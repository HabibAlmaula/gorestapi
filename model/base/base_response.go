package base

type BaseResponse struct {
	Code    int         `json:"code"`
	Succes  bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
