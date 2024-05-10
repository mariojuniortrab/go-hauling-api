package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	infra_adapters "github.com/mariojuniortrab/hauling-api/internal/infra/adapters"
	brand_repository "github.com/mariojuniortrab/hauling-api/internal/infra/repository/brand"
	user_repository "github.com/mariojuniortrab/hauling-api/internal/infra/repository/user"
	brand_routes "github.com/mariojuniortrab/hauling-api/internal/presentation/web/route/brand"
	user_routes "github.com/mariojuniortrab/hauling-api/internal/presentation/web/route/user"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/hauling")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	validator := infra_adapters.NewValidator()
	encrypter := infra_adapters.NewBcryptAdapter()
	tokenizer := infra_adapters.NewJwtAdapter()

	//Repositories
	brandRepository := brand_repository.NewRepositoryMysql(db)
	userRepository := user_repository.NewRepositoryMysql(db)

	//Routes
	brandRouter := brand_routes.NewRouter(brandRepository)
	userRouter := user_routes.NewRouter(userRepository, validator, encrypter, tokenizer)

	//Using chi with an adapter to manage routes
	r := infra_adapters.NewChiRouteAdapter()

	//routing
	brandRouter.Route(r)
	userRouter.Route(r)

	//starting server
	fmt.Println("Server has started")
	http.ListenAndServe(":8000", r)

}
