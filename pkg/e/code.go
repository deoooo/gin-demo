package e

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400
	NeedAuth       = 401
	AuthFail       = 401
	// 保存图片失败
	ERROR_UPLOAD_SAVE_IMAGE_FAIL = 30001
	// 检查图片失败
	ERROR_UPLOAD_CHECK_IMAGE_FAIL = 30002
	// 校验图片错误，图片格式或大小有问题
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT = 30003
)
