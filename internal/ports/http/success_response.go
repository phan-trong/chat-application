package http

type SuccessResponse struct {
	Code int                    `json:"code"`
	Data map[string]interface{} `json:"data"`
}
