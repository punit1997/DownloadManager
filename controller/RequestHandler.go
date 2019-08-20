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
	} else {
		r.JSON(400, gin.H{
			"internal_code": 4001,
			"message":       "unknown type of download",
		})
		return
	}

	file.Download(r)
}

func Show(r *gin.Context) {
	response, ok := models.GetStatus[r.Param("id")]
	if !ok {
		r.JSON(400, gin.H{
			"internal_code": 4002,
			"message":       "unknown download ID",
		})
		return
	}
	r.JSON(200, response)
}
