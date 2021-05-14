package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"testing"
	"time"

	"golang.org/x/net/http2"

	rmq "github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

var jsonString = `
{
	"ID":10
}
`

func TestModule(t *testing.T) {
	f, err := os.OpenFile("./server.crt", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	fileByte, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(fileByte), string(fileByte))
	rmq.NewProducer(
		producer.WithNameServer([]string{"127.0.0.1"}),
		// producer.WithNsResolver(primitive.NewPassthroughResolver(endPoint)),
		producer.WithRetry(2),
		producer.WithGroupName("GID_xxxxxx"),
	)
}

func TestHttpClient(t *testing.T) {
	client := http.Client{
		// Skip TLS dial
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
		},
	}

	resp, err := client.Get("http://localhost:8972")
	if err != nil {
		t.Log(fmt.Errorf("error making request: %v", err))
	}
	time.Sleep(3 * time.Second)
	defer resp.Body.Close()
}

type serverHandler struct{}

func (sh *serverHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	w.Header().Set("server", "h2test")
	w.Write([]byte("this is a http2 test sever"))
}

func TestHttpServe(t *testing.T) {
	server := &http.Server{
		Addr:         ":8080",
		Handler:      &serverHandler{},
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	// http2.Server.ServeConn()
	s2 := &http2.Server{
		IdleTimeout: 1 * time.Minute,
	}
	http2.ConfigureServer(server, s2)
	l, _ := net.Listen("tcp", ":8080")
	defer l.Close()
	fmt.Println("Start server...")
	for {
		rwc, err := l.Accept()
		if err != nil {
			fmt.Println("accept err:", err)
			continue
		}
		s2.ServeConn(rwc, &http2.ServeConnOpts{BaseConfig: server})
	}
}

func TestHttpServe2(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello h2c")
	})
	s := &http.Server{
		Addr:    ":8880",
		Handler: mux,
	}
	http2.ConfigureServer(s, &http2.Server{})
	err := s.ListenAndServe()
	t.Log(err)
}
