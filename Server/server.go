package main

import (
	"fmt"
	"net"
	"net/http"
)

func main() {

	http.HandleFunc("/hi", HiServer)
	http.ListenAndServe(":8080", nil)
}

func HiServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Hello", "BasicHTTP!")

	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	userIP := net.ParseIP(ip)
	if userIP == nil {
		//return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
		fmt.Fprintf(w, "userip: %q is not IP:port", r.RemoteAddr)
		return
	}

	w.Write([]byte(net.ParseIP(ip)))
}
