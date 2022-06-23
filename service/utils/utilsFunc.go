package utils

import (
	"fmt"
	"github.com/tidwall/gjson"
	"strconv"
)

func PareInt(str string) int {
	intRes, _ := strconv.ParseInt(string(str), 10, 32)
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
