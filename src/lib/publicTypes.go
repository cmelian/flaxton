package lib

import (
	"encoding/json"
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
	StartTime int64 	`json:"start_time,omitempty"`
	EndTime int64 		`json:"end_time,omitempty"`
}
type TaskResult struct  {
	TaskID string       `json:"task_id"`
	StartTime int64 	`json:"start_time,omitempty"`
	EndTime int64 		`json:"end_time,omitempty"`
	Error bool          `json:"error"`
	Message string      `json:"message"`
}

func (t *Task) ConvertData(obj interface{}) error {
	b, err := json.Marshal(t.Data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, obj)
	if err != nil {
		return err
	}
	return nil
}