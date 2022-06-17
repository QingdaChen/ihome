package model

type FacilityPo struct {
	Id   int    `json:"fid"`                //设施编号
	Name string `gorm:"size:32"json:"name"` //设施名字
}

type HouseFacPo struct {
	HouseId    int `json:"house_id"`
	FacilityId int `json:"facility_id"`
}
