package main

import (
	"flag"
	"github.com/gorilla/handlers"
	"log"
	"net"
	"net/http"
	"os"
	"path"
)

func main() {
	port := flag.String("port", "8000", "port to listen on")

	flag.Parse()

	shares := flag.Args()

	if len(shares) == 0 {
		shares = []string{"./"}
	}

	for _, dir := range shares {
		dir := path.Clean(dir)
		base := "/" + path.Base(dir)

		if base == "/." {
			base = ""
		}

		http.Handle(base+"/", http.StripPrefix(base, handlers.CombinedLoggingHandler(os.Stdout, http.FileServer(http.Dir(dir)))))
	}

	log.Fatal(http.ListenAndServe(net.JoinHostPort("", *port), nil))
}
