package elasticsearch

type HouseSearchReq struct {
	HouseId    int64  `json:"house_id"    mapstructure:"house_id"`
	CreateTime string `json:"ctime"       mapstructure:"ctime"`
	UserName   string `json:"user_name"   mapstructure:"user_name"`
	UserId     int64  `json:"user_id"     mapstructure:"user_id" `   //房屋主人的用户编号  与用户进行关联
	AreaId     int64  `json:"area_id"     mapstructure:"area_id" `   //归属地的区域编号   和地区表进行关联
	Title      string `json:"title"       mapstructure:"title"`      //房屋标题
	Address    string `json:"address"     mapstructure:"address"`    //地址
	RoomCount  int    `json:"room_count"  mapstructure:"room_count"` //房间数目
	Acreage    int    `json:"acreage"     mapstructure:"acreage"`    //房屋总面积
	MinPrice   int    `json:"min_price"   mapstructure:"min_price"`
	MaxPrice   int    `json:"max_price"   mapstructure:"max_price"`
	Unit       string `json:"unit"        mapstructure:"unit"`        //房屋单元,如 几室几厅
	Capacity   int    `json:"capacity"    mapstructure:"capacity"`    //房屋容纳的总人数
	Beds       string `json:"beds"        mapstructure:"beds"`        //房屋床铺的配置
	Deposit    int    `json:"deposit"     mapstructure:"deposit"`     //押金
	Days       int    `json:"days"        mapstructure:"days"`        //最少入住的天数
	OrderCount int    `json:"order_count" mapstructure:"order_count"` //预定完成的该房屋的订单数
	Page       int    `json:"page"        mapstructure:"page"`        //页码
	PageSize   int    `json:"size"        mapstructure:"size"`        //每一页数量
}
