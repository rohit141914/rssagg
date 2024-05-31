package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/rohit141914/rssagg/internal/database"

	// "github.com/postgres/rssagg/internal/database"
	_ "github.com/lib/pq"

)

func main() {
	// fmt.Println("Hello, World!")

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString==""{
		log.Fatal("Port is not set")
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL==""{
		log.Fatal("DBURL is not found in the environment")
	}
	conn,err :=sql.Open("postgres",dbURL)
	if err!=nil{
		log.Fatal("can't connect to the database: ", err)
	}
	// queries,err := database.New(conn)
	// if err!=nil{
	// 	log.Fatal("can't create to db connection: ", err)
	// }
	// APIConfig := 
	apiCFG := apiConfig{
		DB: database.New(conn),
	}
	router :=chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300, // Maximum value not ignored by any of major browsers
	}),
	)

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/users", apiCFG.handlerCreateUser)

	router.Mount("/v1", v1Router)

	srv :=&http.Server{
		Handler: router,
		Addr: ":"+portString,
	}
	log.Println("Server is running on port: ", portString)
	err=srv.ListenAndServe()
	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("Port is set to: ", portString)
}
