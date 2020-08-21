package storage

import (
	"alex/fishorder-api-v3/app2/models"
	"context"
	"database/sql"
	"fmt"
	"time"

	// Postgres driver
	_ "github.com/lib/pq"
	"golang.org/x/sync/errgroup"
)

// Storage ...
type Storage struct {
	config *Config
	db     *sql.DB
	ctx    context.Context
}

// New ...
func New(ctx context.Context, config *Config) (*Storage, error) {
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("storage: New: Config is not valid: %v", err)
	}
	return &Storage{
		config: config,
		ctx:    ctx,
	}, nil
}

// Connect ...
func (st *Storage) Connect() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		st.config.Host, st.config.Port, st.config.User, st.config.Password, st.config.DBname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	st.db = db
	err = st.db.Ping()
	if err != nil {
		return err
	}
	return nil
}

// InsertProduct ...
func (st *Storage) InsertProduct(p *models.Product) error {
	if err := p.Validate(); err != nil {
		err := fmt.Errorf("storage: InsertProduct: Product is not valid: %v", err)
		return err
	}

	sqlStatement := `
	INSERT INTO products (name, unit)
	VALUES ($1, $2)
	RETURNING id`
	id := 0
	err := st.db.QueryRow(sqlStatement, p.Name, p.Unit).Scan(&id)
	if err != nil {
		return err
	}
	p.ID = id
	return nil
}

// SelectProduct ...
func (st *Storage) SelectProduct(id int) (*models.Product, error) {
	sqlStatement := "SELECT name, unit FROM products WHERE id=$1;"
	row := st.db.QueryRow(sqlStatement, id)
	var p models.Product
	switch err := row.Scan(&p.Name, &p.Unit); err {
	case sql.ErrNoRows:
		return nil, fmt.Errorf("storage: SelectProduct: no rows were returned")
	case nil:
		p.ID = id
		return &p, nil
	default:
		return nil, err
	}
}

// SelectAllProducts ...
func (st *Storage) SelectAllProducts() (*[]models.Product, error) {
	rows, err := st.db.Query(`SELECT id, name, unit FROM products`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		err = rows.Scan(&p.ID, &p.Name, &p.Unit)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &products, nil
}

// InsertOrderItem ...
func (st *Storage) InsertOrderItem(item *models.OrderItem, orderID int) error {
	if err := item.Validate(); err != nil {
		err := fmt.Errorf("storage: InsertOrder: OrderItem is not valid: %v", err)
		return err
	}
	sqlStatement := `
	INSERT INTO order_items(product_id, rq, bn, comment, is_submitted, order_id)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id`

	id := 0
	err := st.db.QueryRow(
		sqlStatement,
		item.Product.ID,
		item.RequiredQuantity,
		item.BatchNumber,
		item.Comment,
		item.IsSubmitted,
		orderID,
	).Scan(&id)
	if err != nil {
		return err
	}
	item.ID = id
	return nil
}

// UpdateOrderItem ...
func (st *Storage) UpdateOrderItem(item *models.OrderItem) error {
	if err := item.Validate(); err != nil {
		err := fmt.Errorf("storage: InsertOrder: OrderItem is not valid: %v", err)
		return err
	}
	sqlStatement := `
	UPDATE order_items 
	SET is_submitted = $2, aq = $3, bn = $4
	WHERE id = $1;`

	res, err := st.db.Exec(
		sqlStatement,
		item.ID,
		item.IsSubmitted,
		item.ActualQuantity,
		item.BatchNumber,
	)

	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

// InsertOrderItems inserts multiple items.
func (st *Storage) InsertOrderItems(items *[]models.OrderItem, orderID int) error {
	group, _ := errgroup.WithContext(st.ctx)
	for _, item := range *items {
		group.Go(func() error {
			return st.InsertOrderItem(&item, orderID)
		})
	}
	if err := group.Wait(); err != nil {
		return err
	}
	return nil
}

// SelectItemsByOrder ...
func (st *Storage) SelectItemsByOrder(o *models.Order) error {
	rows, err := st.db.Query(`
	SELECT id, product_id, rq, aq, bn, comment, is_submitted 
	FROM order_items WHERE order_id = $1`, o.ID)
	if err != nil {
		return err
	}
	defer rows.Close()
	items := make([]models.OrderItem, 0)
	for rows.Next() {
		var (
			i         models.OrderItem
			aq        sql.NullFloat64
			productID int
		)
		err = rows.Scan(
			&i.ID,
			&productID,
			&i.RequiredQuantity,
			&aq,
			&i.BatchNumber,
			&i.Comment,
			&i.IsSubmitted,
		)
		if err != nil {
			return err
		}
		if aq.Valid {
			i.ActualQuantity = float32(aq.Float64)
		}

		product, err := st.SelectProduct(productID)
		if err != nil {
			return err
		}
		i.Product = *product

		items = append(items, i)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	o.Items = items
	return nil
}

// InsertOrder ...
func (st *Storage) InsertOrder(o *models.Order) error {
	if err := o.Validate(); err != nil {
		err := fmt.Errorf("storage: InsertOrder: Order is not valid: %v", err)
		return err
	}

	sqlStatement := `
	INSERT INTO orders (recipient, shipping_date, comment, is_submitted)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	id := 0
	err := st.db.QueryRow(
		sqlStatement,
		o.Recipient,
		o.ShippingDate,
		o.Comment,
		o.IsSubmitted,
	).Scan(&id)
	if err != nil {
		return err
	}
	err = st.InsertOrderItems(&o.Items, id)
	if err != nil {
		return err
	}
	o.ID = id
	return nil
}

// UpdateOrder ...
func (st *Storage) UpdateOrder(o *models.Order) error {
	if err := o.Validate(); err != nil {
		err := fmt.Errorf("storage: InsertOrder: Order is not valid: %v", err)
		return err
	}

	sqlStatement := `
	UPDATE orders 
	SET comment = $2, is_submitted = $3
	WHERE id = $1`

	res, err := st.db.Exec(
		sqlStatement,
		o.ID,
		o.Comment,
		o.IsSubmitted,
	)

	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

// SelectOrdersByDate ...
func (st *Storage) SelectOrdersByDate(date time.Time) (*[]models.Order, error) {
	rows, err := st.db.Query(
		`SELECT id, recipient, shipping_date, comment, is_submitted FROM orders WHERE shipping_date=($1)`,
		date,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var orders []models.Order
	for rows.Next() {
		o := models.Order{}
		err = rows.Scan(
			&o.ID,
			&o.Recipient,
			&o.ShippingDate,
			&o.Comment,
			&o.IsSubmitted,
		)
		if err != nil {
			return nil, err
		}
		err := st.SelectItemsByOrder(&o)
		if err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &orders, nil
}
