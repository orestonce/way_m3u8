package model

import "time"

type Work_D struct {
	ID         string    `json:"ID"`
	Name       string    `json:"name"`
	Url        string    `json:"url"`
	SaveDir    string    `json:"save_dir"`
	State      int       `json:"state"`
	Info       string    `json:"info"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
