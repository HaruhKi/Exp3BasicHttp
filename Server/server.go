package main

import (
	"net"
	"net/http"
	"os"
	"fmt"
)

func main() {

	http.HandleFunc("/hi", HiServer)
	http.ListenAndServe(":9000", nil)
}

func HiServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Hello", "BasicHTTP!")

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				//w.Write([]byte(ipnet.IP.String()))
				fmt.Fprintf(w, ipnet.IP.String())
			}
		}
	}

}
