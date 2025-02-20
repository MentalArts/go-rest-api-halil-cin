package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"fmt"
)

type Response struct {
	Msg string `json:"message"`
}

func main() {
	router := gin.Default()
	router.GET("/ping", handlePing)
	router.GET("/hello/:name", handleHello)
	router.Run(":8000")
}

func handlePing(c *gin.Context) {
	res := Response{Msg: "pong"}
	c.JSON(http.StatusOK, res)
}

func handleHello(c *gin.Context) {
	name := c.Param("name")

	msg := fmt.Sprintf("Welcome, %s", name)

	c.String(http.StatusOK, msg)
}

// Vanilla implementation
// func main() {
// 	http.HandleFunc("GET /ping", handlePing)
// 	log.Println("Server listening...")
// 	log.Fatal(http.ListenAndServe(":8000", nil))
// }

// func handlePing(w http.ResponseWriter, r *http.Request) {
// 	res := Response{Msg: "pong"}
// 	json.NewEncoder(w).Encode(res)
// 	w.WriteHeader(http.StatusOK)
// 	log.Println("Request recieved")
// }
