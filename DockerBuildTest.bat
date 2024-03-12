go env -w GOOS=linux
cd ./build
del drone-k8s
cd ..
go build  -ldflags "-s -w" -o ./build/drone-k8s .
docker rmi drone-k8s
docker rmi testhub.szjixun.cn:9043/public/drone-k8s
docker build . -f .\Dockerfile -t drone-k8s
docker tag drone-k8s testhub.szjixun.cn:9043/public/drone-k8s
docker push testhub.szjixun.cn:9043/public/drone-k8s
pause