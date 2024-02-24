package main

import (
	"credit/api"
	"credit/database"
	"credit/repository"
	"credit/service"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	// Establecer la conexión a la base de datos PostgreSQL
	postgresRepo, err := database.NewPostgreSQLRepository("postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable")
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	defer postgresRepo.Close()

	// Inicializar el repositorio de créditos
	creditRepository := repository.NewCreditRepository()

	// Inicializar el servicio de créditos
	creditService := service.NewCreditService(creditRepository)

	// Iniciar el servidor
	mux := api.InitRoutes(creditService, postgresRepo)
	log.Println("Servidor escuchando en el puerto 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
