package models

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Concurrent struct {
	Urls []string
}

const limitThreads = 5

func (files *Concurrent) Download(r *gin.Context) {

	ResponseId := uuid()

	GetStatus[ResponseId] = &Status{
		Id:           ResponseId,
		StartTime:    time.Now(),
		Status:       "QUEUED",
		DownloadType: "CONCURRENT",
	}

	var ch = make(chan string)
	UrlLocation := make(map[string]string)

	for i := 0; i < limitThreads; i++ {
		go func() {
			for {
				url, ok := <-ch
				if !ok {
					return //close go routine when channel is closed
				}
				resp, _ := http.Get(url)
				nameId := uuid()
				out, _ := os.Create("/Users/punitlakshwani/go/src/github.com/punit1997/DownloadManager/downloads/" + nameId)
				UrlLocation[url] = nameId
				io.Copy(out, resp.Body)
				resp.Body.Close()
			}
		}()
	}

	go func() {
		for _, url := range files.Urls {
			ch <- url
		}
		close(ch)

		GetStatus[ResponseId].EndTime = time.Now()
		GetStatus[ResponseId].Files = UrlLocation
		GetStatus[ResponseId].Status = "SUCCESSFULL"
		return
	}()

	r.JSON(200, gin.H{
		"Id": ResponseId,
	})
}
