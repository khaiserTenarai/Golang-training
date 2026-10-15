package model

type Department struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"size:100;not null;uniqueIndex" json:"name"`
	Location string `gorm:"size:100" json:"location"`
}
