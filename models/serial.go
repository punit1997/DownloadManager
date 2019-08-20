package models

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Serial struct {
	Urls []string
}

func (files *Serial) Download(r *gin.Context) {

	ResponseId := uuid()

	GetStatus[ResponseId] = &Status{
		Id:           ResponseId,
		StartTime:    time.Now(),
		Status:       "QUEUED",
		DownloadType: "SERIAL",
	}

	UrlLocation := make(map[string]string)

	for _, url := range files.Urls {
		resp, _ := http.Get(url)
		nameId := uuid()
		out, _ := os.Create("/Users/punitlakshwani/go/src/github.com/punit1997/DownloadManager/downloads/" + nameId)
		UrlLocation[url] = nameId
		io.Copy(out, resp.Body)
		resp.Body.Close()
	}

	GetStatus[ResponseId].EndTime = time.Now()
	GetStatus[ResponseId].Files = UrlLocation
	GetStatus[ResponseId].Status = "SUCCESSFULL"

	r.JSON(200, gin.H{
		"Id": ResponseId,
	})
}
