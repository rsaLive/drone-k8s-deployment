package main

var PConfigData *PConfig

type PConfig struct {
	Debug            bool   `envName:"PLUGIN_DEBUG"`
	AccessToken      string `envName:"PLUGIN_ACCESS_TOKEN,PLUGIN_TOKEN"`
	AccessSecret     string `envName:"LUGIN_ACCESS_SECRET,PLUGIN_SECRET"`
	Lang             string `envName:"PLUGIN_LANG"`
	MessageType      string `envName:"PLUGIN_MSG_TYPE,PLUGIN_TYPE,PLUGIN_MESSAGE_TYPE"`
	CommitLink       string `envName:"DRONE_COMMIT_LINK"`
	CommitMessage    string `envName:"DRONE_COMMIT_MESSAGE"`
	CommitSha        string `envName:"DRONE_COMMIT_SHA"`
	DroneRepo        string `envName:"DRONE_REPO"`
	DroneTag         string `envName:"DRONE_TAG"`
	BuildStatus      string `envName:"DRONE_BUILD_STATUS"`
	BuildLink        string `envName:"DRONE_BUILD_LINK"`
	DronePort        string `envName:"PLUGIN_DRONE_PORT"`
	ImageName        string `envName:"PLUGIN_IMAGE_NAME"`
	ShaLink          string `envName:"PLUGIN_SHA_LINK,PLUGIN_MESSAGE_SHA_LINK"`
	Event            string `envName:"DRONE_BUILD_EVENT"`
	BuildNumber      string `envName:"PLUGIN_BUILD_NUMBER"`
	KuboardAccessKey string `envName:"PLUGIN_KUBOARD_ACCESSKEY"`
	Namespace        string `envName:"PLUGIN_NAMESPACE"`
	DeploymentName   string `envName:"PLUGIN_DEPLOYMENT_NAME"`
	Dashboard        string `envName:"PLUGIN_DASHBOARD"`
}
