package config

var Data = new(AppConfig)

type AppConfig struct {
	DockerRepo struct {
		Host       string `mapstructure:"host"`
		ProjectEnv struct {
			Dev  string
			Test string
			Prod string
		}
	}
	Kuboard struct {
		UpdateImageUrl string `mapstructure:"update_image_url"`
		Username       string
	}
}

func Init() {
	Data = &AppConfig{
		DockerRepo: struct {
			Host       string `mapstructure:"host"`
			ProjectEnv struct {
				Dev  string
				Test string
				Prod string
			}
		}{
			Host: "xxxx.xxxx.cn:9043",
			ProjectEnv: struct {
				Dev  string
				Test string
				Prod string
			}{
				Dev:  "k8stest",
				Test: "k8stest",
				Prod: "k8prod",
			},
		},
		Kuboard: struct {
			UpdateImageUrl string `mapstructure:"update_image_url"`
			Username       string
		}{
			UpdateImageUrl: "http://172.16.100.22:9070/kuboard-api/cluster/k8s-test/kind/CICDApi/admin/resource/updateImageTag",
			Username:       "admin",
		},
	}
}
