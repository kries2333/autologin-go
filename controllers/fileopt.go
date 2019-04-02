package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"path"
	"strings"
)

type FileOptUploadController struct {
	beego.Controller
}

func (c *FileOptUploadController) Get()  {
	c.TplName = "fileopt.html"
}

func (this *FileOptUploadController) Post()  {
	//image，这是一个key值，对应的是html中input type-‘file’的name属性值
	f, h, _ := this.GetFile("image")
	//得到文件的名称
	fileName := h.Filename
	arr := strings.Split(fileName, ":")
	if len(arr) > 1 {
		index := len(arr) - 1
		fileName = arr[index]
	}
	fmt.Println("文件名称:")
	fmt.Println(fileName)
	//关闭上传的文件，不然的话会出现临时文件不能清除的情况
	f.Close()
	//保存文件到指定的位置
	//static/uploadfile,这个是文件的地址，第一个static前面不要有/
	this.SaveToFile("image", path.Join("static/uploadfile",fileName))
	//显示在本页面，不做跳转操作
	this.TplName = "fileopt.html"
}