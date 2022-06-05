package main

import (
	kitex_gen "ihome/web/user/kitex_gen/smsservice"
	"log"
)

func main() {
	svr := kitex_gen.NewServer(new(SMSServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
