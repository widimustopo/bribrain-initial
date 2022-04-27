package model

type AccumulationRKA struct {
	Name    string `json:"name" gorm:"column:name"`
	RkaName string `json:"rka_name" gorm:"column:rka_name"`
	Visited int64  `json:"visited" gorm:"column:visited"`
	Target  int64  `json:"target" gorm:"column:target"`
	Month   int32  `json:"month" gorm:"column:month"`
	Year    int32  `json:"year" gorm:"column:yaer"`
}
