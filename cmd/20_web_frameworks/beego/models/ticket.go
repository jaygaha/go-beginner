package models

import "time"

type Ticket struct {
	Id          int        `json:"id" orm:"auto"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedBy   string     `json:"created_by"`
	CreatedAt   *time.Time `json:"created_at,omitempty" orm:"auto_now_add;type(datetime)"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" orm:"auto_now;type(datetime)"`
}
