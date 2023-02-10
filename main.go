package main

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/ozgur-soft/deprem.io/controllers"
)

func certificate(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	server := strings.ReplaceAll(hello.ServerName, "www.", "")
	cert, err := tls.LoadX509KeyPair("/etc/letsencrypt/live/"+server+"/fullchain.pem", "/etc/letsencrypt/live/"+server+"/privkey.pem")
	if err != nil {
		log.Println(err)
		return nil, nil
	}
	return &cert, nil
}

func httpsHandler(w http.ResponseWriter, r *http.Request) {
	host, _, err := net.SplitHostPort(r.Host)
	if err != nil {
		host = strings.Split(r.Host, ":")[0]
	}
	http.Redirect(w, r, "https://"+strings.ToLower(host)+r.RequestURI, http.StatusFound)
}

func servehttp() error {
	server := &http.Server{Addr: ":http", ReadTimeout: 2 * time.Minute, WriteTimeout: 2 * time.Minute, IdleTimeout: 2 * time.Minute, Handler: http.HandlerFunc(httpsHandler)}
	server.SetKeepAlivesEnabled(true)
	return server.ListenAndServe()
}

func servehttps() error {
	config := &tls.Config{GetCertificate: certificate, PreferServerCipherSuites: true, MinVersion: tls.VersionTLS12, InsecureSkipVerify: true}
	server := &http.Server{Addr: ":https", ReadTimeout: 2 * time.Minute, WriteTimeout: 2 * time.Minute, IdleTimeout: 2 * time.Minute, TLSConfig: config}
	server.SetKeepAlivesEnabled(true)
	return server.ListenAndServeTLS("", "")
}

func run() error {
	err := make(chan error, 2)
	go func() { err <- servehttp() }()
	go func() { err <- servehttps() }()
	return <-err
}

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
	if e := run(); e != nil {
		log.Fatalln(e)
	}
}
