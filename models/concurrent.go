package models

import (
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
)

type Concurrent struct {
	Urls []string
}

const limitThreads = 1

func (files *Concurrent) Download(r *gin.Context) {

	var ch = make(chan string)

	var wg sync.WaitGroup
	wg.Add(limitThreads)

	for i := 0; i < limitThreads; i++ {
		go func() {
			for {
				url, ok := <-ch
				if !ok {
					wg.Done()
					return
				}
				resp, _ := http.Get(url)
				nameId := uuid()
				out, _ := os.Create("/Users/punitlakshwani/go/src/github.com/punit1997/DownloadManager/downloads/" + nameId)
				io.Copy(out, resp.Body)
				resp.Body.Close()
			}
		}()
	}

	for _, url := range files.Urls {
		ch <- url
	}

	close(ch)
	wg.Wait()
}
