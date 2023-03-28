package main

import (
    "fmt"
    "log"
    "os"
    "net/url"

    "database/sql"
	"api/routers"

    "github.com/joho/godotenv"
    "github.com/gofiber/fiber/v2"
    _ "github.com/denisenkom/go-mssqldb"
)

func init(){
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func dbConnection() *sql.DB{

    query := url.Values{}
    query.Add("app name", "MyAppName")

    u := &url.URL{
        Scheme:   "sqlserver",
        User:     url.UserPassword(os.Getenv("user"), os.Getenv("password")),
        Host:     fmt.Sprintf("%s:%s", os.Getenv("host"), os.Getenv("port")),
    }
    db, err := sql.Open("sqlserver", u.String())

    err = db.Ping()
    if err != nil {
        panic(err)
    }
    
    return db

}

func main() {
    app := fiber.New()

    db := dbConnection()
    fmt.Println("Connection Successful")

    routers.TodoNotes(app, db)
	routers.ApiTestCalls(app, db)

    log.Fatal(app.Listen(":3000"))
}