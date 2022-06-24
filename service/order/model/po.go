package model

type FacilityPo struct {
	Id   int    `json:"fid"`                //设施编号
	Name string `gorm:"size:32"json:"name"` //设施名字
}

type HouseFacPo struct {
	HouseId    int `json:"house_id"`
	FacilityId int `json:"facility_id"`
}

//CommentPo 评论
type CommentPo struct {
	Comment  string `json:"comment"`
	Ctime    string `json:"ctime"`
	UserName string `json:"user_name"`
}

//type HouseDetailInfo struct {
//	Address   string
//	areaName  string
//	Ctime     string
//	HouseId   int
//	Price     int
//	RoomCount int
//	Title     string
//	UserId    int
//	Acreage   string
//	Beds      string
//	Capacity  int
//}
