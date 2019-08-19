package models

import (
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Serial struct {
	Urls []string
}

func uuid() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return string(uuid)
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
