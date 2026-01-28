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
)

func main(){
	r := gin.Default()

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

	// server start
	srv := &http.Server{
		Addr: ":5455",
		Handler: r,
	}

	// go routine pe server
	go func() {
		log.Printf("server started at %s\n", srv.Addr)
		if err := srv.ListenAndServe(); err != nil  && err != http.ErrServerClosed {
			log.Fatalf("server start err %s\n", err)
		}
	}()

	// gracefful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<- quit

	log.Printf("system shutdown recieved💠")

	// ctx 
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		log.Printf("server shutdown err %s\n", err)
	}

	log.Printf("server gracefully shutdown✅")
	
}