package server

import (
	"net/http"
	controller "scs-guard/internal/controllers"

	middleware "scs-guard/internal/middlewares"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "scs-guard/docs" // Import generated docs
)

func (s *Server) MapHandlers(e *echo.Echo) error {
	// Init handlers
	missionHandler := controller.NewMissionHandler(*s.deps.MissionService)

	mw := middleware.NewMiddlewareManager(s.cfg, []string{"*"}, s.logger)
	e.Use(mw.RequestLoggerMiddleware)
	e.Use(mw.ErrorHandlerMiddleware)
	e.Use(mw.ResponseStandardizer)

	// Swagger documentation route
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := e.Group("/api/v1")

	health := v1.Group("/health")
	missionGroup := v1.Group("/missions", mw.JWTAuth)

	// Health check endpoint
	// @Summary Health check
	// @Description Check if the service is running
	// @Tags health
	// @Accept json
	// @Produce json
	// @Success 200 {object} map[string]string "Service is healthy"
	// @Router /api/v1/health [get]
	health.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})
	missionHandler.RegisterRoutes(missionGroup)

	return nil

}
