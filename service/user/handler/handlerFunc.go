package handler

import (
	"encoding/json"
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

func DeleteImgByHandler(sessionId string, m map[string]string) {
	utils.AntsPool.Pool.Submit(func() {
		//查数据库
		mysqlResp := model.GetUserInfo(sessionId)
		if utils.RECODE_OK != mysqlResp.Errno {
			//失败了直接返回
			//TODO 记录
			utils.NewLog().Info("mysql GetUserInfo failed:", mysqlResp)
			return
		}
		user := &model.UserPo{}
		json.Unmarshal(mysqlResp.Data, user)
		if user.Avatar_url != "" {
			//已经上传过头像了,就启动协程删除
			utils.NewLog().Debug("DeleteFileByHandler....")
			err := model.FastDfsClient.Client.DeleteFile(user.Avatar_url)
			if err != nil {
				//TODO 记录
				utils.NewLog().Error("FastDfsClient.Client.DeleteFile fail....")
			}
		}
		//更新数据库
		updateResp := model.UpdateUserInfo(sessionId, m)
		if utils.RECODE_OK != updateResp.Errno {
			//更新失败记录 TODO
			utils.NewLog().Error("FastDfsClient.Client.DeleteFile fail....")
		}
		return

	})
}

func UpdateUserInfoByHandler(sessionID string, m map[string]string) kitex_gen.Response {

	//先更新数据库
	updateResp := model.UpdateUserInfo(sessionID, m)
	if utils.RECODE_OK != updateResp.Errno {
		utils.NewLog().Info("UpdateUserInfo failed")
		return updateResp

	}
	//再更新缓存
	redisResp := UpdateRedisUserInfo(sessionID, m)
	if utils.RECODE_OK != redisResp.Errno {
		//删除失败直接返回
		utils.NewLog().Info("UpdateUserInfo failed")

	}
	return redisResp

}

func UpdateRedisUserInfo(sessionId string, m map[string]string) kitex_gen.Response {
	getResp := model.GetRedisUserInfo(sessionId)
	utils.NewLog().Info("GetRedisUserInfo result:", getResp.Errmsg)
	if utils.RECODE_OK != getResp.Errno {
		utils.NewLog().Info("GetRedisUserInfo fail:", getResp)
		return utils.UserResponse(utils.RECODE_SESSIONERR, nil)
	}
	userByte, err := json.Marshal(m)
	utils.NewLog().Debug("userByte:", string(userByte))
	if err != nil {
		utils.NewLog().Info("json.Marshal error:", getResp)
		return utils.UserResponse(utils.RECODE_SERVERERR, nil)
	}
	user := &model.UserPo{}
	json.Unmarshal(getResp.Data, user)
	json.Unmarshal(userByte, user)
	utils.NewLog().Debug("user:", user)
	data, _ := json.Marshal(user)
	saveResp := model.SaveRedisUserInfo(sessionId, data)
	if utils.RECODE_OK != saveResp.Errno {
		utils.NewLog().Info("SaveRedisUserInfo fail:", saveResp)
		return utils.UserResponse(utils.RECODE_SESSIONERR, nil)
	}
	return utils.UserResponse(utils.RECODE_OK, nil)

}
