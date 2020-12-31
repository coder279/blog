package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"myblog/global"
	"myblog/internal/model"
	"myblog/internal/routers"
	setting2 "myblog/pkg/setting"
	"net/http"
	"time"
)

func init(){
	err := SetUpSetting()
	if err != nil {
		log.Fatalf("init.setupsetting err:%v",err)
	}
}

func SetUpSetting() error {
	setting,err := setting2.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server",&global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App",&global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database",&global.DataBaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeOut *= time.Second
	global.ServerSetting.WriteTimeOut *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine,err = model.NewDBEngine(global.DataBaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr: ":8080",
		Handler: router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1<<20,
	}
	s.ListenAndServe()
}
