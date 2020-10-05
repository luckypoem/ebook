package main

import (
	"runtime"

	"github.com/luckypoem/ebook/httpd"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	http := httpd.Httpd{}
	http.Run("/etc/phoenix/config.json")
}
