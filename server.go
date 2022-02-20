package main

import (
	"famesensor/go-graphql-jwt/database"
	"famesensor/go-graphql-jwt/directives"
	"famesensor/go-graphql-jwt/graph"
	"famesensor/go-graphql-jwt/graph/generated"
	"famesensor/go-graphql-jwt/middlewares"
	"famesensor/go-graphql-jwt/models"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = defaultPort
	}

	// connect database
	DB := database.GetDB()
	defer func() {
		db, _ := DB.DB()
		database.DisconnectDatabase(db)
	}()

	err := database.Migrate(&models.User{})
	if err != nil {
		fmt.Printf("Migrate database error : %v", err)
		os.Exit(0)
	}

	router := mux.NewRouter()
	router.Use(middlewares.AuthMiddleware)

	c := generated.Config{Resolvers: &graph.Resolver{}}
	c.Directives.Auth = directives.Auth

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
