package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"middleware-server/api"
	"middleware-server/database"
	"net/http"
	"os"

	"github.com/cockroachdb/cockroach-go/v2/crdb"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Message struct {
	Value string `json:"value"`
}

var db *sql.DB

// ! handler function
func home(c echo.Context) error {
	return rootHandler(db, c)
}
func ping(c echo.Context) error {
	return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
}
func send(c echo.Context) error {
	return sendHandler(db, c)
}

// ! handlers
func rootHandler(db *sql.DB, c echo.Context) error {
	r, err := api.CountRecords(db)
	if err != nil {
		return c.HTML(http.StatusInternalServerError, err.Error())
	}
	return c.HTML(http.StatusOK, fmt.Sprintf("Hello, Docker! (%d)\n", r))
}

func sendHandler(db *sql.DB, c echo.Context) error {
	m := &Message{}

	// ! Bind 함수는 request body data를 interface에 바인딩하는 함수
	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	err := crdb.ExecuteTx(context.Background(), db, nil, func(tx *sql.Tx) error {
		_, err := tx.Exec("INSERT INTO message (value) VALUES ($1) ON CONFLICT (value) DO UPDATE SET value = excluded.value",
			m.Value)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return nil
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, m)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db, err := database.InitStore()
	if err != nil {
		log.Fatalf("failed to initialise the store: %s", err)
	}
	defer db.Close()

	e.GET("/", home)
	e.GET("/ping", ping)
	e.POST("/send", send)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8081"
	}
	e.Logger.Fatal(e.Start(":" + httpPort))
}
