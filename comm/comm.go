package comm

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Equipment represents a piece of equipment
type Equipment struct {
	ID string `json:"id"`
}

// EquipmentListResponse represents a JSON response for the list of equipment
type EquipmentListResponse struct {
	Equipment []Equipment `json:"equipment"`
}

// CreateServer creates and configure the Echo HTTP server
func CreateServer() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	// Unauthorized endpoint (i.e. login)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!")
	})

	authGroup := e.Group("", func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Verify authorization here
			authorized := true
			if !authorized {
				return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
			}

			return next(c)
		}
	})

	// Add all protected endpoints to authGroup to force them to go
	// through the security middleware.

	authGroup.GET("/equipment", func(c echo.Context) error {
		return c.JSON(http.StatusOK, &EquipmentListResponse{
			Equipment: []Equipment{
				{ID: "1"},
			},
		})
	})

	return e
}
