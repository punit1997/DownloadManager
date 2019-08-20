package controller

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/punit1997/DownloadManager/models"
)

type TypeDownload interface {
	Download(r *gin.Context)
}

func Start(r *gin.Context) {
	body := r.Request.Body
	bodyByte, _ := ioutil.ReadAll(body)
	var request models.Request
	json.Unmarshal(bodyByte, &request)

	var file TypeDownload

	if request.Type == "serial" {
		file = &models.Serial{Urls: request.Urls}
	} else if request.Type == "concurrent" {
		file = &models.Concurrent{Urls: request.Urls}
	}

	file.Download(r)
}
