package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	infra_adapters "github.com/mariojuniortrab/hauling-api/internal/infra/adapters"
	user_mysql_repository "github.com/mariojuniortrab/hauling-api/internal/infra/repository/mysql/user"
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
	urlParser := infra_adapters.NewChiUrlParserAdapter()

	//Repositories
	userRepository := user_mysql_repository.NewRepositoryMysql(db)

	//Routes
	userRouter := user_routes.NewRouter(userRepository, validator, encrypter, tokenizer, urlParser)

	//Using chi with an adapter to manage routes
	r := infra_adapters.NewChiRouteAdapter()

	//routing
	userRouter.Route(r)

	//starting server
	fmt.Println("Server has started")
	http.ListenAndServe(":8000", r)

}
