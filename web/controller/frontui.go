package controller

import (
	"github.com/gin-gonic/gin"
)

// FRONTUIController is the main controller for the FRONT-UI panel, managing sub-controllers.
type FRONTUIController struct {
	BaseController

	settingController     *SettingController
	xraySettingController *XraySettingController
}

// NewFRONTUIController creates a new FRONTUIController and initializes its routes.
func NewFRONTUIController(g *gin.RouterGroup) *FRONTUIController {
	a := &FRONTUIController{}
	a.initRouter(g)
	return a
}

// initRouter sets up the main panel routes and initializes sub-controllers.
func (a *FRONTUIController) initRouter(g *gin.RouterGroup) {
	g = g.Group("/panel")
	g.Use(a.checkLogin)

	g.GET("/", a.index)
	g.GET("/inbounds", a.inbounds)
	g.GET("/settings", a.settings)
	g.GET("/xray", a.xraySettings)

	a.settingController = NewSettingController(g)
	a.xraySettingController = NewXraySettingController(g)
}

// index renders the main panel index page.
func (a *FRONTUIController) index(c *gin.Context) {
	html(c, "index.html", "pages.index.title", nil)
}

// inbounds renders the inbounds management page.
func (a *FRONTUIController) inbounds(c *gin.Context) {
	html(c, "inbounds.html", "pages.inbounds.title", nil)
}

// settings renders the settings management page.
func (a *FRONTUIController) settings(c *gin.Context) {
	html(c, "settings.html", "pages.settings.title", nil)
}

// xraySettings renders the Xray settings page.
func (a *FRONTUIController) xraySettings(c *gin.Context) {
	html(c, "xray.html", "pages.xray.title", nil)
}
