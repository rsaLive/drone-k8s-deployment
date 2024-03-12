package service

import (
	"bytes"
	"drone-k8s/models"
	"drone-k8s/pkg/config"
	"fmt"
	"net/http"
)

// Exec
//
//	@Description: 执行打包
func Exec() {
	switch models.PConfigData.Dashboard {
	case "kuboard":
		KuboardUpdateImage()
	}
}

// KuboardUpdateImage
//
//	@Description: 更新deployment
func KuboardUpdateImage() {
	url := config.Data.Kuboard.UpdateImageUrl
	var projectName string
	switch models.PConfigData.CommitBranch {
	case "daiyb", "dev":
		projectName = config.Data.DockerRepo.ProjectEnv.Test
	case "prod":
		projectName = config.Data.DockerRepo.ProjectEnv.Prod
	}
	imageUrl := fmt.Sprintf("%s/%s/%s", config.Data.DockerRepo.Host, projectName, models.PConfigData.DeploymentName)
	payload := []byte(fmt.Sprintf(`{"kind":"deployments","namespace":"%s","name":"%s",
"images":{"%s":"%s:%s"}}`, models.PConfigData.Namespace, models.PConfigData.DeploymentName, imageUrl, imageUrl, models.PConfigData.BuildNumber))
	headers := map[string]string{
		"Content-Type": "application/yaml",
		"Cookie":       fmt.Sprintf("KuboardUsername=admin; KuboardAccessKey=%s", models.PConfigData.KuboardAccesskey),
	}
	fmt.Println(url)
	fmt.Println(string(payload))
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Response status code:", resp.Status)
}

func KuboardRestartWorker() {
	url := "http://172.16.100.22:9070/kuboard-api/cluster/k8s-test/kind/CICDApi/admin/resource/updateImageTag"
	payload := []byte(fmt.Sprintf(`{"kind":"deployments","namespace":"%s","name":"%s"}`, models.PConfigData.Namespace, models.PConfigData.DeploymentName))
	headers := map[string]string{
		"Content-Type": "application/yaml",
		"Cookie":       fmt.Sprintf("KuboardUsername=admin; KuboardAccessKey=%s", models.PConfigData.KuboardAccesskey),
	}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Response status code:", resp.Status)
}
