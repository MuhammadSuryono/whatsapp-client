package pagination

import "time"

type ResponseListLog struct {
	TotalData  int64       `json:"total_data"`
	Records    interface{} `json:"records"`
	LastUpdate time.Time   `json:"last_update"`
	TotalPage  int         `json:"total_page"`
}
