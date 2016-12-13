package main

import (
	"time"
)

type Todo struct {
	Id          int       `json:"id"`
	CreatedById int       `json:"userid"`
	Name        string    `json:"name"`
	Completed   bool      `json:"completed"`
	Due         time.Time `json:"due"`
}
