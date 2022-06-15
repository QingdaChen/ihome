package model

type UserPo struct {
	ID            int           `json:"id"`            //用户编号
	Name          string        `json:"name"`          //用户名
	Password_hash string        `json:"password_hash"` //用户密码加密的
	Mobile        string        `json:"mobile"`        //手机号
	Real_name     string        `json:"real_name"`     //真实姓名  实名认证
	Id_card       string        `json:"id_card" `      //身份证号  实名认证
	Avatar_url    string        `json:"avatar_url" `   //用户头像路径       通过fastdfs进行图片存储
	Houses        []*House      //用户发布的房屋信息  一个人多套房
	Orders        []*OrderHouse //用户下的订单       一个人多次订单
}
