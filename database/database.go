package database

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
)

// Initialize the App with router and DB
type App struct {
	Router *mux.Router
	DB     *pgx.Conn
}

func InitializeApp() *App {
	app := new(App)
	app.ConnectToDB()
	return app
}

// Make DB connection
func (a *App) ConnectToDB() *pgx.Conn {
	var dbname, username, password, host, port = "gin_goorm_rest", "sm", "", "localhost", "5432"

	// QueryEscape is used to escape the breaking special chars
	connectionString :=
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s", url.QueryEscape(username), url.QueryEscape(password), url.QueryEscape(host), port, dbname)

	var err error
	conn, err := pgx.Connect(context.Background(), connectionString)

	a.DB = conn
	if err != nil {
		// log if error
		log.Fatal(err)
	}
	return conn
}
