package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/joyrex2001/kubedock/internal/container"
	"github.com/joyrex2001/kubedock/internal/kubernetes"
	"github.com/joyrex2001/kubedock/internal/server/httputil"
	"github.com/joyrex2001/kubedock/internal/util/keyval"
)

// containerRouter is the object that facilitate all container
// related API endpoints.
type Router struct {
	factory    container.Factory
	kubernetes kubernetes.Kubernetes
}

// New will instantiate a containerRouter object.
func New(router *gin.Engine, db keyval.Database, kube kubernetes.Kubernetes) *Router {
	f := container.NewFactory(db)
	cr := &Router{
		factory:    f,
		kubernetes: kube,
	}
	cr.initRoutes(router)
	return cr
}

// initRoutes will add all suported routes.
func (cr *Router) initRoutes(router *gin.Engine) {
	router.POST("/containers/create", cr.ContainerCreate)
	router.POST("/containers/:id/start", cr.ContainerStart)
	router.GET("/containers/:id/json", cr.ContainerInfo)
	router.DELETE("/containers/:id", cr.ContainerDelete)
	router.POST("/containers/:id/exec", cr.ContainerExec)
	router.GET("/containers/:id/logs", cr.ContainerLogs)
	router.POST("/exec/:id/start", cr.ExecStart)
	router.GET("/exec/:id/json", cr.ExecInfo)
	router.PUT("/containers/:id/archive", cr.PutArchive)
	router.GET("/networks", cr.NetworksList)
	router.POST("/networks/create", cr.NetworksCreate)
	router.GET("/images/json", cr.ImageList)
	router.POST("/images/create", cr.ImageCreate)
	router.GET("/images/:image/*json", cr.ImageJSON)
	router.GET("/_ping", cr.Ping)
	router.HEAD("/_ping", cr.Ping)
	router.GET("/info", cr.Info)
	router.GET("/version", cr.Version)
	router.GET("/healthz", cr.Healthz)

	// not supported at the moment
	router.POST("/containers/:id/stop", httputil.NotImplemented)
	router.POST("/containers/:id/kill", httputil.NotImplemented)
	router.GET("/containers/json", httputil.NotImplemented)
	router.GET("/containers/:id/top", httputil.NotImplemented)
	router.GET("/containers/:id/changes", httputil.NotImplemented)
	router.GET("/containers/:id/export", httputil.NotImplemented)
	router.GET("/containers/:id/stats", httputil.NotImplemented)
	router.POST("/containers/:id/resize", httputil.NotImplemented)
	router.POST("/containers/:id/restart", httputil.NotImplemented)
	router.POST("/containers/:id/update", httputil.NotImplemented)
	router.POST("/containers/:id/rename", httputil.NotImplemented)
	router.POST("/containers/:id/pause", httputil.NotImplemented)
	router.POST("/containers/:id/unpause", httputil.NotImplemented)
	router.POST("/containers/:id/attach", httputil.NotImplemented)
	router.GET("/containers/:id/attach/ws", httputil.NotImplemented)
	router.POST("/containers/:id/wait", httputil.NotImplemented)
	router.HEAD("/containers/:id/archive", httputil.NotImplemented)
	router.GET("/containers/:id/archive", httputil.NotImplemented)
	router.POST("/containers/prune", httputil.NotImplemented)
	router.GET("/networks/reaper_default", httputil.NotImplemented)
}