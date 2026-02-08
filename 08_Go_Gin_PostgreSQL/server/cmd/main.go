package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/config"
	"github.com/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/utils"
	"github.com/gin-gonic/gin"
)

func main(){

	// load config
	config.LoadEnv()

	// db connect
	utils.ConnectPostgres()
	utils.ConnectRedis()
 

	// create server gin 
	r := gin.Default()

	r.GET("/", func (c*gin.Context)  {

		c.JSON(200, gin.H{
			"msg": "Server started",
		})
		
	})

	r.GET("/slow", func (c*gin.Context)  {
		time.Sleep(5*time.Second)
		c.JSON(200, gin.H{
			"msg": "task done",
		})
	})


	// create server with port
	srv := &http.Server{
		Addr: ":6065",
		Handler: r,
	}

	// start server
	go func() {
		log.Printf("Server started at port %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Couldn't start server")
		}
	}()

	// make channel for gs
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<- quit

	log.Printf("System shutdown received")

	// ctx 
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		log.Printf("Couldn't shut down server")
	}

	log.Printf("System Gracefully shutdown✅ ")


}