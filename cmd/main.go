package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/weezqyd/hello-go/handlers"
	"log"
	"os"
)

func main() {
	var port, dburl string
	flag.StringVar(&port, "port", ":3000", "port to listen")
	flag.StringVar(&dburl, "dsn", "homestead:secret@tcp(192.168.10.10:3306)/hello", "database dsn")
	flag.Parse()
	
	conn, err := sql.Open("mysql", dburl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()
	
	controller := &handlers.Controller{DB: conn}
	app := fiber.New()
	app.Get("/", controller.Welcome)
	app.Get("/widgets", controller.Widgets)
	
	log.Fatal(app.Listen(port))
}
