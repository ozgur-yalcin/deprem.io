package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ozgur-soft/deprem.io/cache"
	"github.com/ozgur-soft/deprem.io/controllers"
	"github.com/ozgur-soft/deprem.io/database"
)

func main() {
	router := gin.Default()
	router.GET("/", controllers.Anasayfa)
	router.GET("/iletisim/:id", controllers.Iletisim)
	router.GET("/iletisim", controllers.IletisimAra)
	router.POST("/iletisim", controllers.IletisimEkle)
	router.GET("/yardim/:id", controllers.Yardim)
	router.GET("/yardim", controllers.YardimAra)
	router.POST("/yardim", controllers.YardimEkle)
	router.GET("/yardimet/:id", controllers.Yardimet)
	router.GET("/yardimet", controllers.YardimetAra)
	router.POST("/yardimet", controllers.YardimetEkle)
	router.GET("/yardimet/rapor", controllers.YardimetRapor)
	router.GET("/yardimkaydi/:id", controllers.YardimKaydi)
	router.POST("/yardimkaydi", controllers.YardimKaydiEkle)
	router.GET("/cache/flushall", controllers.Flushall)
	database.Connect()
	cache.Connect()
	router.Run()
}
