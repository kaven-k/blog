package image_service

import (
	"fmt"
	"io"
	"mime/multipart"
	"path"
	"server/global"
	"server/models"
	"server/models/ctype"
	"server/plugins/qiniu"
	"server/utils"
	"strings"
)

var (
	// WhiteImageList 图片上传的白名单
	WhiteImageList = []string{
		"jpg",
		"jpeg",
		"png",
		"webp",
		"ico",
		"tiff",
		"gif",
		"svg",
	}
)

type FileUploadResponse struct {
	Filename  string `json:"file_name"`  // 文件名
	IsSuccess bool   `json:"is_success"` // 是否上传成功
	Msg       string `json:"msg"`        // 消息
}

// ImageUploadService 处理图片文件上传的方法（向数据库里面添加数据）
// 如果重复就不添加  这里需要传递一个 file
func (ImageService) ImageUploadService(file *multipart.FileHeader) (res FileUploadResponse) {
	fileName := file.Filename
	basePath := global.Config.Upload.Path // basePath 指向的是文件上传的地址
	filePath := path.Join(basePath, file.Filename)
	// 对于已经存在的部分，先定义好
	res.Filename = filePath // 默认是这个，默认是上传失败
	// 文件白名单判断
	nameList := strings.Split(fileName, ".")
	suffix := strings.ToLower(nameList[len(nameList)-1])
	// 判断某个字符串是否存在于函数中
	if !utils.InList(suffix, WhiteImageList) {
		res.Msg = "非法文件"
		return
	}

	// 判断大小,有的图片大有的图片小，大的就上传小的就不上传，
	// 写一个响应的数据，哪些图片上传成功哪些图片上传失败
	size := float64(file.Size) / float64(1024*1024)
	if size >= float64(global.Config.Upload.Size) {
		res.Msg = fmt.Sprintf("图片大小超过设定大小！当前大小为：%2fMB ,图片规定大小为：%dMB", size, global.Config.Upload.Size)
		return
	}

	// 读取文件内容 hash
	fileObj, err := file.Open()
	if err != nil {
		global.Log.Error(err)
	}
	byteData, err := io.ReadAll(fileObj)
	imageHash := utils.Md5(byteData)
	// 去数据库中查这个图片是否存在
	var bannerModel models.BannerModel
	err = global.DB.Take(&bannerModel, "hash = ?", imageHash).Error
	if err == nil {
		// 找到了
		res.Msg = "图片已存在"
		res.Filename = bannerModel.Path
		return
	}
	fileType := ctype.Local // 指的是文件上传到本地还是七牛
	res.Msg = "图片上传成功"
	res.IsSuccess = true
	// 是否上传到七牛云
	if global.Config.QiNiu.Enable {
		filePath, err = qiniu.UploadImage(byteData, fileName, global.Config.QiNiu.Prefix)
		if err != nil {
			global.Log.Error(err)
			res.Msg = err.Error()
			return
		}
		res.Filename = filePath
		res.Msg = "上传七牛成功"
		fileType = ctype.QiNiu
	}
	// 写入数据库
	global.DB.Create(&models.BannerModel{
		Path:          filePath,
		Hash:          imageHash,
		Name:          fileName,
		ImageLocation: fileType,
	})
	return
}
