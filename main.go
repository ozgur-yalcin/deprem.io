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
	httpPort = ":9999" // ssl için ":https" kullanılmalıdır
)

func main() {
	http.HandleFunc("/", controllers.Anasayfa)
	http.HandleFunc("/iletisim", controllers.Iletisim)
	http.HandleFunc("/iletisim/ara", controllers.IletisimAra)
	http.HandleFunc("/iletisim/ekle", controllers.IletisimEkle)
	http.HandleFunc("/yardim", controllers.Yardim)
	http.HandleFunc("/yardim/ara", controllers.YardimAra)
	http.HandleFunc("/yardim/ekle", controllers.YardimEkle)
	http.HandleFunc("/yardimet", controllers.Yardimet)
	http.HandleFunc("/yardimet/ara", controllers.YardimetAra)
	http.HandleFunc("/yardimet/ekle", controllers.YardimetEkle)
	http.HandleFunc("/yardimet/rapor", controllers.YardimetRapor)
	http.HandleFunc("/flushall", controllers.Flushall)
	http.HandleFunc("/getstats", controllers.GetStats)
	server := http.Server{Addr: httpHost + httpPort, ReadTimeout: 30 * time.Second, WriteTimeout: 30 * time.Second}
	// ssl için server.ListenAndServeTLS(".cert dosyası", ".key dosyası") kullanılmalıdır.
	if e := server.ListenAndServe(); e != nil {
		log.Fatalln(e)
	}
}
