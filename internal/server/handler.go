package server

import (
	"net/http"
	controller "scs-guard/internal/controllers"

	middleware "scs-guard/internal/middlewares"

	"github.com/labstack/echo/v4"
)

func (s *Server) MapHandlers(e *echo.Echo) error {
	// Init handlers
	missionHandler := controller.NewMissionHandler(*s.deps.MissionService)

	mw := middleware.NewMiddlewareManager(s.cfg, []string{"*"}, s.logger)
	e.Use(mw.RequestLoggerMiddleware)
	e.Use(mw.ErrorHandlerMiddleware)
	v1 := e.Group("/api/v1")

	health := v1.Group("/health")
	missionGroup := v1.Group("/missions", mw.JWTAuth)

	health.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})
	missionHandler.RegisterRoutes(missionGroup)

	return nil

}
