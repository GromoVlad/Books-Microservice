package baseResponse

type BaseResponse struct {
	Data    any  `json:"data"`
	Success bool `json:"success"`
}

type Response struct {
	Status string `json:"status"`
}
