package main

import (
	"clothes-shop/configs"
	"errors"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Products struct {
	ProductId int    `db:"product_id" json:"product_id"`
	Gender    string `db:"gender" json:"gender"`
	Style     string `db:"style" json:"style"`
	Size      string `db:"size" json:"size"`
	Price     int    `db:"price" json:"price"`
}

type Orders struct {
	OrderId   int    `db:"order_id" json:"order_id"`
	Status    string `db:"status" json:"status"`
	OrderDate string `db:"order_date" json:"order_date"`
	PaidDate  string `db:"paid_date" json:"paid_date"`
	Address   string `db:"address" json:"address"`
}

type ProductDetails struct {
	ProductId []int    `db:"product_id" json:"product_id"`
	Gender    []string `db:"gender" json:"gender"`
	Style     []string `db:"style" json:"style"`
	Size      []string `db:"size" json:"size"`
	Price     []int    `db:"price" json:"price"`
	Quantity  []int    `db:"quantity" json:"quantity"`
}

type OrderDetails struct {
	ProductDetails ProductDetails `db:"product_details" json:"product_details"`
	Quantity       int            `db:"quantity" json:"quantity"`
	Address        string         `db:"address" json:"address"`
}

type QueryParams struct {
	Gender    []string `query:"gender"`
	Style     []string `query:"style"`
	Size      []string `query:"size"`
	Price     []int    `query:"price"`
	Status    []string `query:"status"`
	StartDate string   `query:"start_date"`
	EndDate   string   `query:"end_date"`
	Limit     int      `query:"limit"`
	Page      int      `query:"page"`
}

var db *sqlx.DB

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(errors.New("error loading .env file"))
	}

	cfg := new(configs.Configs)

	cfg.App.Host = os.Getenv("FIBER_HOST")
	cfg.App.Port = os.Getenv("FIBER_PORT")

	cfg.PostgreSQL.Host = os.Getenv("DB_HOST")
	cfg.PostgreSQL.Port = os.Getenv("DB_PORT")
	cfg.PostgreSQL.Database = os.Getenv("DB_DATABASE")
	cfg.PostgreSQL.Username = os.Getenv("DB_USERNAME")
	cfg.PostgreSQL.Password = os.Getenv("DB_PASSWORD")
	cfg.PostgreSQL.SSLMode = os.Getenv("DB_SSL_MODE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.PostgreSQL.Host,
		cfg.PostgreSQL.Port,
		cfg.PostgreSQL.Username,
		cfg.PostgreSQL.Password,
		cfg.PostgreSQL.Database,
		cfg.PostgreSQL.SSLMode,
	)

	db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		panic(err)
	}

	addr := fmt.Sprintf("%s:%s", cfg.App.Host, cfg.App.Port)

	app := fiber.New()

	app.Get("/products", func(c *fiber.Ctx) error {
		queryParams := new(QueryParams)
		err := c.QueryParser(queryParams)
		if err != nil {
			panic(errors.New("cannot connect to Database server"))
		}

		products, err := getProducts(queryParams)
		if err != nil {
			panic(err)
		}

		return c.Status(fiber.StatusOK).JSON(products)
	})

	app.Get("/orders", func(c *fiber.Ctx) error {
		queryParams := new(QueryParams)
		err := c.QueryParser(queryParams)
		if err != nil {
			panic(err)
		}

		orders, err := getOrders(queryParams)
		if err != nil {
			panic(err)
		}

		return c.Status(fiber.StatusOK).JSON(orders)
	})

	app.Post("/orders", func(c *fiber.Ctx) error {
		orderDetails := new(OrderDetails)
		err := c.BodyParser(orderDetails)
		if err != nil {
			panic(err)
		}

		err = createOrder(orderDetails)
		if err != nil {
			panic(err)
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"status": "order has been created",
		})
	})

	app.Listen(addr)
}

func filter(queryParam []string, fieldName string) string {
	if len(queryParam) <= 0 {
		return "( 1 = 1 ) "
	}

	addQuery := fmt.Sprintf(" %s = '%s' ", fieldName, queryParam[0])
	for index := 1; index < len(queryParam); index++ {
		addQuery += fmt.Sprintf("OR %s = '%s' ", fieldName, queryParam[index])
	}
	return fmt.Sprintf("(%s) ", addQuery)
}

func pagination(limit, page int) string {
	if limit == 0 {
		limit = 10
	}
	if page == 0 {
		page += 1
	}
	return fmt.Sprintf("LIMIT %d OFFSET %d", limit, (page-1)*limit)
}

func getProducts(queryParams *QueryParams) ([]Products, error) {
	query := "SELECT * FROM products WHERE "
	query += filter(queryParams.Gender, "gender")
	query += "AND " + filter(queryParams.Style, "style")
	query += "AND " + filter(queryParams.Size, "size")
	query += pagination(queryParams.Limit, queryParams.Page)

	rows, err := db.Queryx(query)
	if err != nil {
		return nil, err
	}

	products := []Products{}
	for rows.Next() {
		product := Products{}
		rows.StructScan(&product)
		products = append(products, product)
	}

	return products, nil
}

func getOrders(queryParams *QueryParams) ([]Orders, error) {
	query := "SELECT * FROM orders WHERE "
	query += fmt.Sprintf("( paid_date BETWEEN '%s' AND '%s'::DATE + INTERVAL '1' DAY ) ", queryParams.StartDate, queryParams.EndDate)
	query += "AND " + filter(queryParams.Status, "status")
	query += pagination(queryParams.Limit, queryParams.Page)

	fmt.Println(query)

	rows, err := db.Queryx(query)
	if err != nil {
		return nil, err
	}

	orders := []Orders{}
	for rows.Next() {
		order := Orders{}
		rows.StructScan(&order)
		orders = append(orders, order)
	}

	return orders, nil
}

func createOrder(orderDetails *OrderDetails) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////
	rows, err := tx.Queryx("INSERT INTO orders(address) VALUES($1) RETURNING order_id", orderDetails.Address)
	if err != nil {
		tx.Rollback()
		return err
	}

	lastInsertId := 0
	rows.Next()
	err = rows.Scan(&lastInsertId)
	if err != nil {
		tx.Rollback()
		return err
	}

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////
	query := fmt.Sprintf("INSERT INTO order_item(order_id, product_id, gender, style, size, price, quantity) VALUES\n(%d, %d, '%s', '%s', '%s', %d, %d)",
		lastInsertId,
		orderDetails.ProductDetails.ProductId[0],
		orderDetails.ProductDetails.Gender[0],
		orderDetails.ProductDetails.Style[0],
		orderDetails.ProductDetails.Size[0],
		orderDetails.ProductDetails.Price[0],
		orderDetails.ProductDetails.Quantity[0],
	)

	for index := 1; index < len(orderDetails.ProductDetails.ProductId); index++ {
		query += fmt.Sprintf(",\n(%d, %d, '%s', '%s', '%s', %d, %d)",
			lastInsertId,
			orderDetails.ProductDetails.ProductId[index],
			orderDetails.ProductDetails.Gender[index],
			orderDetails.ProductDetails.Style[index],
			orderDetails.ProductDetails.Size[index],
			orderDetails.ProductDetails.Price[index],
			orderDetails.ProductDetails.Quantity[index],
		)
	}

	result, err := tx.Exec(query)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}

	if affected <= 0 {
		tx.Rollback()
		return errors.New("the number of affected rows is zero")
	}

	tx.Commit()

	return nil
}
