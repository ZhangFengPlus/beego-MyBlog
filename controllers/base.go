package controllers

import (
	"beego-blog/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strings"
)

type baseController struct {
	beego.Controller
	o orm.Ormer
	controllerName string
	actionName     string
}

//后台登录第一步
func (p *baseController) Prepare() {

	//获取此方法名   AdminController
	controllerName, actionName := p.GetControllerAndAction()
	// 用strings.ToLower  转换 小写 提取  0-5 之前 的 字母  admin
	p.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])

	//Login 转换为小写 login
	p.actionName = strings.ToLower(actionName)

	//切换数据库
	p.o = orm.NewOrm()


	//判断是admin  但是 控制器不是 login 的 session 没有 用户 去登录
	if strings.ToLower(p.controllerName) == "admin" && strings.ToLower(p.actionName) != "login" {
		if p.GetSession("user") == nil {
			p.History("未登录", "/admin/login")
			//p.Ctx.WriteString(p.controllerName +"==="+ p.actionName)
		}
	}

	//初始化前台页面相关元素
	if strings.ToLower(p.controllerName) == "blog" {

		p.Data["actionName"] = strings.ToLower(actionName)
		var result []*models.Config

		p.o.QueryTable(new(models.Config).TableName()).All(&result)


		configs := make(map[string]string)
		for _, v := range result {
			configs[v.Name] = v.Value
		}
		p.Data["config"] = configs
	}

}

//错误跳转
func (p *baseController) History(msg string, url string) {
	if url == "" {
		p.Ctx.WriteString("<script>alert('" + msg + "');window.history.go(-1);</script>")
		p.StopRun()
	} else {
		p.Redirect(url, 302)
	}
}

//获取用户IP地址
func (p *baseController) getClientIp() string {
	s := strings.Split(p.Ctx.Request.RemoteAddr, ":")
	return s[0]
}
