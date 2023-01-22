package model

import "time"

type TodoId int

type Todo struct {
	Id     TodoId    `json:"id"`
	Title  string    `json:"title"`
	Detail string    `json:"detail"`
	Date   time.Time `json:"date"`
}
