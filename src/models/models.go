package models

type Response struct {
	Echo      string `json:"echo"`
	Timestamp string `json:"timestamp"`
	Env       string `json:"Environment"`
	Version   string `json:"buildversion"`
}
