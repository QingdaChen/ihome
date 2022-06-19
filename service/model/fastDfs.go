package model

import (
	"github.com/tedcy/fdfs_client"
	"ihome/service/conf"
	"ihome/service/user/kitex_gen"
	"ihome/service/utils"
)

type FastDfs struct {
	Client *fdfs_client.Client
	Err    error
}

var FastDfsClient *FastDfs

func init() {
	FastDfsClient = &FastDfs{}
	//fmt.Println(os.Getwd())
	client, err := fdfs_client.NewClientWithConfig(conf.FastDfsCfgPath)
	if err != nil {
		utils.NewLog().Error("fdfs_client init error", err)
	}
	FastDfsClient.Client = client
	FastDfsClient.Err = err
}

//UploadImg 上传图像
func (fastDfs *FastDfs) UploadImg(imgBase64 string, imgType string) kitex_gen.Response {
	buf, err := utils.Base64ToBuf(imgBase64)
	if err != nil {
		utils.NewLog().Error("Base64ToBuf error:", err)
		return utils.UserResponse(utils.RECODE_SERVERERR, nil)
	}
	if err != nil {
		utils.NewLog().Error("ExtensionsByType error:", err)
		return utils.UserResponse(utils.RECODE_SERVERERR, nil)
	}
	var filePath string
	utils.NewLog().Info("buf:  ", buf.Len())
	filePath, err = fastDfs.Client.UploadByBuffer(buf.Bytes(), imgType)
	if err != nil {
		utils.NewLog().Error("fastDfs.Client.UploadByBuffer error:", err, ":"+filePath)
		return utils.UserResponse(utils.RECODE_SERVERERR, nil)
	}
	//avatarUrl := fmt.Sprintf("%s/%s", conf.NginxUrl, filePath)
	return utils.UserResponse(utils.RECODE_OK, []byte(filePath))
}

//DeleteFile 删除文件
func (fastDfs *FastDfs) DeleteFile(fileName string) error {
	return fastDfs.Client.DeleteFile(fileName)
}
