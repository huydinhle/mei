package main

import (
	"encoding/base64"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		if checkAuth(w, r) {
			promhttp.Handler().ServeHTTP(w, r)
			return
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="MY REALM"`)
		w.WriteHeader(401)
		w.Write([]byte("401 Unauthorized\n"))
	})

	http.Handle("/metrics1", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, time.Now().String())
	fmt.Fprintf(w, "\nAppVersion:"+os.Getenv("VERSION"))
}

func checkAuth(w http.ResponseWriter, r *http.Request) bool {
	s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 {
		return false
	}

	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		return false
	}

	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		return false
	}

	return pair[0] == "user" && pair[1] == "pass"
}
