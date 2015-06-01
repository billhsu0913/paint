package controllers

import "github.com/revel/revel"
import "mx.com/log"

import "encoding/base64"
import "strings"
import "fmt"
import "io/ioutil"

import "crypto/md5"

type Show struct {
	*revel.Controller
}

const (
	// Need Modify
	filepath = "/Users/test/test"
)

func (c Show) Index() revel.Result {

	return c.Render()
}

func (c Show) PutImage() revel.Result {

	var datainfo string
	c.Params.Bind(&datainfo, "datainfo")
	var datas = strings.Split(datainfo, ",")

	data, err := base64.StdEncoding.DecodeString(datas[len(datas)-1]) //成图片文件并把文件写入到buffer
	if err != nil {
		log.Error("<putImage> decode string failed :", err)
		return c.RenderJson(map[string]interface{}{
			"success": false,
			"info":    "Invalid Image",
		})
	}

	md5Sum := md5.Sum([]byte(data))
	fileName := fmt.Sprintf("%x", md5Sum)
	err = ioutil.WriteFile(filepath+"/"+fileName+".png", data, 0666)
	if err != nil {
		log.Error("<putImage> write file failed:", err)
		return c.RenderJson(map[string]interface{}{
			"success": false,
			"info":    "存储图片失败，请稍后重试",
		})
	}

	return c.RenderJson(map[string]interface{}{
		"success":  true,
		"filename": fileName,
	})
}

func (c Show) GetImage() revel.Result {

	var file string
	c.Params.Bind(&file, "file")

	fileByte, err := ioutil.ReadFile(filepath + "/" + file + ".png")
	if err != nil {
		log.Error("<GetImage> readfile failed: ", err)
	}

	fileBase := base64.StdEncoding.EncodeToString(fileByte)
	c.RenderArgs["file"] = fileBase

	return c.Render()
}
