package controllers

import (
	"net/url"
	"time"

	"github.com/mindoc-org/mindoc/conf"
	"github.com/mindoc-org/mindoc/models"
)

type WKController struct {
	BaseController
}

func (c *WKController) Prepare() {
	c.BaseController.Prepare()
	//如果没有开启匿名访问，则跳转到登录页面
	if !c.EnableAnonymous && c.Member == nil {
		c.Redirect(conf.URLFor("AccountController.Login")+"?url="+url.PathEscape(conf.BaseUrl+c.Ctx.Request.URL.RequestURI()), 302)
	}
}

func (c *WKController) Index() {
	c.Prepare()
	c.TplName = "v2/watchkeeper/dashboard.tpl"

	currentTime := time.Now()

	wkInfo := models.NewWatchkeeperResult().FindByTime(currentTime)
	// memberId := 0
	// if c.Member != nil {
	// 	memberId = c.Member.MemberId
	// }

	c.Data["PAGE_TITLE"] = c.Tr("common.home")
	c.Data["CONTENT_NAME"] = c.Tr("common.operation_center")
	c.Data["app_primary"] = wkInfo.AppPrimary
	c.Data["team_leader"] = wkInfo.TeamLeader
	c.Data["platform"] = wkInfo.Platform

}

func (c *WKController) GetWatchKeeperInfo() {
	currentTime := time.Now()

	wkInfo := models.NewWatchkeeperResult().FindByTime(currentTime)

	c.JsonResult(200, "OK", wkInfo)
}
