FROM golang:alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
# Add this go mod download command to pull in any dependencies
RUN go mod download
# Our project will now successfully build with the necessary go libraries included.
RUN go build -o main .
# Our start command which kicks off
# our newly created binary executable
EXPOSE 8080 5432

CMD ["/app/main"]