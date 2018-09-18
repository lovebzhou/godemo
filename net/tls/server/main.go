package main

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

func HandleClientConnect(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Receive Connect Request From ", conn.RemoteAddr().String())
	buffer := make([]byte, 1024)
	for {
		len, err := conn.Read(buffer)
		if err != nil {
			log.Println(err.Error())
			break
		}
		fmt.Printf("Receive Data: %s\n", string(buffer[:len]))
		//发送给客户端
		_, err = conn.Write([]byte("服务器收到数据:" + string(buffer[:len])))
		if err != nil {
			break
		}
	}
	fmt.Println("Client " + conn.RemoteAddr().String() + " Connection Closed.....")
}

func main() {
	crt, err := tls.LoadX509KeyPair("/etc/certs/server.crt", "/etc/certs/server.key")
	if err != nil {
		log.Fatalln(err.Error())
	}
	tlsConfig := &tls.Config{}
	tlsConfig.Certificates = []tls.Certificate{crt}
	// Time returns the current time as the number of seconds since the epoch.
	// If Time is nil, TLS uses time.Now.
	tlsConfig.Time = time.Now
	// Rand provides the source of entropy for nonces and RSA blinding.
	// If Rand is nil, TLS uses the cryptographic random reader in package
	// crypto/rand.
	// The Reader must be safe for use by multiple goroutines.
	tlsConfig.Rand = rand.Reader

	//
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world!")
	})
	go func() {
		// http.ListenAndServeTLS(":8443", "/etc/certs/server.crt", "/etc/certs/server.key", nil)

		server := http.Server{TLSConfig: tlsConfig}
		if l, err := net.Listen("tcp", ":8443"); err == nil {
			log.Println("start https server on", l.Addr().String())

			server.ServeTLS(l, "", "")
		}
	}()

	l, err := tls.Listen("tcp", ":5443", tlsConfig)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("start ssl socket server on", l.Addr().String())

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		} else {
			go HandleClientConnect(conn)
		}
	}
}
