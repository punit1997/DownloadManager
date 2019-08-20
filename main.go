package main

import (
	"github.com/punit1997/DownloadManager/routes"
)

func main() {
	r := routes.InitRoute()
	r.Run(":8081")
}
