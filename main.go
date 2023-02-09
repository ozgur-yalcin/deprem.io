package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ozgur-soft/deprem.io/controllers"
)

// Sunucu bilgileri
const (
	httpHost = "0.0.0.0"
	httpPort = ":http" // ssl için ":https" kullanılmalıdır
)

func main() {
	http.HandleFunc("/", controllers.Anasayfa)
	http.HandleFunc("/iletisim", controllers.Iletisim)
	http.HandleFunc("/yardim", controllers.Yardim)
	http.HandleFunc("/yardimet", controllers.Yardimet)
	http.HandleFunc("/yardimet/export", controllers.YardimetExport)
	http.HandleFunc("/flushall", controllers.Flushall)
	http.HandleFunc("/getstats", controllers.GetStats)
	server := http.Server{Addr: httpHost + httpPort, ReadTimeout: 30 * time.Second, WriteTimeout: 30 * time.Second}
	// ssl için server.ListenAndServeTLS(".cert dosyası", ".key dosyası") kullanılmalıdır.
	if e := server.ListenAndServe(); e != nil {
		log.Fatalln(e)
	}
}
