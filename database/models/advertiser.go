package models

import "time"

type Advertiser struct {
	Id          int       `json:"id" db:"id" form:"id" binding:"-"`
	Title       string    `json:"title" db:"title" form:"title" binding:"required"`
	Create_date time.Time `json:"create_date" db:"create_date" form:"create_date" binding:"-"`
}

type AdvertiserUpdate struct {
	Id    int    `json:"id" db:"id" form:"id" binding:"required"`
	Title string `json:"title" db:"title" form:"title" binding:"required"`
}
