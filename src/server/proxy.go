package server

import (
	"capture/src/config"
	"sync"
	"capture/src/mylog"
	"net/http"
	"time"
)

func Start(conf *config.Cfg, tlsConfig *config.TlsConfig, wg *sync.WaitGroup) {
	handler, err := InitConfig(conf, tlsConfig)
	if err != nil {
		mylog.Fatalf("InitConfig error: %s", err)
	}

	server := &http.Server{
		Addr:         ":" + *conf.Port,
		Handler:      handler,
		ReadTimeout:  1 * time.Hour,
		WriteTimeout: 1 * time.Hour,
	}

	go func() {
		mylog.Printf("Capture Listening On: %s", *conf.Port)
		if *conf.Tls {
			mylog.Println("Listen And Serve HTTP TLS")
			err = server.ListenAndServeTLS("gomitmproxy-ca-cert.pem", "gomitmproxy-ca-pk.pem")
		} else {
			mylog.Println("Listen And Serve HTTP")
			err = server.ListenAndServe()
		}
		if err != nil {
			mylog.Fatalf("Unable To Start HTTP proxy: %s", err)
		}

		wg.Done()
		mylog.Printf("Gomitmproxy Stop!!!!")
	}()

	return
}
