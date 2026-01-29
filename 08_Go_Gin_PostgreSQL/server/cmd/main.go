package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AbdulRahman-04/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/config"
	"github.com/AbdulRahman-04/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/utils"
	"github.com/gin-gonic/gin"
)

func main(){
	// config load
	config.LoadEnv()

	// database load
	utils.ConnectPostgres()
 
	// gin server 
	r:= gin.Default()

	r.GET("/", func (c*gin.Context)  {
		c.JSON(200, gin.H{
			"msg": "server started",
		})
		
	})

	r.GET("/slow", func (c*gin.Context)  {
		time.Sleep(5*time.Second)
		c.JSON(200, gin.H{
			"msg": "task done",
		})
		
	})

	srv := &http.Server{
		Addr: ":6065",
		Handler: r,
	}

	go func() { 
		log.Printf("server started at port %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
             log.Fatalf("err %s", err)
			 return
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<- quit

	log.Printf("System shutdown recieved💀")

	// ctx 
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		log.Printf("err %s", err)
		return
	}

	log.Printf("Server gracefully shutdown✅")


}