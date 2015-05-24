package lib

import (
	"time"
)

type TransferContainerCall struct {
	Name 		string          `json:"name"`
	Cmd 		string 			`json:"cmd"`
	ImageName 	string			`json:"image_name"`
	ImageId 	string          `json:"image_id"`
	NeedToRun	bool        	`json:"need_to_run"`
	Authorization string 		`json:"authorization"`
	Destination string        	`json:"destination"`  // Container Name or IP address or ID
}

type TransferResponse struct {
	Error bool 		`json:"error"`
	Message string 	`json:"message"`
	Done bool 		`json:"done"`
	TaskId string   `json:"task_id"`
}

// TaskStack types
const (
	TaskContainerTransfer = "container_transfer"
)

type TaskType string

type Task struct  {
	TaskID string       `json:"task_id"`
	Type TaskType  		`json:"type"`
	Data interface{} 	`json:"data"`
	Cron bool 			`json:"cron"`
	StartTime time.Time `json:"start_time"`
	EndTime time.Time 	`json:"end_time"`
}
type TaskResult struct  {
	TaskID string       `json:"task_id"`
	StartTime time.Time `json:"start_time"`
	EndTime time.Time 	`json:"end_time"`
	Error bool          `json:"error"`
	Message string      `json:"message"`
}