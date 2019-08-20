package models

import (
	"time"
)

type Status struct {
	Id           string
	StartTime    time.Time
	EndTime      time.Time
	Status       string
	DownloadType string
	Files        map[string]string
}

var GetStatus = make(map[string]*Status)
