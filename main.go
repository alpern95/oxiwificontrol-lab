package main

import (
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/alpern95/go-restful-api/book"
	"github.com/alpern95/go-restful-api/db"
	"github.com/alpern95/go-restful-api/auth"
)

func main() {
    db.InitDB()
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	wsContainer := restful.NewContainer()
	wsContainer.Add(book.BookController{}.AddRouters())

    wsContainer.Add(auth.UserController{}.AddRouters())
	// Add container filter to enable CORS
	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"Content-Type", "Accept", "Authorization"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		CookiesAllowed: false,
		Container:      wsContainer}
	wsContainer.Filter(cors.Filter)

	host := "192.168.1.32:8080"
	log.Printf("listening on: %s", host)
	server := &http.Server{Addr: host, Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}
