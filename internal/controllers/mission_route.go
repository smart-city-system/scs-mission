package http

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func (h *MissionHandler) RegisterRoutes(g *echo.Group) {
	g.PATCH("/complete", h.CompleteStep())
	g.PUT("/update", h.UpdateIncidentInfo())
	g.GET("/me", func(c echo.Context) error {
		userID, ok := c.Get("user_id").(string)
		fmt.Println(userID)
		if !ok {
			return echo.NewHTTPError(401, "user_id not found in context")
		}
		return h.GetAssignments(userID)(c)
	})
}
