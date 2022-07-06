package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"
)

var (
	name string
	port string

	Version string
)

func init() {
	flag.StringVar(&name, "name", os.Getenv("WHOAMI_NAME"), "give me a name")
	flag.StringVar(&port, "port", os.Getenv("WHOAMI_PORT"), "give me a port number")
}

func main() {
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/", whoami)
	if port == "" {
		port = "80"
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func healthz(w http.ResponseWriter, r *http.Request) {
	for k, vv := range r.Header {
		for _, v := range vv {
			w.Header().Add(k, v)
		}
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

func whoami(w http.ResponseWriter, r *http.Request) {
	u, _ := url.Parse(r.URL.String())
	wait := u.Query().Get("wait")
	if len(wait) > 0 {
		duration, err := time.ParseDuration(wait)
		if err == nil {
			time.Sleep(duration)
		}
	}

	if name != "" {
		_, _ = fmt.Fprintln(w, "Name:", name)
	}
	if Version != "" {
		_, _ = fmt.Fprintln(w, "Version:", Version)
	}

	hostname, _ := os.Hostname()
	_, _ = fmt.Fprintln(w, "Hostname:", hostname)

	ifaces, _ := net.Interfaces()
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			_, _ = fmt.Fprintln(w, "IP:", ip)
		}
	}
	_, _ = fmt.Fprintln(w, "RemoteAddr:", r.RemoteAddr)
	if err := r.Write(w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
