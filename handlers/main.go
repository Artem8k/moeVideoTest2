package routes

import (
	"test/handlers/adCompanyRoutes"
	"test/handlers/advertisementRoutes"
	"test/handlers/advertiserRoutes"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var router = gin.Default()

func Run(db *sqlx.DB) {
	getRoutes(db)
	router.Run(":5000")
}

func getRoutes(db *sqlx.DB) {
	rg := router.Group("/api")

	advertiserRoutes.AddAdvertiserRoutes(rg, db)
	adCompanyRoutes.AddAdCompanyRoutes(rg, db)
	advertisementRoutes.AddAdvertisementRoutes(rg, db)
}
