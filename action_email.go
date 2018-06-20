package main

import (
	"bytes"
	"github.com/MySocialApp/mysocialapp-event-handler/models"
	"html/template"
	"io/ioutil"
	"log"
	"path/filepath"
	"os"
)

type ActionEmail struct {
	From    ActionEmailContact   `yaml:"from"`
	To      []ActionEmailContact `yaml:"to"`
	Cc      []ActionEmailContact `yaml:"cc"`
	Ccc     []ActionEmailContact `yaml:"ccc"`
	Subject string               `yaml:"subject"`
	Body    struct {
		Html     bool   `yaml:"html"`
		Template string `yaml:"template"`
	} `yaml:"body"`
	Pdf *struct {
		Filename string `yaml:"filename"`
		Template string `yaml:"template"`
	} `yaml:"pdf"`
}

func (a *ActionEmail) Init() {
}

func (a *ActionEmail) Do(data interface{}, config *Config) error {
	var event msaevents.EventCreatedUser
	ConvertUsingJson(data, &event)
	appConfig, err := AppConfigService.GetConfig(event.ConfigId)
	if err != nil {
		return err
	}
	tData := ActionEmailTemplateData{Config: config, AppConfig: appConfig, Event: &event}
	email := config.Modules.Email.Prepare().
		Subject(a.viewTemplate(&tData, a.Subject)).
		From(a.viewTemplate(&tData, a.From.Email)).
		Body(a.viewFileTemplate(&tData, a.Body.Template), a.Body.Html)
	if a.To != nil {
		for _, to := range a.To {
			email.AddTo(to.Email)
		}
	}
	if a.Cc != nil {
		for _, cc := range a.Cc {
			email.AddTo(cc.Email)
		}
	}
	if a.Ccc != nil {
		for _, ccc := range a.Ccc {
			email.AddTo(ccc.Email)
		}
	}
	if a.Pdf != nil {
		b := bytes.NewBufferString(a.viewFileTemplate(&tData, a.Pdf.Template))
		email.AddHtmlToPdfFile(b, a.Pdf.Filename)
	}
	return email.Send()
}

func (a *ActionEmail) viewTemplate(data *ActionEmailTemplateData, t string) string {
	tpl := template.New("")
	tpl, err := tpl.Parse(t)
	if err != nil {
		log.Printf("fail to parse template (%s)", err.Error())
		return t
	}
	buff := bytes.NewBufferString("")
	if err := tpl.Execute(buff, data); err != nil {
		log.Printf("fail to generate template (%s)", err.Error())
		return t
	}
	return buff.String()
}

func (a *ActionEmail) viewFileTemplate(data *ActionEmailTemplateData, filename string) string {
	path := filepath.Dir(*configFilename)
	content, err := ioutil.ReadFile(path + string(os.PathSeparator) + filename)
	if err != nil {
		return err.Error()
	}
	return a.viewTemplate(data, string(content))
}

type ActionEmailContact struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
}

type ActionEmailTemplateData struct {
	Config    *Config
	AppConfig *msaevents.AppConfig
	Event     *msaevents.EventCreatedUser
}
