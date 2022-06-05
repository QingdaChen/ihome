package main

import (
	"context"
	"encoding/json"
	"github.com/afocus/captcha"
	"ihome/service/captcha/kitex_gen"
	"ihome/service/captcha/model"
	"ihome/service/captcha/utils"
	"image/color"
)

// CaptchaServiceImpl implements the last service interface defined in the IDL.
type CaptchaServiceImpl struct{}

// GetCaptcha implements the CaptchaServiceImpl interface.
func (s *CaptchaServiceImpl) GetCaptcha(ctx context.Context, req *kitex_gen.Request) (resp *kitex_gen.Response, err error) {
	utils.NewLog().Info("start...", err)
	cap := captcha.New()
	//cap.SetFont("")
	if err := cap.SetFont("./conf/comic.ttf"); err != nil {
		utils.NewLog().Error("cap.SetFont error...", err)
		panic(err.Error())
	}

	cap.SetSize(128, 64)
	cap.SetDisturbance(captcha.MEDIUM)
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})
	img, str := cap.Create(5, captcha.NUM)
	codeImg, _ := json.Marshal(img)
	utils.NewLog().Info("end...", err)
	response := &kitex_gen.Response{}
	response.Img = codeImg

	utils.NewLog().Info(str)
	err = model.SaveImgCode(req.Uuid, str)
	if err != nil {
		utils.NewLog().Error("SaveImgCode error...", err)
	}

	return response, nil
}
