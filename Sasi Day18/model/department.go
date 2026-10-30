package model

import 
	"time"

type Department struct {
	ID               int64 	   `json:"id"`
	Name             string    `json:"name"`
	Code             string    `json:"code"`
	HeadcountLimit   int       `json:"headcount_limit"`
	CurrentHeadcount int       `json:"current_headcount"`
	CreatedAt        time.Time `json:"created_at"`
}