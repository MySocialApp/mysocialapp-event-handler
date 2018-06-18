package msaevents

import (
	"time"
	"github.com/patrickmn/go-cache"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"net/http"
)

func NewAppConfigService() *AppConfigService {
	a := AppConfigService{}
	a.init()
	return &a
}

type AppConfigService struct {
	Cache *cache.Cache
}

func (a *AppConfigService) init() {
	a.Cache = cache.New(1*time.Minute, 1*time.Minute)
}

func  (a *AppConfigService) GetConfig(configId string) (*AppConfig, error) {
	if r, ok := a.Cache.Get(a.getCacheIndex(configId)); ok {
		appConfig := r.(AppConfig)
		return &appConfig, nil
	}
	appConfig, err := a.DownloadConfig(configId)
	if err != nil {
		return nil, err
	}
	a.Cache.Set(a.getCacheIndex(configId), *appConfig, cache.DefaultExpiration)
	return appConfig, nil
}

func  (a *AppConfigService) getCacheIndex(i string) string {
	return fmt.Sprintf("config:%s", i)
}


func (a *AppConfigService) DownloadConfig(configId string) (*AppConfig, error) {
	u := fmt.Sprintf("https://api.mysocialapp.io/api/v1/config/%s", configId)
	resp, err := http.Get(u)
	if err != nil {
		fmt.Printf("fail to load config %s: %s\n", u, err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("fail to read body config %s: %s", u, err.Error())
		return nil, err
	}
	var appConfig AppConfig
	if err := json.Unmarshal(body, &appConfig); err != nil {
		fmt.Printf("fail to unmarshal config body %s: %s", u, err.Error())
		return nil, err
	}
	return &appConfig, nil
}

type AppConfig struct {

}