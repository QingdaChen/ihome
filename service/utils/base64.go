package utils

import (
	"bytes"
	"encoding/base64"
)

func Base64ToBuf(imgBase64 string) (*bytes.Buffer, error) {
	data, err := base64.StdEncoding.DecodeString(imgBase64)
	if err != nil {
		NewLog().Error("base64 DecodeString error:", err)
		return nil, err
	}
	buf := bytes.NewBuffer(data)
	return buf, nil

}
