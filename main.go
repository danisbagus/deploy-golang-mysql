package main

import (
	"fmt"
	"go-project/config"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var status string

func main() {

	db := config.DBInit()

	if err := db.Ping(); err!=nil{
		status = fmt.Sprintf("can't connect database, error on ping: %s", err.Error())
	}else{
		status = fmt.Sprintf("success to connect database")
	}

	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("PORT")

	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Open golang application using docker \n")
        fmt.Fprintln(w, status)
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "About menu \n")
	})
	
	fmt.Println("start server using port:", port)
	fmt.Println("db connection:",status)

	address := fmt.Sprintf(":%s",port)

	err = http.ListenAndServe(address, nil)

	if err != nil {
		fmt.Println(err.Error())
	}
	
}