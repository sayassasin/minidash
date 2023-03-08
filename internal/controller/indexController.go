package controller

import (
	"github.com/labstack/echo/v4"
	"minidash/internal/types"
	"net/http"
	"strings"
)

type IndexController struct {
	Config *types.Configuration
}

type indexWrapper struct {
	Links  *[]types.Link
	Filter *map[string]string
	Title  *string
}

func NewIndexController(config *types.Configuration) *IndexController {
	return &IndexController{Config: config}
}

func (c *IndexController) GetIndex(ctx echo.Context) error {
	m := make(map[string]string)

	for _, link := range c.Config.Links {
		group := strings.ToLower(link.Group)
		group = strings.ReplaceAll(group, " ", "")
		m[group] = link.Group
	}

	wrapper := indexWrapper{
		Links:  &c.Config.Links,
		Filter: &c.Config.Filter,
		Title:  c.Config.MetaData.Title,
	}
	return ctx.Render(http.StatusOK, "index.html", &wrapper)
}
