package main

import (
	"github.com/JuXue54/filecoinAPI/controller"
	"net/http"
	"context"
	"log"
	"os"
	"os/signal"
	"time"
)

func main(){
	mux:=http.NewServeMux()

	mux.HandleFunc("/api/pool/filecoin", controller.mainnet())

	timeOut:=time.Second*60
	srv:=&http.Server{
		Addr:	":8080",
		Handler: mux,
		ReadTimeout: timeOut,
		WriteTimeout: timeOut,
		IdleTimeout: timeOut*2,
	}

	go func(){
		if err:=srv.ListenAndServe(); err!=nil{
			log.Fatalf("listen and serve http server fail:\n%v", err)
		}
	}()

	exit:=make(chan os.Signal)
	signal.Notify(exit,os.Interrupt)
	<-exit
	ctx,cacel:=context.WithTimeout(context.Background(),timeOut)
	defer cacel()
	err:=srv.Shutdown(ctx)
	log.Println("shutting down now. ",err)
	os.Exit(0)
}