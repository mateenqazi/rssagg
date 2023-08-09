package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mateenqazi/rssagg/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the env")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB URL is not found in the env")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database.", err)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", HandlerReadiness)
	v1Router.Get("/error", HandlerErr)
	v1Router.Post("/users", apiCfg.handlerUsersCreate)
	v1Router.Get("/users", apiCfg.MiddlewareAuth(apiCfg.handlerUsersGet))

	router.Mount("/v1", v1Router)

	srv := http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Println("server starting on  port ", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
