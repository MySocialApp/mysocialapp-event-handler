package main

import (
	"encoding/json"
	"fmt"
	msamodule "github.com/MySocialApp/mysocialapp-event-handler/modules"
	"gopkg.in/yaml.v2"
	"log"
	"net/http"
	"reflect"
)

type Config struct {
	Http struct {
		Bind string `yaml:"bind"`
	} `yaml:"http-bind"`
	Language struct {
		Default string `yaml:"default"`
	} `yaml:"language"`
	EventHandler []Handler `yaml:"handlers"`
	Modules      struct {
		Email *msamodule.EmailModule `yaml:"email"`
	} `yaml:"modules"`
}

type Handler struct {
	EventType string        `yaml:"type"`
	Path      string        `yaml:"path"`
	Method    string        `yaml:"method"`
	Actions   []interface{} `yaml:"actions"`
}

func (h *Handler) GetMethod() string {
	if h.Method != "" {
		return h.Method
	}
	return http.MethodPost
}

type Action interface {
	Init()
	Do(event interface{}, config *Config) error
}

func GetAction(event interface{}) Action {
	v := reflect.ValueOf(event)
	if !v.IsValid() {
		log.Printf("error reflecting value for %v", event)
		return nil
	}
	typeValue := v.MapIndex(reflect.ValueOf("type"))
	if !typeValue.IsValid() || typeValue.IsNil() {
		log.Printf("type value invalid for %v", event)
		return nil
	}
	switch fmt.Sprintf("%s", typeValue.Interface()) {
	case "email":
		var a ActionEmail
		ConvertUsingYaml(event, &a)
		return &a
		break
	default:
		log.Printf("action type not found (%s) for %v", typeValue.Interface(), event)
	}

	return nil
}

// TODO : must have a way to improve this
func ConvertUsingYaml(event interface{}, action interface{}) {
	out, _ := yaml.Marshal(event)
	yaml.Unmarshal(out, action)
}

func ConvertUsingJson(event interface{}, action interface{}) {
	out, _ := json.Marshal(event)
	json.Unmarshal(out, action)
}
