package main

import (
	"database/sql"
	"fmt"
	"net/http"

	brand_repository "github.com/mariojuniortrab/hauling-api/internal/infra/repository/brand"
	user_repository "github.com/mariojuniortrab/hauling-api/internal/infra/repository/user"
	"github.com/mariojuniortrab/hauling-api/internal/infra/web/routes"
	brand_routes "github.com/mariojuniortrab/hauling-api/internal/infra/web/routes/brand"
	user_routes "github.com/mariojuniortrab/hauling-api/internal/infra/web/routes/user"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/hauling")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//Repositories
	brandRepository := brand_repository.NewRepositoryMysql(db)
	userRepository := user_repository.NewRepositoryMysql(db)

	//Routes
	brandRouter := brand_routes.NewRouter(brandRepository)
	userRouter := user_routes.NewRouter(userRepository)

	//Using chi with an adapter to manege routes
	r := routes.NewChiRouteAdapter()

	//routing
	brandRouter.Route(r)
	userRouter.Route(r)

	//starting server
	fmt.Println("Server has started")
	http.ListenAndServe(":8000", r)

}
