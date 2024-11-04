package images_api

import (
	"github.com/gin-gonic/gin"
	"io/fs"
	"os"
	"server/global"
	"server/models/res"
	"server/services"
	"server/services/image_service"
)

// ImageUploadView 上传多个图片，返回图片的url
func (ImagesApi) ImageUploadView(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	fileList, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("不存在的图片文件！", c)
		return
	}
	// 判断路径是否存在
	// 转换成列表逐级进行判断
	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {
		// 递归创建
		err = os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
	}

	// 不存在就创建
	var resList []image_service.FileUploadResponse
	// 实例化一个app
	for _, file := range fileList {
		//上传文件并保存
		serviceRes := services.ServiceApp.ImageService.ImageUploadService(file)
		if !serviceRes.IsSuccess {
			resList = append(resList, serviceRes)
			continue
		}
		// 上传成功的
		if !global.Config.QiNiu.Enable {
			// 本地还得保存一下
			err = c.SaveUploadedFile(file, serviceRes.Filename) // 在什么情况下会走这里，需要判断
			if err != nil {
				global.Log.Error(err)
				serviceRes.Msg = err.Error()
				serviceRes.IsSuccess = false
				resList = append(resList, serviceRes)
				continue
			}
		}
		resList = append(resList, serviceRes)
	}
	res.OkWithData(resList, c)
}
