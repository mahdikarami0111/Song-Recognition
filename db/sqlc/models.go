// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import ()

type Request struct {
	ID     int64  `json:"id"`
	Email  string `json:"email"`
	Status string `json:"status"`
	Songid string `json:"songid"`
}
