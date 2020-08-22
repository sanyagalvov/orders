Vue.component("product_list_view", {
    props: ["products"],
    template: `
    <div>
        <productFormCard></productFormCard>
        <div class="row row-cols-1 row-cols-md-3">
            <productCard
                v-for="product in products"
                v-bind:key="product.id"
                v-bind:product="product"
            ></productCard>
        </div>
    </div>`
})
Vue.component("productCard", {
    props: ["product"],
    template: `
    <div class="col mb-4">
        <div class="card">
            <h5 class="card-header">{{ product.name }}</h5>
            <div class="card-body">
                <p class="card-text">
                    ID: <b>{{ product.id }}</b><br/>
                    Measure unit: <b>{{ product.unit }}</b>
                </p>
            </div>
        </div>
    </div>
    `
})
Vue.component("productFormCard", {
    data: function () {
        return {
            name: "",
            unit: ""
        }
    },
    methods: {
        send: function() {
            send_product(this.name, this.unit)
            this.name = ""
            this.unit = ""
        }
    },
    computed: {
        isButtonDisabled: function() {
            if ((this.name === "") || (this.unit === "")) {
                return true
            }
            return false
        }
    },
    template: `
    <div class="card mb-4">
        <div class="card-header">
            <ul class="nav justify-content-end">
                <h5 class="mb-0 mr-auto">Add new product</h5>
                <a data-toggle="collapse" href="#collapse" role="button" aria-expanded="false" aria-controls="collapseExample">
                    Show/Hide
                </a>
            </ul>
        </div>
        <div class="collapse" id="collapse">
        <div class="card-body">
            <div>
                <div class="form-group">
                    <label class="card-text" for="productName">Product name</label>
                    <input v-model="name" class="form-control" id="productName" aria-describedby="emailHelp">
                </div>
                <div class="form-group">
                    <label class="card-text" for="productUnit">Product measure unit</label>
                    <input v-model="unit" class="form-control" id="productUnit">
                </div>
                <button v-on:click="send(name, unit)" class="btn btn-primary" v-bind:disabled="isButtonDisabled">Submit</button>
            </div>
        </div>
        </div>
    </div>
    `
})

Vue.component("order_list_view", {
    props: ["orders", "products"],
    template: `
    <div>
        <orderFormCard :products="products"></orderFormCard>
        <datePickerCard></datePickerCard>
        <div class="row row-cols-1 row-cols-md-2">
            <orderCard
                v-for="order in orders"
                v-bind:key="order.id"
                v-bind:order="order"
            ></orderCard>
        </div>
    </div>`
})
Vue.component("orderCard", {
    props: ["order"],
    computed: {
        allItems: function() {
            return this.order.items.length
        },
        submittedItems: function() {
            return this.order.items.filter(function(value, index, array) {return value.is_submitted}).length
        },
        isButtonDisabled: function() {
            for (let i = 0; i < this.order.items.length; i++) {
                if (this.order.items[i].is_submitted != true){
                    return true
                }
            }
            return false
        }
    },
    methods: {
        send_order: function() {
            this.order = send_update_order(this.order)
        }
    },
    template: `
    <div class="col mb-4">
        <div class="card">
            <h5 class="card-header">{{ order.recipient }}</h5>
            <div class="card-body">
                <p class="card-text">
                    Shipping date: <b>{{ order.date.slice(0, 10) }}</b><br>
                    Items: <b>{{ submittedItems }}/{{ allItems }}</b>
                    <a 
                        data-toggle="collapse" 
                        :href="'#id'+String(order.id)" 
                        role="button" 
                        aria-expanded="false" 
                        aria-controls="collapseExample" 
                        v-if="order.items.length">
                        Show/Hide
                    </a><br>
                    Status: <b class="text-success" v-if="order.is_submitted">Submitted! ;)<br></b>
                            <b class="text-danger" v-else>Not submitted<br></b>
                </p>
                <div class="form-group mb-0">
                    <label for="comment">Comment: </label>
                    <textarea class="form-control" id="comment" rows=1 v-model="order.comment" :disabled="order.is_submitted"></textarea>
                </div>
            </div>
            <div class="collapse" :id="'id'+String(order.id)">
            <ul class="list-group list-group-flush">
                <orderItem 
                    v-for="item in order.items"
                    v-bind:key="item.id"
                    v-bind:item="item"
                ></orderItem>
            </ul>
            </div>
            <div class="card-body pt-0" v-if="order.is_submitted != true">
                <button class="btn btn-primary" v-bind:disabled="isButtonDisabled" v-on:click="send_order()">Submit</button>
            </div>

        </div>
    </div>
    `
})
Vue.component("orderItem", {
    props: ["item"],
    data: function(){
        return {
            aq: "",
            bn: ""
        }
    },
    computed: {
        isButtonDisabled: function() {
            if (this.aq === "" || this.bn === ""){
                return true
            }
            return false
        }
    },
    methods: {
        send_item: function() {
            this.item = send_update_order_item(this.item, this.aq, this.bn)
        }
    },
    template: `
        <li class="list-group-item">
            <label for="aq"><b>{{ item.product.name }}</b></label>
            <div class="input-group" v-if="!item.is_submitted">
                <div class="input-group-prepend">
                    <span class="input-group-text">{{item.rq}} {{ item.product.unit }}</span>
                </div>
                <input type="number" v-model="aq" step=0.01 id="aq" class="form-control">
                <div class="input-group-append">
                    <span class="input-group-text">{{ item.product.unit }}</span>
                    <button class="btn btn-outline-primary" type="button" v-bind:disabled="isButtonDisabled" v-on:click="send_item()">
                        <svg width="1.3em" height="1.3em" viewBox="0 0 16 16" class="bi bi-box-seam" fill="currentColor" xmlns="http://www.w3.org/2000/svg">
                            <path fill-rule="evenodd" d="M8.186 1.113a.5.5 0 0 0-.372 0L1.846 3.5l2.404.961L10.404 2l-2.218-.887zm3.564 1.426L5.596 5 8 5.961 14.154 3.5l-2.404-.961zm3.25 1.7l-6.5 2.6v7.922l6.5-2.6V4.24zM7.5 14.762V6.838L1 4.239v7.923l6.5 2.6zM7.443.184a1.5 1.5 0 0 1 1.114 0l7.129 2.852A.5.5 0 0 1 16 3.5v8.662a1 1 0 0 1-.629.928l-7.185 2.874a.5.5 0 0 1-.372 0L.63 13.09a1 1 0 0 1-.63-.928V3.5a.5.5 0 0 1 .314-.464L7.443.184z"/>
                        </svg>   
                    </button>
                </div>
            </div>
            <div class="input-group my-3"  v-if="!item.is_submitted">
                <div class="input-group-prepend">
                    <span class="input-group-text" id="batch-number">Batch №</span>
                </div>
                <input type="text" v-model="bn" class="form-control" aria-describedby="batch-number">
            </div>
            <p class="card-text" v-else>
                Required Quantity: <b>{{item.rq}}</b><br/>
                Actual Quantity: <b>{{item.aq}}</b><br/>
                Batch number: <b>{{item.bn}}</b>
            </p>
        </li>
    `
})
Vue.component("datePickerCard", {
    data:function(){
        return {
            date: "",
            timer: null
        }
    },
    methods: {
        update: function(){
            update_orders(this.date)
        },
        setTimer: function(){
            this.update()
            this.timer = setInterval(this.update, 3000)
        },
        clearTimer: function(){
            this.update()
            clearInterval(this.timer)
            this.timer = null
        }
    },
    template: `
        <div class="input-group mb-4">
            <div class="input-group-prepend">
                <span class="input-group-text bg-light">Pick a date</span>
            </div>
            <input type="date" v-model="date" class="form-control">
            <div class="input-group-append">
                <button class="btn btn-primary" v-on:click="update()" :disabled="date == ''">Load</button>
                <button class="btn btn-primary" v-on:click="setTimer()" v-if="!timer" :disabled="date == ''">Track automaticaly</button>
                <button class="btn btn-primary" v-on:click="clearTimer()" v-if="timer">Stop tracking</button>
            </div>
        </div>
    `
})
Vue.component("orderFormCard", {
    props: ["products"],
    data: function () {
        return {
            recipient: "",
            date: "",
            items: [],
            comment: ""
        }
    },
    methods: {
        send: function() {
            send_order(this.recipient, this.date, this.items, this.comment)
            this.recipient = ""
            this.date = ""
            this.items = []
            this.comment = ""
        }
    },
    computed: {
        isButtonDisabled: function() {
            if (this.recipient == "") {
                return true
            }
            if (this.date == "") {
                return true
            }
            if (this.items.length <= 0) {
                return true
            }
            for (let i = 0; i < this.items.length; i++) {
                if (!this.items[i].product || !this.items[i].rq) {
                    return true
                }
            }
            return false
        }
    },
    template: `
    <div class="card-group mb-4">
        <div class="card">
            <div class="card-header">
                <ul class="nav justify-content-end">
                    <h5 class="mb-0 mr-auto">New order</h5>
                    <a data-toggle="collapse" href="#collapseOrderForm" role="button" aria-expanded="false" aria-controls="collapseExample">
                        Show/Hide
                    </a>
                </ul>
            </div>
            
            <div class="collapse" id="collapseOrderForm">
            <div class="card-body pb-0">
                <div class="form-group">
                    <label for="recipient">Recipient: </label>
                    <input type="text" class="form-control" v-model="recipient" id="recipient">
                </div>
                <div class="form-group">
                    <label for="shippingDate">Shipping date: </label>
                    <input type="date" class="form-control" v-model="date" id="shippingDate">
                </div>
                <div class="form-group">
                    <label for="comment">Comment: </label>
                    <textarea class="form-control" v-model="comment" id="comment" rows=2></textarea>
                </div>


                <div class="mt-auto mb-3 text-right">
                    <button type="submit" v-on:click="send()" class="btn btn-primary" v-bind:disabled="isButtonDisabled">Submit</button>
                    <button v-on:click="items.push({})" class="btn btn-primary">+</button>
                    <button v-on:click="items.pop()" class="btn btn-primary" v-bind:disabled="!items.length">-</button>
                </div>
            </div>
            </div>



        </div>

        <div class="card">
        
            <h5 class="card-header">Items</h5>
            <div class="collapse" id="collapseOrderForm">
            <ul class="list-group list-group-flush border-top-0">
                <orderItemForm 
                    v-for="item in items"
                    :key="item.id"
                    :item="item"
                    :products="products"
                ></orderItemForm>
            </ul>
            </div>
        </div>

    </div>
    `
})
Vue.component("orderItemForm", {
    model: {
        prop: 'item'
    },
    props: ["products", "item"],
    template: `
    <li class="list-group-item">
        <div class="input-group mb-2">
            <div class="input-group-prepend">
                <label class="input-group-text" for="product">Product</label>
            </div>
            <select class="custom-select" v-model="item.product" id="product">
                <option v-for="product in products" v-bind:value="product">
                    {{ product.name }}
                </option>
            </select>
        </div>
        <div class="input-group mb-2">
            <div class="input-group-prepend">
                <label class="input-group-text" for="rq"> Req. Q. </label>
            </div>
            <input v-model="item.rq" type="number" step=0.01 class="form-control ml-auto" id="rq" placeholder="0">
            <div class="input-group-append" v-if="item.product">
                <span class="input-group-text" id="unit">{{ item.product.unit }}</span>
            </div>
        </div>
    </li>
    `
})


var app = new Vue({
    el: '#app',
    data: {
        order_view: false,
        product_list_view: false,
        order_list_view: false,

        products: [],
        orders: []
    },
    created () {
        update_products()
        this.product_list_view = true
    },
    methods: {
        show_order_list_view: function() {
            this.order_view = false
            this.product_list_view = false
            this.order_list_view = true
        },
        show_product_list_view: function() {
            update_products()
            this.order_view = false
            this.order_list_view = false
            this.product_list_view = true
        }
    }
})

function update_products() {
    fetch("/api/products")
    .then(function (response) {
        return response.json();
    })
    .then(function (data) {
        app.products = data
    })
    .catch(function (error) {
        console.log("Error: " + error);
    });
}

function send_product(name, unit) {
    var product = {
        name: name,
        unit: unit
    }

    fetch("/api/products", {
        method: "POST",
        mode: "same-origin",
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(product)
    })
    .then(function (response) {
        return response.json();
    })
    .then(function (data) {
        console.log(data)
        if (app.products == null) {
            app.products = data
        } else {
        app.products.push(data)
        }
    })
    .catch(function (error) {
        console.log("Error: " + error);
    });
}

var bell = new Audio("/x/bell.mp3")
function update_orders(date) {
    
    fetch("/api/orders/" + date)
    .then(function (response) {
        return response.json();
    })
    .then(function (data) {
        if (JSON.stringify(app.orders) != JSON.stringify(data)) {
            bell.play()
            app.orders = data
        }
        console.log(app.orders)
    })
    .catch(function (error) {
        console.log("Error: " + error);
    });
}

function send_update_order_item(item, aq, bn) {

    item.aq = Number(aq)
    item.bn = bn
    item.is_submitted = true
    console.log(item)

    fetch("/api/order-items", {
        method: "POST",
        mode: "same-origin",
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(item)
    })
    .then(function (response) {
        return response.json();
    })
    .then(function (data) {
        console.log(data);
        item = data
    })
    .catch(function (error) {
        console.log("Error: " + error);
    });
    return item
}

function send_update_order(order) {
    order.is_submitted = true
    console.log(order)

    fetch("/api/update-order", {
        method: "POST",
        mode: "same-origin",
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(order)
    })
    .then(function (response) {
        return response.json();
    })
    .then(function (data) {
        console.log(data);
        order = data
    })
    .catch(function (error) {
        console.log("Error: " + error);
    });
    return order
}

function send_order(recipient, date, items, comment) {
    var order = {
        recipient: recipient,
        date: date + "T18:25:43.511Z",
        items: items,
        comment: comment,
        is_submitted: false
    }

    order.items.forEach(element => {
        element.rq = Number(element.rq)
        element.is_submitted = false
    });

    fetch("/api/orders", {
        method: "POST",
        mode: "same-origin",
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(order)
    })
    .then(function (response) {
        return response.json();
    })
    .then(function (data) {
        console.log(data);
        app.orders.push(data);
    })
    .catch(function (error) {
        console.log("Error: " + error);
    });
}