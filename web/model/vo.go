package model

type RegisterRequest struct {
	Phone    string `json:"mobile"`
	Password string `json:"password"`
	SmsCode  string `json:"sms_code"`
}

type LoginRequest struct {
	Phone    string `json:"mobile"`
	Password string `json:"password"`
}

type SessionVo struct {
	Name string `json:"name"` //用户名
}

type UserVo struct {
	ID         int    `json:"id" map:"id"`                 //用户编号
	Name       string `json:"name" map:"name"`             //用户名
	Mobile     string `json:"mobile" map:"mobile"`         //手机号
	Real_name  string `json:"real_name" map:"real_name"`   //真实姓名  实名认证
	Id_card    string `json:"id_card" map:"id_card"`       //身份证号  实名认证
	Avatar_url string `json:"avatar_url" map:"avatar_url"` //用户头像路径       通过fastdfs进行图片存储

}

type UpdateVo struct {
	Name string `json:"name"` //用户名

}

type HouseVo struct {
	Houses []House `json:"houses" map:"id"`
}

type UserHouseVo struct {
	Ctime     string `map:"ctime" json:"ctime"`
	HouseId   uint   `map:"house_id" json:"house_id"`
	UserId    uint   `map:"user_id" json:"user_id"`       //房屋主人的用户编号  与用户进行关联
	AreaName  string `map:"area_name" json:"area_name"`   //归属地的区域编号   和地区表进行关联
	Title     string `json:"title" map:"title"`           //房屋标题
	Address   string `json:"address" map:"address"`       //地址
	RoomCount int    `json:"room_count" map:"room_count"` //房间数目
	//Acreage       int    `json:"acreage" map:"acreage"`       //房屋总面积
	Price int `json:"price" map:"price"`
	//Unit          string `json:"unit" map:"unit"`                       //房屋单元,如 几室几厅
	//Capacity      int    `json:"capacity" map:"capacity"`               //房屋容纳的总人数
	//Beds          string `json:"beds" map:"beds"`                       //房屋床铺的配置
	//Deposit       int    `json:"deposit" map:"deposit"`                 //押金
	//MinDays       int    `json:"min_days" map:"min_days"`               //最少入住的天数
	//MaxDays       int    `json:"max_days" map:"max_days"`               //最多入住的天数 0表示不限制
	OrderCount int    `json:"order_count" map:"order_count"` //预定完成的该房屋的订单数
	ImageUrl   string `json:"image_url" map:"image_url"`     //房屋主图片路径
	UserAvatar string `json:"user_avatar" map:"user_avatar"`
	//Facilities      []*Facility   `gorm:"many2many:house_facilities" json:"facilities"`                     //房屋设施   与设施表进行关联
}

type PostOrderReq struct {
	HouseId   string `json:"house_id" binding:"required"`
	StartDate string `json:"start_date" binding:"required"`
	EndDate   string `json:"end_date" binding:"required"`
}
