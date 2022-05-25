package main

import (
	"log"
	"encoding/json"
	"fmt"
	"net/http"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database Connection
// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
var dsn = "root:faded@tcp(127.0.0.1:3306)/go_test_models?charset=utf8mb4&parseTime=True&loc=Local"
var db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

type GoTestModel struct {
	Name string 
	Year string 
}

func main() {
	fmt.Println("My sql, not yours")
	http.HandleFunc("/createstuff", GoDatabaseCreate)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func GoDatabaseCreate(w http.ResponseWriter, r *http.Request) {
	GoTestModel := GoTestModel{
		Name: "Gabe",
		Year: "2022",
	}

	db.Create(&GoTestModel)
	if err := db.Create(&GoTestModel).Error; err != nil {
		log.Fatalln(err)
	}

	json.NewEncoder(w).Encode(GoTestModel)

	fmt.Println("Fields Added", GoTestModel)
}