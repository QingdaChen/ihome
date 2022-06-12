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
	ID         int    `json:"user_id"`     //用户编号
	Name       string `json:"name"`        //用户名
	Mobile     string `json:"mobile" `     //手机号
	Real_name  string `json:"real_Name" `  //真实姓名  实名认证
	Id_card    string `json:"id_card" `    //身份证号  实名认证
	Avatar_url string `json:"avatar_url" ` //用户头像路径       通过fastdfs进行图片存储

}

type UpdateVo struct {
	Name string `json:"name"` //用户名
}
