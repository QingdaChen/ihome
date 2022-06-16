package model

import "github.com/jinzhu/gorm"

type HousePo struct {
	gorm.Model                    //房屋编号
	UserId          uint          `json:"user_id"`                  //房屋主人的用户编号  与用户进行关联
	AreaId          uint          `json:"area_id"`                  //归属地的区域编号   和地区表进行关联
	Title           string        `gorm:"size:64" json:"title"`     //房屋标题
	Address         string        `gorm:"size:512"json:"address"`   //地址
	Room_count      int           `gorm:"default:1"json:"address"`  //房间数目
	Acreage         int           `gorm:"default:0" json:"acreage"` //房屋总面积
	Price           int           `json:"price" json:"price"`
	Unit            string        `gorm:"size:32;default:''" json:"unit"`               //房屋单元,如 几室几厅
	Capacity        int           `gorm:"default:1" json:"capacity"`                    //房屋容纳的总人数
	Beds            string        `gorm:"size:64;default:''" json:"beds"`               //房屋床铺的配置
	Deposit         int           `gorm:"default:0" json:"deposit"`                     //押金
	Min_days        int           `gorm:"default:1" json:"min_days"`                    //最少入住的天数
	Max_days        int           `gorm:"default:0" json:"max_days"`                    //最多入住的天数 0表示不限制
	Order_count     int           `gorm:"default:0" json:"order_count"`                 //预定完成的该房屋的订单数
	Index_image_url string        `gorm:"size:256;default:''" json:"index_image_url"`   //房屋主图片路径
	Facilities      []*FacilityPo `gorm:"many2many:house_facilities" json:"facilities"` //房屋设施   与设施表进行关联
	Images          []*HouseImage `json:"img_urls"`                                     //房屋的图片   除主要图片之外的其他图片地址
	Orders          []*OrderHouse `json:"orders"`                                       //房屋的订单    与房屋表进行管理
}

type FacilityPo struct {
	Id   int    `json:"fid"`                //设施编号
	Name string `gorm:"size:32"json:"name"` //设施名字
}

type HouseFacPo struct {
	HouseId    int `json:"house_id"`
	FacilityId int `json:"facility_id"`
}
