package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main(){
 
	// create gin server 
	r := gin.Default()

	r.GET("/", func (c*gin.Context)  {
		c.JSON(200, gin.H{
			"msg": "Hi api",
		})
	})

	r.GET("/slow", func (c*gin.Context)  {

		time.Sleep(4*time.Second)
		c.JSON(200, gin.H{
			"msg": "task done",
		})
	})
 
	// start server 
	srv := &http.Server{
		Addr: ":6065",
		Handler: r,
	}

	// go routine pe start server
	go func() {
		log.Printf("server started at port %s\n", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("err: %s", err)
		}
	}()

	// graceful shurdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<- quit

	log.Printf("System going to shutdown soon💀")

	// ctx 
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		fmt.Println(err)
	}

	log.Printf("Server shutdown gracefully!✅")


}