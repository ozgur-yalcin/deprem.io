package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ozgur-soft/deprem.io/controllers"
)

// Sunucu bilgileri
const (
	httpHost = "localhost"
	httpPort = ":http" // ssl için ":https" kullanılmalıdır
)

func main() {
	http.HandleFunc("/", controllers.Anasayfa)
	http.HandleFunc("/iletisim", controllers.Iletisim)
	http.HandleFunc("/yardim", controllers.Iletisim)
	http.HandleFunc("/yardimet", controllers.Iletisim)
	server := http.Server{Addr: httpHost + httpPort, ReadTimeout: 30 * time.Second, WriteTimeout: 30 * time.Second}
	// ssl için server.ListenAndServeTLS(".cert dosyası", ".key dosyası") kullanılmalıdır.
	if e := server.ListenAndServe(); e != nil {
		log.Fatalln(e)
	}
}
