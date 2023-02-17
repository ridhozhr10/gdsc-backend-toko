package engine

import (
	"github.com/gin-gonic/gin"
	"github.com/ridhozhr10/gdsc-backend-toko/entities"
	"github.com/ridhozhr10/gdsc-backend-toko/handler"
	"github.com/ridhozhr10/gdsc-backend-toko/repository/psql"
	"github.com/ridhozhr10/gdsc-backend-toko/service/product"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Run(port string) error {
	e := gin.Default()

	// setup db
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	db.AutoMigrate(entities.Product{})

	// repository setup
	productRepository := psql.NewProductPSQLRepo(db)

	// service setup
	productService := product.NewProductService(productRepository)

	// handler setup
	handler := handler.NewGinHttp(productService)

	// setup route
	e.GET("/product", handler.GetAllProduct)
	e.GET("/product/:id", handler.GetOneProduct)
	e.POST("/product/", handler.CreateProduct)
	e.PATCH("/product/:id", handler.UpdateProduct)
	e.DELETE("/product/:id", handler.DeleteProduct)

	return e.Run(port)
}
