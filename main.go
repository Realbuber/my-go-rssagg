package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/go-chi/cors"
	"github.com/go-chi/chi"
)

func main(){

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString =="" {
		log.Fatal("PORT is not found in environemnt")
	}

	router := chi.NewRouter()
	//just open everything to request for testing.
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:    []string{"https://*","http://*"},
		AllowedMethods:    []string{"GET","POST","DELETE","OPTIONS"},
		AllowedHeaders:    []string{"*"},
		ExposedHeaders:    []string{"Link"},
		AllowCredentials: false,
		MaxAge:			   300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err",handlerErr)



	//add a v1 prefix
	router.Mount("/v1",v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:	":" + portString,
	}

	log.Printf("Server starting on port %v",portString)

	err := srv.ListenAndServe()
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Port:",portString)
}
