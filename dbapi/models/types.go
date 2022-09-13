package models

type CFG struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
