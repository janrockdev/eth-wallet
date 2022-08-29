package models

type Rule struct {
	ID        string `json:"id"`
	Rule      string `json:"rule"`
	Clearance string `json:"clearance"`
}

type Token struct {
	ID        string `json:"id"`
	Token     string `json:"token"`
	Clearance string `json:"clearance"`
}

type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
