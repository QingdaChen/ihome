package utils

import (
	"fmt"
	"github.com/tidwall/gjson"
	"strconv"
	"time"
)

func PareInt(str string) int {
	intRes, _ := strconv.ParseInt(str, 10, 32)
	return int(intRes)

}

func ConcatImgUrl(nginxDns string, imgUrl string) string {
	return fmt.Sprintf("%s/%s", nginxDns, imgUrl)
}

func ConcatRedisKey(str1 string, val interface{}) string {
	switch val.(type) {
	case int:
		return str1 + "_" + IntToString(val.(int))
	case string:
		return str1 + "_" + val.(string)
	case int64:
		return str1 + "_" + IntToString(val.(int64))
	}
	return ""
}

func IntToString(num interface{}) string {
	switch num.(type) {
	case int64:
		return strconv.FormatInt(num.(int64), 10)
	case int:
		return strconv.Itoa(num.(int))
	}
	return ""
}

func GetFromJson(json, key string) string {
	return gjson.Get(json, key).String()
}

func TimeParse(t string) (time.Time, error) {
	//return time.ParseInLocation("2006-01-02 15:04:05", t, time.Local)
	return time.ParseInLocation("2006-01-02", t, time.Local)
}

//GetDays 输入日期 yy-mm-dd HH:mm:ss
func GetDays(startDay, endDay string) (int, error) {
	sd, err := TimeParse(startDay)
	if err != nil {
		NewLog().Info("utils.TimeParse sd error:", err)
		return 0, err
	}
	ed, err2 := TimeParse(endDay)
	if err2 != nil {
		NewLog().Info("utils.TimeParse ed error:", err)
		return 0, err2
	}
	if sd.After(ed) {
		NewLog().Info("utils.After error:", err)
		return 0, err2
	}
	sub := ed.Sub(sd)
	days := sub.Hours() / 24
	NewLog().Debug("days:", days)
	return int(days), nil
}
