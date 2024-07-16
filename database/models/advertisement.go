package models

import "time"

type Advertisement struct {
	Id           int       `json:"id" db:"id" form:"id" binding:"-"`
	Title        string    `json:"title" db:"title" form:"title" binding:"required"`
	CPM          int       `json:"cpm" db:"cpm" form:"CPM" binding:"required"`
	Create_date  time.Time `json:"create_date" db:"create_date" form:"create_date" binding:"-"`
	AdCompany_id int       `json:"adCompany_id" db:"adcompany_id" form:"adCompany_id" binding:"required"`
}

type AdvertisementUpdate struct {
	Id    int    `json:"id" db:"id" form:"id" binding:"required"`
	Title string `json:"title" db:"title" form:"title" binding:"required"`
	CPM   int    `json:"CPM" db:"CPM" form:"CPM" binding:"required"`
}
