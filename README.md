# 项目说明
- 本项目是项目drone打包完成后，调用此插件，将镜像更新到k8s中，k8s是使用的kuboard，项目中有CI/CD接口，调用一下就行，接口中使用的accessKey和accessSecret是环境变量传过来的

# 打包步骤
- 执行批处理DockerBuildTest.bat
- 将镜像上传到镜像仓库，在drone中配置使用即可

# drone.yaml中的配置示例
```yaml
  - name: Deploy app
    image: testhub.xxxxx.cn:9043/xxxxx/drone-k8s:latest
    settings:
      debug: true
      build_number: ${DRONE_BUILD_NUMBER}
      kuboard_accessKey: xxxxxxx
      namespace: "chain-server"
      deployment_name: "xxxxx-server"
      dashboard: kuboard
      build_repo:
        from_secret: build_repo
```
