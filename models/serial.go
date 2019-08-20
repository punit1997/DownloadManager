package models

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Serial struct {
	Urls []string
}

func (files *Serial) Download(r *gin.Context) {

	for _, url := range files.Urls {
		resp, _ := http.Get(url)
		nameId := uuid()
		out, _ := os.Create("/Users/punitlakshwani/go/src/github.com/punit1997/DownloadManager/downloads/" + nameId)
		io.Copy(out, resp.Body)
		resp.Body.Close()
	}

	ResponseId := uuid()
	r.JSON(200, gin.H{
		"Id": ResponseId,
	})
}
