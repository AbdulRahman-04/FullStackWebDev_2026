package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AbdulRahman-04/FullStackWebDev_2026/09Backend_Practice/server/internal/config"
	"github.com/gin-gonic/gin"
)


func main() {

	// load config
	config.LoadEnv()

	// initialise gin server 
	r := gin.Default()

	// basic route 
	r.GET("/", func (c*gin.Context)  {
		c.JSON(200, gin.H{
			"msg": "server started",
		})
	})

	r.GET("/slow", func (c*gin.Context)  {
		time.Sleep(5*time.Second)
		c.JSON(200, gin.H{
			"msg": "task done✅",
		})
	})

	// create server and give it a port address
	srv := &http.Server{
		Addr: ":"+config.AppConfig.Port,
		Handler: r,
	}

	// start the server 
	go func() {
		log.Printf("Server started at port %s", srv.Addr)
		err := srv.ListenAndServe();
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Couldn't start server, err : %v", err)
		}
	}()

	// create graceful shut down 
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<- quit

	log.Printf("Graceful shutdown received⚠️")
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		log.Printf("err %v", err)
	}

	log.Printf("System gracefully shutdown✅")


}