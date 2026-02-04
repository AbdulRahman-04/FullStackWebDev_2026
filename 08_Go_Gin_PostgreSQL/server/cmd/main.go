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
 
	// make gin server 
	r := gin.Default()


	// routes 
	r.GET("/", func (c*gin.Context)  {

		c.JSON(200, gin.H{
			"msg": "Serer started",
		})
		
	})

	r.GET("/slow", func (c*gin.Context)  {
 
		time.Sleep(5*time.Second)

		c.JSON(200, gin.H{
			"msg": "task done",
		})
		
	})


	// create server 
	srv := &http.Server{
		Addr: ":6065",
		Handler: r,
	}

	// start server 
	go func() {
		log.Printf("Server started at port %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("err %s", err)
			return
		}
	}()

	// create channel
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<- quit

	log.Printf("system shutdown recieved⚠️")

	// create ctx
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		log.Printf("err %s", err)
		return
	}

	log.Printf("Shystem Gracefully Shutdown✅")

}