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

	r:= gin.Default()

	r.GET("/", func (c*gin.Context)  {

		c.JSON(200, gin.H{
			"msg": "hi",
		})
		
	})

	r.GET("/slow", func (c*gin.Context)  {
		time.Sleep(6*time.Second)
		c.JSON(200, gin.H{
			"msg": "task done",
		})
		
	})

	// server start 
	srv := &http.Server{
		Addr: ":5545",
		Handler: r,
	}

	// go routine 
	go func() {
		log.Printf("server started at port %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed{
			log.Printf("err : %s", err)
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<- quit

	log.Printf("system shutdown recieved💀")

	// ctx 
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		fmt.Println(err)
	}

	log.Printf("server gracefully shutdown✅")
}