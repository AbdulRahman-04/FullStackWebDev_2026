package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/config"
	publicAuth "github.com/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/controllers/public"
	"github.com/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/models"
	"github.com/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/utils"
)

func main() {

	// --------------------
	// LOAD CONFIG
	// --------------------
	config.LoadEnv()

	// --------------------
	// CONNECT DBs
	// --------------------
	utils.ConnectPostgres()
	utils.ConnectRedis()

	// START BACKGROUND WORKERS
	utils.StartEmailWorker()
	utils.StartSMSWorker()

	// --------------------
	// DB MIGRATIONS (TEMP)
	// --------------------
	if err := utils.PostgresDB.AutoMigrate(
		&models.User{},
		&models.Admin{},
	); err != nil {
		log.Fatalf("❌ DB migration failed: %v", err)
	}
	log.Println("✅ DB migration completed")

	// --------------------
	// CREATE SERVER
	// --------------------
	r := gin.Default()

	// ================================
	// PUBLIC AUTH ROUTES (TESTING)
	// ================================
	public := r.Group("/api/public")
	{
		// signup / login
		public.POST("/user/signup", publicAuth.UserSignup)
		public.POST("/user/signin", publicAuth.UserSignin)

		// verify
		public.GET("/user/emailverify/:token", publicAuth.EmailVerify)
		public.POST("/user/phoneverify", publicAuth.PhoneVerify)

		// refresh + forgot
		public.POST("/user/refresh", publicAuth.RefreshToken)
		public.POST("/user/forgot-password", publicAuth.ForgotPassword)
	}

	// --------------------
	// HEALTH CHECK
	// --------------------
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "Server started",
		})
	})

	// --------------------
	// HTTP SERVER
	// --------------------
	srv := &http.Server{
		Addr:    ":6065",
		Handler: r,
	}

	// --------------------
	// START SERVER
	// --------------------
	go func() {
		log.Printf("🚀 Server started at %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Server error: %v", err)
		}
	}()

	// --------------------
	// GRACEFUL SHUTDOWN
	// --------------------
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Println("❌ Server forced to shutdown")
	}

	log.Println("✅ Server gracefully stopped")
}
