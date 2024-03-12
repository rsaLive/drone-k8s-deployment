package plugin

import (
	"drone-k8s/models"
	"fmt"
	"github.com/spf13/viper"
	"reflect"
)

func Bind() (PConfig models.PConfig) {
	viper.AutomaticEnv() // 允许读取环境变量
	vPConfig := reflect.ValueOf(&PConfig).Elem()
	tPConfig := vPConfig.Type()
	for i := 0; i < tPConfig.NumField(); i++ {
		field := tPConfig.Field(i)
		envName := field.Tag.Get("envName")
		fieldValue := vPConfig.Field(i)
		switch field.Type.Name() {
		case "bool":
			v := viper.GetBool(envName)
			if fieldValue.CanSet() {
				vPConfig.Field(i).SetBool(v)
			} else {
				fmt.Printf("Cannot set value for field %s\n", field.Name)
			}
		case "string":
			v := viper.GetString(envName)
			if fieldValue.CanSet() {
				vPConfig.Field(i).SetString(v)
			} else {
				fmt.Printf("Cannot set value for field %s\n", field.Name)
			}
		case "int", "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64":
			v := viper.GetInt(envName)
			if fieldValue.CanSet() {
				vPConfig.Field(i).SetInt(int64(v))
			} else {
				fmt.Printf("Cannot set value for field %s\n", field.Name)
			}

		case "float", "float32", "float64":
			v := viper.GetFloat64(envName)
			if fieldValue.CanSet() {
				vPConfig.Field(i).SetFloat(v)
			} else {
				fmt.Printf("Cannot set value for field %s\n", field.Name)
			}
		}
	}
	return
}
