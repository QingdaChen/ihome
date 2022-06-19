package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"ihome/service/utils"
	"ihome/web/conf"
	"io"
	"strconv"
	"strings"
	"time"
)

func PareInt(str string) int {
	intRes, _ := strconv.ParseInt(string(str), 10, 32)
	return int(intRes)

}

func ConcatImgUrl(nginxDns string, imgUrl string) string {
	return fmt.Sprintf("%s/%s", nginxDns, imgUrl)
}

func ConcatRedisKey(str1, str2 string) string {
	return str1 + "_" + str2
}

func IntToString(num int) string {
	return strconv.Itoa(num)
}

func GetFromJson(json, key string) string {
	return gjson.Get(json, key).String()
}

func CheckFile(ctx *gin.Context, imgIndex string) (resp map[string]interface{}, fileTYpe string, imgBase64 string) {
	start := time.Now()
	file, head, err := ctx.Request.FormFile(imgIndex)
	//head, err := ctx.FormFile(imgIndex)
	//file, _ := head.Open()
	utils.NewLog().Debugf("upload time %s s:", time.Since(start))
	//校验文件上传
	if err != nil {
		utils.NewLog().Debug("ctx.Request.FormFile error:", err)
		return Response(RECODE_IOERR, nil), "", ""
	}
	utils.NewLog().Debug("house_image....")

	fileType := strings.Split(head.Filename, ".")[1]
	utils.NewLog().Debug("fileType: ", fileType)
	if fileType != "jpg" && fileType != "png" && fileType != "jpeg" {
		utils.NewLog().Debug("fileTYpe error:", fileType)
		return Response(RECODE_IOERR, nil), "", ""
	}
	if head.Size > conf.UploadFileMaxSize {
		utils.NewLog().Debug("file size error:")
		return Response(RECODE_SERVERERR, nil), "", ""
	}
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, file)
	//base64编码
	imageBase64 := base64.StdEncoding.EncodeToString(buf.Bytes())
	return Response(RECODE_OK, nil), fileType, imageBase64
}
