package app

import (
	"fmt"
	"microservicego/mvc/controllers"
	"net/http"
)

// StartApp : func
func StartApp() {
	fmt.Println("app started")
	http.HandleFunc("/users", controllers.GetUser)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
