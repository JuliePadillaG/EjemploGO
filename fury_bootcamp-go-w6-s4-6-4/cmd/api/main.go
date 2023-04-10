package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/mercadolibre/go-meli-toolkit/gomelipass"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/cmd/api/routes"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	// NO MODIFICAR

	var conectionPolicy string
	if len(os.Args) > 1 {
		conectionPolicy = os.Args[1]
	}

	var db *sql.DB
	var err error
	if conectionPolicy == "local" {
		db, err = sql.Open("mysql", "root:@/melisprint")
		if err != nil {
			panic(err)
		}
	} else {
		dbUsername := "bgow6s464_WPROD"
		dbPassword := gomelipass.GetEnv("DB_MYSQL_DESAENV07_BGOW6S464_BGOW6S464_WPROD")
		// You can use the variables (with READ ONLY permissions): 
		// gomelipass.GetEnv("DB_MYSQL_DESAENV07_BGOW6S464_BGOW6S464_LOCAL_REPLICA_ENDPOINT"; to connect to the local replica, or... 
		// gomelipass.GetEnv("DB_MYSQL_DESAENV07_BGOW6S464_BGOW6S464_REMOTE_REPLICA_ENDPOINT"; for the remote replica 
		dbHost := gomelipass.GetEnv("DB_MYSQL_DESAENV07_BGOW6S464_BGOW6S464_ENDPOINT")
		dbName := "bgow6s464"
		connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", dbUsername, dbPassword, dbHost, dbName)
		db, err = sql.Open("mysql", connectionString)
		if err != nil {
			panic(err)
		}
		// Check that the database is available and accessible
		err = db.Ping()
		if err != nil {
			panic(err)
		}
	}

	// db, err := sql.Open("mysql", "root:@/melisprint")
	// if err != nil {
	// 	panic(err)
	// }

	eng := gin.Default()

	router := routes.NewRouter(eng, db)
	router.MapRoutes()

	docs.SwaggerInfo.Host = "localhost:8080"
	eng.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := eng.Run(); err != nil {
		panic(err)
	}
}

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/mercadolibre/fury_go-core/pkg/web"
// 	"github.com/mercadolibre/fury_go-platform/pkg/fury"
// )

// func main() {
// 	if err := run(); err != nil {
// 		log.Fatal(err)
// 	}
// }

// func run() error {
// 	app, err := fury.NewWebApplication()
// 	if err != nil {
// 		return err
// 	}

// 	app.Post("/hello", helloHandler)

// 	return app.Run()
// }

// func helloHandler(w http.ResponseWriter, r *http.Request) error {
// 	return web.EncodeJSON(w, fmt.Sprintf("%s, world!", r.URL.Path[1:]), http.StatusOK)
// }
