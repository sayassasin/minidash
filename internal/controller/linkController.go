package controller

import (
	"github.com/labstack/echo/v4"
	"minidash/internal/types"
	"net/http"
)

type LinkController struct {
	Config *types.Configuration
}

func NewLinkController(config *types.Configuration) *LinkController {
	return &LinkController{Config: config}
}

func (c *LinkController) GetLinks(ctx echo.Context) error {
	if len(ctx.QueryParams()) == 0 {
		return ctx.Render(http.StatusOK, "links.html", &c.Config.Links)
	}
	filter := ctx.QueryParam("groups")
	links := c.Config.Links
	filteredGroup := c.Config.Filter[filter]
	var filteredLinks []types.Link
	for _, link := range links {
		if link.Group == filteredGroup {
			filteredLinks = append(filteredLinks, link)
		}
	}

	return ctx.Render(http.StatusOK, "links", &filteredLinks)
}
