package models

type Segment struct {
	Name string `json:"segment_name" binding:"required" db:"name"`
}
