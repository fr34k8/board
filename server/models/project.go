package models

type Project struct {
	Id uint `json:"id"`
	Key string `json:"key"`
	Counter uint `json:"counter"`
	Title string `json:"title"`
}
