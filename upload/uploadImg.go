package upload

import (
	"context"
	"gin-blog/errormsg"
	"gin-blog/utils"
	"mime/multipart"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

var Zone = utils.Zone
var AccessKey = utils.AccessKey
var SecretKey = utils.SecretKey
var Bucket = utils.Bucket
var ImgUrl = utils.QiNiuServe

func UploadImg(file multipart.File, fileSize int64) (string, int) {
	// 获取凭证
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	putPolicy.Expires = 7200 //2小时有效期
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)
	// 对上传进行参数配置
	config := storage.Config{
		// 空间名
		Zone: selectZone(Zone),
		// 是否用cdn加速
		UseCdnDomains: false,
		// 是否用https
		UseHTTPS: false,
	}
	// 构建上传对象
	formUploader := storage.NewFormUploader(&config)
	ret := storage.PutRet{}
	//
	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &storage.PutExtra{})
	if err != nil {
		return "", errormsg.ERROR
	}
	return ImgUrl + ret.Key, errormsg.SUCCESS
}

// Zone设置
func selectZone(id int) *storage.Zone {
	switch id {
	case 1:
		return &storage.ZoneHuadong
	case 2:
		return &storage.ZoneHuabei
	case 3:
		return &storage.ZoneHuanan
	default:
		return &storage.ZoneHuadong
	}
}
