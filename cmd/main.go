package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"html/template"
	"log"
	"minidash/internal/controller"
	"minidash/internal/types"
	"strings"
)

type Greet struct {
	Greeting string `json:"greeting"`
	Name     string `json:"name"`
}

func main() {

	t := &types.Template{
		Templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	conf := ReadConfig()

	ctrl := controller.NewLinkController(conf)
	idxCtr := controller.NewIndexController(conf)

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		conf = ReadConfig()
		ctrl.Config = conf
	})

	e := echo.New()
	e.Use(middleware.Logger())
	//e.Static("/static", "static")
	e.Renderer = t

	//ctrl := controller.NewLinkController(*conf)

	e.GET("/links", ctrl.GetLinks)
	e.GET("/", idxCtr.GetIndex)
	//e.GET("/", func(c echo.Context) error {
	//	return c.String(http.StatusOK, "Hello, World!")
	//})

	e.Logger.Fatal(e.Start(":1323"))
}

func ReadConfig() *types.Configuration {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	var conf types.Configuration
	err = viper.Unmarshal(&conf)
	if err != nil {
		log.Fatal(err)
	}

	m := make(map[string]string)

	for _, link := range conf.Links {
		group := strings.ToLower(link.Group)
		group = strings.ReplaceAll(group, " ", "")
		m[group] = link.Group
	}
	conf.Filter = m
	return &conf
}
