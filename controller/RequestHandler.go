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

func Show(r *gin.Context) {

	response := models.GetStatus[r.Param("id")]

	r.JSON(200, gin.H{
		"id":            response.Id,
		"start_time":    response.StartTime,
		"end_time":      response.EndTime,
		"status":        response.Status,
		"download_type": response.DownloadType,
		"files":         response.Files,
	})
}
