package main

import (
	"flag"
	"github.com/MySocialApp/mysocialapp-event-handler/models"
	"github.com/kataras/iris"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var (
	AppConfigService = msaevents.NewAppConfigService()

	configFilename = flag.String("config", "", "YAML config file")
)

func main() {
	flag.Parse()
	conf := loadConfig(*configFilename)
	if conf == nil {
		log.Fatal("fail to load conf")
	}

	app := iris.New()
	publicApi := app.Party("/api/v1/public")

	webHookApi := publicApi.Party("/event")
	log.Printf("%+v", conf)
	for _, h := range conf.EventHandler {
		log.Printf("load path %s", h.Path)
		handler := h
		actions := []Action{}
		for _, action := range handler.Actions {
			a := GetAction(action)
			if a != nil {
				actions = append(actions, a)
			} else {
				log.Printf("actions not found for %v", action)
			}
		}
		webHookApi.Handle(handler.GetMethod(), handler.Path, func(ctx iris.Context) {
			var e interface{}
			if err := ctx.ReadJSON(&e); err != nil {
				log.Printf("fail to read request body: %s", err.Error())
				return
			}
			for _, action := range actions {
				if err := action.Do(e, conf); err != nil {
					log.Printf("action failed (%s): %s",handler.Path, err.Error())
				}
			}
		})
	}

	app.Run(iris.Addr(conf.Http.Bind), iris.WithoutVersionChecker)
}

func loadConfig(filename string) *Config {
	var e Config
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("fail to load file conf '%s': %s", filename, err.Error())
		return nil
	}
	if err := yaml.Unmarshal(fileContent, &e); err != nil {
		log.Fatalf("fail to read conf file: %s", err.Error())
		return nil
	}
	return &e
}
