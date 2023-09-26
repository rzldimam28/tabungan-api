package dto

type SuccessResponse struct {
	Success bool `json:"success"`
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Success bool `json:"success"`
	Remark Remark `json:"remark"`
}

type Remark struct {
	ErrMessage string `json:"err_message"`
}