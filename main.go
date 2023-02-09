package main

import (
	"log"
	"net/http"
	"time"
)

// Sunucu bilgileri
const (
	httpHost = "localhost"
	httpPort = ":http" // ssl için ":https" kullanılmalıdır
)

func main() {
	http.HandleFunc("/", view)
	server := http.Server{Addr: httpHost + httpPort, ReadTimeout: 30 * time.Second, WriteTimeout: 30 * time.Second}
	// ssl için server.ListenAndServeTLS(".cert dosyası", ".key dosyası") kullanılmalıdır.
	if e := server.ListenAndServe(); e != nil {
		log.Fatalln(e)
	}
}

func view(w http.ResponseWriter, r *http.Request) {
}
