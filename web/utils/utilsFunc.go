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
	"net/http"
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
	//file, head, err := ctx.Request.FormFile(imgIndex)
	head, err := ctx.FormFile(imgIndex)
	if err != nil {
		utils.NewLog().Debug("ctx.Request.FormFile error:", err)
		return Response(RECODE_IOERR, nil), "", ""
	}
	file, _ := head.Open()
	utils.NewLog().Debugf("upload time %s s:", time.Since(start))
	//校验文件上传
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

func TimeParse(t string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02", t, time.Local)
}

//func TimeParse2(t string) (time.Time, error) {
//	return time.ParseInLocation("2006-01-02 15:04:05", t, time.Local)
//}

func CheckDays(ctx *gin.Context, startDay, endDay string) (int, error) {
	sd, err := TimeParse(startDay)
	if err != nil {
		utils.NewLog().Info("utils.TimeParse sd error:", err)
		ctx.JSON(http.StatusOK, Response(utils.RECODE_REQERR, nil))
		return 0, err
	}
	ed, err2 := TimeParse(endDay)
	if err2 != nil {
		utils.NewLog().Info("utils.TimeParse ed error:", err)
		ctx.JSON(http.StatusOK, Response(utils.RECODE_REQERR, nil))
		return 0, err2
	}
	if sd.After(ed) {
		utils.NewLog().Info("utils.After error:", err)
		ctx.JSON(http.StatusOK, Response(utils.RECODE_REQERR, nil))
		return 0, err2
	}
	sub := ed.Sub(sd)
	days := sub.Hours() / 24
	utils.NewLog().Debug("days:", days)
	return int(days), nil
}

func CheckDate(ctx *gin.Context, startDay, endDay string) bool {
	sd, err := TimeParse(startDay)
	if err != nil {
		utils.NewLog().Info("utils.TimeParse sd error:", err)
		ctx.JSON(http.StatusOK, Response(utils.RECODE_REQERR, nil))
		return false
	}
	ed, err2 := TimeParse(endDay)
	if err2 != nil {
		utils.NewLog().Info("utils.TimeParse ed error:", err)
		ctx.JSON(http.StatusOK, Response(utils.RECODE_REQERR, nil))
		return false
	}
	if sd.After(ed) {
		utils.NewLog().Info("utils.After error:", err)
		ctx.JSON(http.StatusOK, Response(utils.RECODE_REQERR, nil))
		return false
	}
	return true
}
