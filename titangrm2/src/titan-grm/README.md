### grm-service默认参数
```
Name:   "server_address" 
EnvVar: "GRM_SERVER_ADDRESS",
Usage:  "Bind address for the server. :8080",

Name:   "server_namespace" 
EnvVar: "GRM_SERVER_NAMESPACE",
Usage:  "Bind address for the server. :8080",

Name:   "registry_address"
EnvVar: "GRM_REGISTRY_ADDRESS",
Usage:  "Comma-separated list of registry addresses.consul:8500",

Name:   "data_dir"
EnvVar: "GRM_DATA_DIR",
Usage:  "GRM data dir.  /opt/titangrm/data",

Name:   "config_dir"
EnvVar: "GRM_CONFIG_DIR",
Usage:  "GRM config dir. /opt/titangrm/config",

```
### k8s内部服务端口
consul : 8500
grm-api : 8441
data-manager : 8442
titan-auth : 8443
data-importer: 8444
storage-manager: 8445

### k8s服务对外端口
grm-api ：30475

### 镜像构建
docker build --no-cache -t titan-grm -f titangrm2/src/titan-grm/Dockerfile .
docker tag titan-grm:latest 192.168.1.149:5000/titan-grm
docker push 192.168.1.149:5000/titan-grm

测试：
docker run -d -e "GRM_SERVER=data-manager" -e "GRM_REGISTRY_ADDRESS=192.168.1.189:8500" titan-grm

http://192.168.1.149:7070

http://192.168.1.149:1234

192.168.1.149:31686

