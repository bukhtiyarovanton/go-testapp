// main
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bukhtiyarovanton/go-testapp/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

var db *gorm.DB

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
		//log.Fatal("$PORT must be set")
	}

	var err error
	db, err = gorm.Open("postgres", "host=localhost dbname=go-testappDB user=AntonBukhtiyarov password=taNk1985 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&models.Todo{}, &models.User{})

	count := 0
	db.Model(&models.User{}).Count(&count)
	if count == 0 {
		for i := 0; i < 2; i++ {
			password, _ := bcrypt.GenerateFromPassword([]byte(fmt.Sprintf("password%v", i+1)), bcrypt.DefaultCost)

			user := models.User{
				Name:     fmt.Sprintf("User %v", i+1),
				Email:    fmt.Sprintf("user%v@mail.com", i+1),
				Password: string(password),
			}

			db.Create(&user)
		}
	}

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":"+port, router))
}
