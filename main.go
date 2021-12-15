package content_sdk

import (
	"errors"
	"github.com/imroc/req"
	"os"
	"path/filepath"
)

var (
	StatusFail = errors.New("响应不正确")
)

type Sdk struct {
	Host string // 服务器地址
}

func New() *Sdk {
	return &Sdk{
		Host: "http://a.resok.cn",
	}
}

type GetAccessResp struct {
	Token string `json:"token"`
}

type ImgCheckResp struct {
	Pass bool   `json:"pass"`
	Msg  string `json:"msg"`
}

// GetAccessToken 获取token
func (c *Sdk) GetAccessToken(appId string) (string, error) {
	resp, err := req.Get(c.Host + "/access_token/" + appId)
	if err != nil {
		return "", err
	}
	if resp.Response().StatusCode != 200 {
		return "", StatusFail
	}
	var r GetAccessResp
	err = resp.ToJSON(&r)
	if err != nil {
		return "", err
	}
	return r.Token, nil
}

// RefreshAccessToken 刷新token
func (c *Sdk) RefreshAccessToken(appId string) error {
	resp, err := req.Get(c.Host + "/reset_access_token/" + appId)
	if err != nil {
		return err
	}
	if resp.Response().StatusCode != 200 {
		return StatusFail
	}
	return nil
}

// TextSecCheck 文字安全校验 返回校验是否通过 true为通过 false为不通过
func (c *Sdk) TextSecCheck(text string) bool {
	var r = map[string]interface{}{
		"content": text,
	}
	resp, err := req.Post(c.Host+"/hit_text", req.BodyJSON(&r))
	if err != nil {
		return false
	}
	if resp.Response().StatusCode != 200 {
		return false
	}
	return true
}

// ImageSecCheck 图片安全校验 返回校验是否通过
func (c *Sdk) ImageSecCheck(imgPath string) (bool, error) {
	file, err := os.Open(imgPath)
	if err != nil {
		return false, err
	}
	p := filepath.Base(imgPath)
	ext := filepath.Ext(imgPath)
	name := p[0 : len(p)-len(ext)]
	resp, err := req.Post(c.Host+"/img_check", req.FileUpload{
		File:      file,
		FieldName: "file", // FieldName is form field name
		FileName:  name + ext,
	})
	if err != nil {
		return false, err
	}
	var r ImgCheckResp
	err = resp.ToJSON(&r)
	if err != nil {
		return false, err
	}
	return r.Pass, nil
}