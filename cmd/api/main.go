package main

import (
	"database/sql"
	"fmt"
	"net/http"

	brand_repository "github.com/mariojuniortrab/hauling-api/internal/infra/repository/brand"
	"github.com/mariojuniortrab/hauling-api/internal/infra/web/routes"
	brand_routes "github.com/mariojuniortrab/hauling-api/internal/infra/web/routes/brand"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/hauling")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//Repositories
	brandRepository := brand_repository.NewBrandRepositoryMysql(db)

	//Routes
	brandRouter := brand_routes.NewRouter(brandRepository)

	//Using chi with an adapter to manege routes
	r := routes.NewChiRouteAdapter()

	//routing
	brandRouter.Route(r)

	//starting server
	fmt.Println("Server has started")
	http.ListenAndServe(":8000", r)

}
