package elasticsearch

import "github.com/olivere/elastic/v7"

type ESHousePo struct {
	HouseId    int64  `json:"house_id"    mapstructure:"house_id"`
	CreateTime string `json:"ctime"       mapstructure:"ctime"`
	UserName   string `json:"user_name"   mapstructure:"user_name"`
	UserAvatar string `json:"user_avatar" mapstructure:"user_avatar"`
	UserId     int64  `json:"user_id"     mapstructure:"user_id" ` //房屋主人的用户编号  与用户进行关联
	AreaId     int64  `json:"area_id"     mapstructure:"area_id" ` //归属地的区域编号   和地区表进行关联
	AreaName   string `json:"area_name"   mapstructure:"area_name" `
	Title      string `json:"title"       mapstructure:"title"`      //房屋标题
	Address    string `json:"address"     mapstructure:"address"`    //地址
	RoomCount  int    `json:"room_count"  mapstructure:"room_count"` //房间数目
	Acreage    int    `json:"acreage"     mapstructure:"acreage"`    //房屋总面积
	Price      int    `json:"price"       mapstructure:"price"`
	Unit       string `json:"unit"        mapstructure:"unit"`        //房屋单元,如 几室几厅
	Capacity   int    `json:"capacity"    mapstructure:"capacity"`    //房屋容纳的总人数
	Beds       string `json:"beds"        mapstructure:"beds"`        //房屋床铺的配置
	Deposit    int    `json:"deposit"     mapstructure:"deposit"`     //押金
	MinDays    int    `json:"min_days"    mapstructure:"min_days"`    //最少入住的天数
	MaxDays    int    `json:"max_days"    mapstructure:"max_days"`    //最多入住的天数 0表示不限制
	OrderCount int    `json:"order_count" mapstructure:"order_count"` //预定完成的该房屋的订单数
	ImageUrl   string `json:"image_url"   mapstructure:"image_url"`   //房屋主图片路径
}

type ESSearch struct {
	MustQuery    []elastic.Query
	MustNotQuery []elastic.Query
	ShouldQuery  []elastic.Query
	RangeQuery   []elastic.RangeQuery
	Filters      []elastic.Query
	Sorters      []elastic.Sorter
	From         int //分页
	Size         int
}
