package route

import (
	"crawler/internal/app/mu"
	adminAuth "crawler/internal/route/admin/auth"
	"crawler/internal/route/admin/node"
	"crawler/internal/route/admin/site"
	idxAuth "crawler/internal/route/index/auth"
	"crawler/internal/route/index/favor"
	"crawler/internal/route/index/hot"
	"crawler/internal/route/middleware"
	"crawler/internal/route/oauth"
	"github.com/gin-contrib/cors"
	"os"
	"path/filepath"
)

func RegisterStatic() {
	r := mu.App.Gin

	pwd, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	path := filepath.Dir(pwd)

	r.StaticFile("/", path+"/public/index.html")
	r.StaticFile("/admin", path+"/public/admin.html")

	r.StaticFile("favicon.png", path+"/public/favicon.png")
	r.Static("/static", path+"/public/static")
}

func RegisterRoutes() {
	r := mu.App.Gin

	c := cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
	})
	r.Use(c)

	// Auth操作
	r.GET("/auth_config", oauth.Config)
	r.GET("/oauth/auth", oauth.Auth)
	r.GET("/oauth/callback", oauth.Callback)

	// 前端路由
	r.GET("/config", hot.Tabs)
	r.GET("/logout", idxAuth.Logout)
	idx := r.Group("")
	idx.Use(middleware.ApiAuth(false))
	{
		idx.GET("/list", hot.List)
	}

	api := r.Group("/api")
	api.Use(middleware.ApiAuth(true))
	{
		api.GET("/info", idxAuth.Info)

		// 收藏管理
		api.GET("/favor/config", favor.Config)
		api.GET("/favor/list", favor.List)
		api.POST("/favor/add", favor.Add)
		api.POST("/favor/remove", favor.Remove)
	}

	// 后台路由管理
	admin := r.Group("/admin")
	admin.Use(middleware.AdminAuth())
	{
		admin.GET("/debug", site.Debug)
		admin.GET("/info", adminAuth.Info)
		// 节点管理
		admin.GET("/node", node.Info)
		admin.GET("/node/list", node.List)
		admin.POST("/node/upsert", node.CreateOrUpdateNode)
		admin.GET("/node/del", node.Del)

		// 站点管理
		admin.GET("/site", site.Info)
		admin.GET("/site/list", site.List)
		admin.POST("/site/update", site.UpdateSite)
	}
}
