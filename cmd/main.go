package main

import (
	"fmt"
	"github.com/MCPutro/maxchatTest/internal/controller"
	"github.com/MCPutro/maxchatTest/internal/repository"
	"github.com/MCPutro/maxchatTest/internal/usecase"
	"log"
	"net/http"
)

func main() {

	repo := repository.NewProductRepo()
	productUseCase := usecase.NewProductUseCase(repo)

	// seed
	if err := productUseCase.SeedProductsFromJSON("./data/initialData.json"); err != nil {
		log.Fatalf("Gagal melakukan seed data: %v", err)
	}

	productHandler := controller.NewProductHandler(productUseCase)

	mux := http.NewServeMux()

	// Routing
	mux.HandleFunc("GET /ping", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Pong")
	})
	mux.HandleFunc("POST /products", productHandler.AddNewProduct)
	mux.HandleFunc("GET /products", productHandler.GetAllProducts)
	mux.HandleFunc("GET /products/tech/{tech}", productHandler.GetProductsByTech)
	mux.HandleFunc("GET /products/model/{model}", productHandler.GetProductsByModel)
	mux.HandleFunc("PUT /products/update", productHandler.UpdateProduct)
	mux.HandleFunc("DELETE /products/{code}", productHandler.DeleteProduct)

	// Start the server
	port := ":9999"
	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}
	log.Println("Listening...")
	server.ListenAndServe() // Run the http server
}
