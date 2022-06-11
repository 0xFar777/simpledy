package model

type Favorite struct {
	UserId  int64
	VideoId int64
	User    UserInformation `gorm:"ForeignKey:UserId"`
	Video   Video           `gorm:"ForeignKey:VideoId"`
}
