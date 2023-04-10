package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/cmd/api/handler"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/buyer"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/carry"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/employee"
	inboundorder "github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/inbound_order"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/locality"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/product"
	productbatches "github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/product_batches"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/product_records"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/purchase_orders"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/section"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/seller"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/warehouse"
)

type Router interface {
	MapRoutes()
}

type router struct {
	eng *gin.Engine
	rg  *gin.RouterGroup
	db  *sql.DB
	pr  *gin.RouterGroup
}

func NewRouter(eng *gin.Engine, db *sql.DB) Router {
	return &router{eng: eng, db: db}
}

func (r *router) MapRoutes() {
	r.setGroup()

	r.buildSellerRoutes()
	r.buildProductRoutes()
	r.buildSectionRoutes()
	r.buildProductBatchesRoutes()
	r.buildWarehouseRoutes()
	r.buildEmployeeRoutes()
	r.buildBuyerRoutes()
	r.buildPurchaseOrdersRoutes()
	r.buildInBoundOrder()
	r.buildProductRecordsRoutes()
	r.buildLocalityRoutes()
	r.buildCarryRoutes()
	r.buildHealthCheckRoute()
}

func (r *router) setGroup() {
	r.rg = r.eng.Group("/api/v1")
	r.pr = r.eng.Group("/api/v1/products")
}

func (r *router) buildHealthCheckRoute() {

	r.eng.GET("/ping", func(ctx *gin.Context) {
		w := ctx.Writer
		w.Write([]byte("pong"))
	})
}

func (r *router) buildSellerRoutes() {
	repo := seller.NewRepository(r.db)
	service := seller.NewService(repo)
	handler := handler.NewSeller(service)
	r.rg.GET("/sellers", handler.GetAll())
	r.rg.GET("/sellers/:id", handler.Get())
	r.rg.POST("/sellers", handler.Create())
	r.rg.PATCH("/sellers/:id", handler.Update())
	r.rg.DELETE("/sellers/:id", handler.Delete())
}

func (r *router) buildSectionRoutes() {
	repo := section.NewRepository(r.db)
	service := section.NewService(repo)
	handler := handler.NewSection(service)

	r.rg.GET("/sections", handler.GetAll())
	r.rg.GET("/sections/:id", handler.Get())
	r.rg.DELETE("/sections/:id", handler.Delete())
	r.rg.PATCH("/sections/:id", handler.Update())
	r.rg.POST("/sections", handler.Create())
}

func (r *router) buildProductBatchesRoutes() {
	repository := productbatches.NewRepository(r.db)
	service := productbatches.NewService(repository)
	handler := handler.NewProductBatches(service)

	r.rg.POST("/productbatches", handler.Create())
	r.rg.GET("/reportProducts/", handler.Get())
	
}

func (r *router) buildProductRoutes() {
	repo := product.NewRepository(r.db)
	service := product.NewService(repo)
	handler := handler.NewProduct(service)

	r.pr.GET("/", handler.GetAll())
	r.pr.GET("/:id", handler.Get())
	r.pr.POST("/", handler.Create())
	r.pr.PATCH("/:id", handler.Update())
	r.pr.DELETE("/:id", handler.Delete())
	r.pr.GET("/reportRecords", handler.GetReportRecords())
}

func (r *router) buildWarehouseRoutes() {
	repo := warehouse.NewRepository(r.db)
	service := warehouse.NewService(repo)
	handler := handler.NewWarehouse(service)
	r.rg.GET("/warehouses/:id", handler.Get())
	r.rg.GET("/warehouses", handler.GetAll())
	r.rg.POST("/warehouses", handler.Create())
	r.rg.PATCH("/warehouses/:id", handler.Update())
	r.rg.DELETE("/warehouses/:id", handler.Delete())
}

func (r *router) buildEmployeeRoutes() {
	repo := employee.NewRepository(r.db)
	service := employee.NewService(repo)
	handler := handler.NewEmployee(service)

	er := r.rg.Group("/employees")
	er.GET("", handler.GetAll())
	er.GET("/:id", handler.Get())
	er.POST("", handler.Create())
	er.DELETE("/:id", handler.Delete())
	er.PATCH("/:id", handler.Update())

	er.GET("/reportInboundOrders", handler.Report_InboundOrders())
}

func (r *router) buildBuyerRoutes() {
	// Example
	repo := buyer.NewRepository(r.db)
	service := buyer.NewService(repo)
	handler := handler.NewBuyer(service)

	pr := r.rg.Group("buyers")
	pr.POST("", handler.Create())
	pr.GET("", handler.GetAll())
	pr.GET("/:id", handler.Get())
	pr.DELETE("/:id", handler.Delete())
	pr.PATCH("/:id", handler.Update())
}

func (r *router) buildPurchaseOrdersRoutes() {
	repo := purchase_orders.NewRepository(r.db)
	service := purchase_orders.NewService(repo)
	handler := handler.NewPurchaseOrders(service)

	pr := r.rg.Group("purchaseOrders")
	pr.POST("", handler.Create())
	
  ps := r.rg.Group("buyers")
  ps.GET("reportPurchaseOrders", handler.Get())
}

func (r *router) buildInBoundOrder() {
	repo := inboundorder.NewRepository(r.db)
	service := inboundorder.NewService(repo)
	handler := handler.NewInBound_Order(service)

	bor := r.rg.Group("/inboundOrders")
	bor.GET("", handler.GetAll())
	bor.POST("", handler.Create())
  }
func (r *router) buildProductRecordsRoutes() {
	repo := product_records.NewRepository(r.db)
	service := product_records.NewService(repo)
	handler := handler.NewProductRecord(service)

	r.rg.POST("/productRecords", handler.Create())
}

func (r *router) buildLocalityRoutes() {
	repo := locality.NewRepository(r.db)
	service := locality.NewService(repo)
	handler := handler.NewLocality(service)

	r.rg.POST("/localities", handler.Create())
	r.rg.GET("/localities/reportSellers", handler.GetAllSellersByLocality())
	r.rg.GET("/localities/reportCarries", handler.GetReport())
}

func (r *router) buildCarryRoutes() {
	repo := carry.NewRepository(r.db)
	service := carry.NewService(repo)
	handler := handler.NewCarry(service)

	r.rg.POST("/carries", handler.Create())

}
