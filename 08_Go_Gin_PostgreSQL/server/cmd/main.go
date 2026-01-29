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
			"msg": "hi",
		})
		
	})

	r.GET("/slow", func (c*gin.Context)  {

		time.Sleep(5*time.Second)
		c.JSON(200, gin.H{
			"msg": "task done",
		})
		
	})


	// server start 
	srv := &http.Server{
		Addr: ":6065",
		Handler: r,
	}

	go func() {
		log.Printf("server started at port %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed{

			log.Fatalf("couldn't start server")
             return
		}
	}()

	// grace ful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<- quit

	log.Printf("System shutdown recieved💀")

	// ctx 
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatalf("couldn't shutdown server❌")
		return
	}

	log.Printf("server shutdown gracefull!✅")

}