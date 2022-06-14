package handler

import (
	"encoding/json"
	"ihome/service/user/conf"
	"ihome/service/user/kitex_gen"
	"ihome/service/user/model"
	"ihome/service/utils"
)

func GetUserInfoByHandler(sessionId string) kitex_gen.Response {
	//TODO 限流
	utils.NewLog().Debug("GetUserInfo start")
	//加分布式锁,避免高并发下同时对大量对数据库用户信息的访问
	//model.InitLock(conf.RedisLockKey + "-user-" + sessionId)
	//model.Lock.Lock()
	//defer model.TryRelease()
	//先查redis
	redisResp := model.GetRedisUserInfo(sessionId)
	if utils.RECODE_OK == redisResp.Errno {
		//redis查到了直接返回
		utils.NewLog().Info("GetRedisUserInfo success!")
		return redisResp
	}
	//查不到查数据库
	mysqlResp := model.GetUserInfo(sessionId)
	if utils.RECODE_OK != mysqlResp.Errno {
		//失败了直接返回
		utils.NewLog().Info("mysql GetUserInfo failed:", mysqlResp)
		return mysqlResp
	}
	//写入redis
	saveResp := model.SaveRedisUserInfo(sessionId, mysqlResp.Data)
	if utils.RECODE_OK != saveResp.Errno {
		//保存失败直接返回
		utils.NewLog().Error("redis save error:", saveResp)

	}
	return mysqlResp
}

func DeleteFileByHandler(data []byte) {
	utils.AntsPool.Pool.Submit(func() {
		user := &model.User{}
		json.Unmarshal(data, user)
		if user.Avatar_url != "" {
			//已经上传过头像了,就启动协程删除
			utils.NewLog().Debug("DeleteFileByHandler....")
			model.FastDfsClient.Client.DeleteFile(user.Avatar_url)
		}
	})
}

func UpdateUserInfoByHandler(sessionID string, key, value string) kitex_gen.Response {
	//先删除缓存
	delResp := model.DeleteKey(conf.UserRedisIndex + "_" + sessionID)
	if utils.RECODE_OK != delResp.Errno {
		//删除失败直接返回
		utils.NewLog().Info("DeleteKey failed")
		return delResp

	}
	//再更新数据库
	updateResp := model.UpdateUserInfo(sessionID, key, value)
	if utils.RECODE_OK != updateResp.Errno {
		utils.NewLog().Info("UpdateUserInfo failed")

	}
	return updateResp

}
