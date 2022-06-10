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
