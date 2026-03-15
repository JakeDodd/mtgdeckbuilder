package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/JakeDodd/mtgdeckbuilder/models"
	"github.com/JakeDodd/mtgdeckbuilder/service"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()
	host := os.Getenv("pg_host")
	dbport, _ := strconv.Atoi(os.Getenv("pg_port"))
	user := os.Getenv("pg_user")
	password := os.Getenv("pg_password")
	dbname := os.Getenv("pg_dbname")
	sslmode := os.Getenv("sslmode")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		host, dbport, user, password, dbname, sslmode)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the environment")
	}
	port = ":" + port

	//read file and write to database

	// e is an Echo http server, see above import
	// dont know much about this library
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/random-card", func(c echo.Context) error {
		card, err := database.GetRandomCard(db)
		if err != nil {
			log.Fatal(err)
		}
		prints, err := database.GetPrintsByName(db, card.CardName)
		if err != nil {
			log.Fatal(err)
		}
		var english_print models.Prints
		for i := 0; i < len(prints); i++ {
			if prints[i].Lang == "en" {
				english_print = prints[i]
				return c.JSON(200, english_print)
			}
		}
		return c.JSON(200, prints[0])
	})

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Skipper:    nil,
		Root:       "react/dist",
		Index:      "index.html",
		HTML5:      true,
		Browse:     false,
		IgnoreBase: false,
		Filesystem: nil,
	}))

	log.Printf("Server starting on port: %v", port)

	// Here we start the server, this starts a loop which will only end if an error
	// is thrown while processing a request, which is why we log.Fatal the returned value
	e.Logger.Fatal(e.Start(port))

}
