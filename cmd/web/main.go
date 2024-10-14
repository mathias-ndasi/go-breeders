package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

const port = ":4000";

type application struct{
	templateMap map[string]*template.Template
	config appConfig
}

type appConfig struct {
	useCache bool
}

func main() {
	app := application{
		templateMap: make(map[string]*template.Template),
	};

	flag.BoolVar(&app.config.useCache, "cache", false, "Use template cache");
	flag.Parse();
	
	server := &http.Server {
		Addr:    port,
        Handler: app.routes(),
		IdleTimeout: 30 * time.Second,
		ReadTimeout: 30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	
	fmt.Println("Server is running on port ", port);

	error := server.ListenAndServe();
	if error!= nil {
        log.Panic(error);
    }
}