package remote

import (
	"errors"
	"ihome/web/conf"
	"ihome/web/utils"
)

func GetService(serviceName string) interface{} {
	var service interface{}
	var err error
	utils.NewLog().Info("serviceName:", serviceName)
	switch serviceName {
	//case conf.CaptchaServiceIndex:
	//	service, err = CaptchaService.Service, CaptchaService.Err
	case conf.UserServiceIndex:
		service, err = UserService.Service, UserService.Err
		utils.NewLog().Info("init ...:", service, err)
	//case conf.HouseServiceIndex:
	//	service, err = HouseService.Service, HouseService.Err
	default:
		utils.NewLog().Info("default")
		service = nil
		err = errors.New("no service error")
	}
	if err != nil {
		utils.NewLog().Error("get service error:", err)
		//ctx.JSON(http.StatusOK, Response(RECODE_SERVERERR, nil))
		return nil

	}
	return service

}
