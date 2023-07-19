package models

type PaginationResponse struct {
	Code      int64       `json:"code"`
	RowsCount int64       `json:"rows_count"`
	Limit     int64       `json:"limit"`
	Page      int64       `json:"page"`
	Data      interface{} `json:"data"`
}
