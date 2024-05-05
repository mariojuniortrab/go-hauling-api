package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	brand_repository "github.com/mariojuniortrab/hauling-api/internal/infra/repository/brand"
	brand_handler "github.com/mariojuniortrab/hauling-api/internal/infra/web/handlers/brand"
	brand_usecase "github.com/mariojuniortrab/hauling-api/internal/usecase/brand"
	validation "github.com/mariojuniortrab/hauling-api/internal/validation"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/hauling")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	brandRepository := brand_repository.NewBrandRepositoryMysql(db)
	createBrandUseCase := brand_usecase.NewCreateBrandUseCase(brandRepository)
	listBrandUseCase := brand_usecase.NewListBrandUseCase(brandRepository)
	createBrandValidation := validation.NewCreateBrandValidation(brandRepository)

	createBrandHandler := brand_handler.NewCreateBrandHandler(createBrandUseCase, createBrandValidation)
	listBrandHandler := brand_handler.NewListBrandHandler(listBrandUseCase)

	r := chi.NewRouter()
	r.Post("/brands", createBrandHandler.Handle)
	r.Get("/brands", listBrandHandler.Handle)

	fmt.Println("Server has started")
	http.ListenAndServe(":8000", r)

}
