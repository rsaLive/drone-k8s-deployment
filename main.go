package main

import (
	"drone-k8s/internal/service"
	"drone-k8s/models"
	"drone-k8s/pkg/config"
	"drone-k8s/pkg/plugin"
	"fmt"
	"reflect"
)

func main() {
	//os.Setenv("PLUGIN_DASHBOARD", "kuboard")
	//os.Setenv("DRONE_COMMIT_BRANCH", "daiyb")
	config.Init()
	oData := plugin.Bind()
	models.PConfigData = &oData
	if models.PConfigData.Debug {
		v := reflect.ValueOf(oData)
		t := v.Type()
		for i := 0; i < t.NumField(); i++ {
			fmt.Println(t.Field(i).Name, "===", v.Field(i).Interface())
		}
	}
	service.Exec()
}
