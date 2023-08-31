package models

type User struct {
	Id int `form:"user_id" binding:"required" db:"id"`
}
