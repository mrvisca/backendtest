package routes

import (
	"backendtest/controllers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiRoute() {
	// Memanggil fungsi route dari framework gin golang
	router := gin.Default()

	// // Menambahkan cors pada settingan route gin golang
	// router.Use(cors.Default())

	// // Menggunakan middleware CORS
	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:8080", "http://example.com"}, // Ganti dengan origin yang diizinkan
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// }))

	// Route API Endpoint
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello Go developers")
	})

	router.GET("/language", controllers.GetLanguange)
	router.GET("/language/:id", controllers.GetByIdLanguage)
	router.POST("/language", controllers.PostLanguage)
	router.PATCH("/language/:id", controllers.UpdateLanguage)
	router.DELETE("/language/:id", controllers.DeleteLanguage)

	// Menangani metode atau endpoint lainnya
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
	})
	router.GET("/check-palindrome", controllers.CheckPalindrome)

	// Menampilkan log server berjalan dengan port 8080
	log.Println("Server started on: http://127.0.0.1:8080")

	// Menjalankan server ke port 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
