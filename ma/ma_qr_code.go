package ma

import (
	"io"
	"io/ioutil"
)

type QrCodeService interface {
	// 获取小程序页面二维码.
	// path 扫码进入的小程序页面路径，最大长度 128 字节，不能为空；对于小游戏，可以只传入 query 部分，来实现传参效果，如：传入 "?foo=bar"，即可在 wx.getLaunchOptionsSync 接口中的 query 参数获取到 {foo:"bar"}。
	// width 二维码的宽度，单位 px。最小 280px，最大 1280px
	CreateQrcodeBytes(path string, width int) ([]byte, error)
	CreateQrCodeReader(path string, width int) (io.ReadCloser, error)
	// 获取小程序页面二维码.
	// path 扫码进入的小程序页面路径，最大长度 128 字节，不能为空；对于小游戏，可以只传入 query 部分，来实现传参效果，如：传入 "?foo=bar"，即可在 wx.getLaunchOptionsSync 接口中的 query 参数获取到 {foo:"bar"}。
	CreateQrcode(path string) ([]byte, error)

	// 接口A: 获取小程序码.
	// path 扫码进入的小程序页面路径，最大长度 128 字节，不能为空；对于小游戏，可以只传入 query 部分，来实现传参效果，如：传入 "?foo=bar"，即可在 wx.getLaunchOptionsSync 接口中的 query 参数获取到 {foo:"bar"}。
	// width 二维码的宽度，单位 px。最小 280px，最大 1280px
	// auto_color 自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调
	// line_color auto_color 为 false 时生效，使用 rgb 设置颜色 例如 {"r":"xxx","g":"xxx","b":"xxx"} 十进制表示
	// is_hyaline 是否需要透明底色，为 true 时，生成透明底色的小程序码
	CreateWxaCodeBytes(path string, width int, autoColor bool, lineColor CodeLineColor, isTransparent bool) ([]byte, error)
	CreateWxaCodeReader(path string, width int, autoColor bool, lineColor CodeLineColor, isTransparent bool) (io.ReadCloser, error)
	// 接口A: 获取小程序码.
	// path 扫码进入的小程序页面路径，最大长度 128 字节，不能为空；对于小游戏，可以只传入 query 部分，来实现传参效果，如：传入 "?foo=bar"，即可在 wx.getLaunchOptionsSync 接口中的 query 参数获取到 {foo:"bar"}。
	CreateWxaCode(path string) ([]byte, error)

	// 接口B: 获取小程序码（永久有效、数量暂无限制）.
	// scene 最大32个可见字符，只支持数字，大小写英文以及部分特殊字符：!#$&'()*+,/:;=?@-._~，其它字符请自行编码为合法字符（因不支持%，中文无法使用 urlencode 处理，请使用其他编码方式）
	// page 必须是已经发布的小程序存在的页面（否则报错），例如 pages/index/index, 根路径前不要填加 /,不能携带参数（参数请放在scene字段里），如果不填写这个字段，默认跳主页面
	// width 二维码的宽度，单位 px，最小 280px，最大 1280px
	// auto_color 自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调，默认 false
	// line_color auto_color 为 false 时生效，使用 rgb 设置颜色 例如 {"r":"xxx","g":"xxx","b":"xxx"} 十进制表示
	// is_hyaline 是否需要透明底色，为 true 时，生成透明底色的小程序
	CreateWxaCodeUnlimitedBytes(scene, page string, width int, autoColor bool, lineColor CodeLineColor, isTransparent bool) ([]byte, error)
	CreateWxaCodeUnlimitedReader(scene, page string, width int, autoColor bool, lineColor CodeLineColor, isTransparent bool) (io.ReadCloser, error)
	// 接口B: 获取小程序码（永久有效、数量暂无限制）.
	// scene 最大32个可见字符，只支持数字，大小写英文以及部分特殊字符：!#$&'()*+,/:;=?@-._~，其它字符请自行编码为合法字符（因不支持%，中文无法使用 urlencode 处理，请使用其他编码方式）
	// page 必须是已经发布的小程序存在的页面（否则报错），例如 pages/index/index, 根路径前不要填加 /,不能携带参数（参数请放在scene字段里），如果不填写这个字段，默认跳主页面
	CreateWxaCodeUnlimited(scene, page string) ([]byte, error)
}

type BaseCode struct {
	Width     int           `json:"width"`
	AutoColor bool          `json:"auto_color"`
	IsHyaline bool          `json:"is_hyaline"`
	LineColor CodeLineColor `json:"line_color"`
}

type WxaCode struct {
	BaseCode
	Path string `json:"path"`
}

type WxaCodeUnlimited struct {
	BaseCode
	Scene string `json:"scene"`
	Page  string `json:"page"`
}

type CodeLineColor struct {
	R string
	B string
	G string
}

func DefaultCodeLineColor() CodeLineColor {
	return CodeLineColor{R: "0", B: "0", G: "0"}
}

type QrCodeServiceImpl struct {
	service Service
}

func newQrCodeService(service Service) QrCodeService {
	return &QrCodeServiceImpl{
		service: service,
	}
}

func (q *QrCodeServiceImpl) CreateQrcodeBytes(path string, width int) (bytes []byte, err error) {

	data := map[string]interface{}{"path": path, "width": width}
	err = q.service.post(&bytes, QrCodeUrl, "", data, q.service.GetAccessToken())
	return
}

func (q *QrCodeServiceImpl) CreateQrCodeReader(path string, width int) (reader io.ReadCloser, err error) {

	data := map[string]interface{}{"path": path, "width": width}
	err = q.service.post(reader, QrCodeUrl, "", data, q.service.GetAccessToken())
	return
}

func (q *QrCodeServiceImpl) CreateQrcode(path string) ([]byte, error) {

	return q.CreateQrcodeBytes(path, 430)
}

func (q *QrCodeServiceImpl) CreateWxaCodeBytes(path string, width int, autoColor bool, lineColor CodeLineColor, isTransparent bool) (bytes []byte, err error) {

	data := WxaCode{
		Path: path,
		BaseCode: BaseCode{
			Width:     width,
			AutoColor: autoColor,
			LineColor: lineColor,
			IsHyaline: isTransparent,
		},
	}
	err = q.service.post(&bytes, QrWxaCodeUrl, "", data, q.service.GetAccessToken())
	return
}

func (q *QrCodeServiceImpl) CreateWxaCodeReader(path string, width int, autoColor bool, lineColor CodeLineColor, isTransparent bool) (bytes io.ReadCloser, err error) {

	data := WxaCode{
		Path: path,
		BaseCode: BaseCode{
			Width:     width,
			AutoColor: autoColor,
			LineColor: lineColor,
			IsHyaline: isTransparent,
		},
	}
	err = q.service.post(bytes, QrWxaCodeUrl, "", data, q.service.GetAccessToken())
	return
}

func (q *QrCodeServiceImpl) CreateWxaCode(path string) ([]byte, error) {

	return q.CreateWxaCodeBytes(path, 430, true, DefaultCodeLineColor(), false)
}

func (q *QrCodeServiceImpl) CreateWxaCodeUnlimitedBytes(scene, path string, width int, autoColor bool, lineColor CodeLineColor,
	isTransparent bool) (bytes []byte, err error) {

	data := WxaCodeUnlimited{
		Page:  path,
		Scene: scene,
		BaseCode: BaseCode{
			Width:     width,
			AutoColor: autoColor,
			LineColor: lineColor,
			IsHyaline: isTransparent,
		},
	}
	err = q.service.post(&bytes, QrCodeUnlimitedUrl, "", data, q.service.GetAccessToken())
	return
}

func (q *QrCodeServiceImpl) CreateWxaCodeUnlimitedReader(scene, path string, width int, autoColor bool, lineColor CodeLineColor, isTransparent bool) (bytes io.ReadCloser, err error) {

	data := WxaCodeUnlimited{
		Page:  path,
		Scene: scene,
		BaseCode: BaseCode{
			Width:     width,
			AutoColor: autoColor,
			LineColor: lineColor,
			IsHyaline: isTransparent,
		},
	}
	err = q.service.post(bytes, QrCodeUnlimitedUrl, "", data, q.service.GetAccessToken())
	return
}

func (q *QrCodeServiceImpl) CreateWxaCodeUnlimited(scene, path string) ([]byte, error) {

	return q.CreateWxaCodeUnlimitedBytes(scene, path, 430, true, DefaultCodeLineColor(), false)
}

func (q QrCodeServiceImpl) BytesToFile(filename string, bytes []byte) error {
	// 覆盖式写入
	return ioutil.WriteFile(filename, bytes, 0664)
}
