package models

type Response struct {
	Method        string              `json:"method"`
	Headers       map[string][]string `json:"headers"`
	Payload       string              `json:"payload"`
	RemoteAddress string              `json:"remoteAddress"`
	Protocol      string              `json:"protocol"`
	ContentLength int64               `json:"contentLength"`
}
