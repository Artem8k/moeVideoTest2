package models

import "time"

type AdCompany struct {
	Id            int       `json:"id" db:"id" form:"id" binding:"-"`
	Title         string    `json:"title" db:"title" form:"title" binding:"required"`
	Create_date   time.Time `json:"create_date" db:"create_date" form:"create_date" binding:"-"`
	Advertiser_id int       `json:"advertiser_id" db:"advertiser_id" form:"advertiser_id" binding:"required"`
}

type AdCompanyUpdate struct {
	Id    int    `json:"id" db:"id" form:"id" binding:"required"`
	Title string `json:"title" db:"title" form:"title" binding:"required"`
}
