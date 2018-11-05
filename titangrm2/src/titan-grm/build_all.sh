#!/bin/bash

sudo docker images|grep none|awk '{print $3}'|xargs sudo docker rmi

# titan
docker build --no-cache -t titan .

# titan-grm
cd /home/workspace/titangrm2/Services
svn up third_party
svn up titangrm2
docker build --no-cache -t 192.168.1.149:5000/titan-grm --file titangrm2/src/titan-grm/Dockerfile .
docker push 192.168.1.149:5000/titan-grm

rancher kubectl delete pods -l app=data-manager -n titangrm-dev
rancher kubectl delete pods -l app=grm-api -n titangrm-dev
rancher kubectl delete pods -l app=storage-manager -n titangrm-dev
rancher kubectl delete pods -l app=titan-auth -n titangrm-dev

# titan-auth
rancher kubectl create -f titan-auth-deploy.yaml

# data-manager
rancher kubectl create -f data-manager-deploy.yaml

# storage-manager
rancher kubectl create -f storage-manager-deploy.yaml

# data-importer
export GOPATH=/home/workspace/TitanGRM2.0/src/Services/third_party:/home/workspace/TitanGRM2.0/src/Services/titangrm2
export GOPATH=/home/build_images/third_party:/home/build_images/titangrm2

cd /home/workspace/TitanGRM2.0/src/Services
export GOPATH=$PWD/third_party:$PWD/titangrm2
go build titan-grm
mv titan-grm /sdc/grm/images/data-importer/
cd /sdc/grm/images/data-importer/
docker build --no-cache -t 192.168.1.149:5000/data-importer .
docker push 192.168.1.149:5000/data-importer
rancher kubectl delete pods -l app=data-importer -n titangrm-dev

#docker run -d  -e "GRM_REGISTRY_ADDRESS=192.168.1.149:31720" -p 8444:8080 -v /sdc/grm/config/FileSysDomain_linux.json:/opt/titangrm/config/FileSysDomain.json -v /sdc/grm/data:/opt/titangrm/data titangrm/data-importer

#docker push 192.168.1.149:5000/data-importer

# grm-api
# docker run -d -e "GRM_SERVER=grm-api" -e "GRM_REGISTRY_ADDRESS=192.168.1.149:31720" -p 8441:8080 titan-grm
rancher kubectl create -f grm-api-deploy.yaml

#####################################################################

/usr/local/bin/DataWorker --mode load --jobid load_bcecf4e2b7fc11e8b79806df3d2386c9 --load_files /sdc/grm/config/dataworker/load/load_bcecf4e2b7fc11e8b79806df3d2386c9/file_list.txt --datatype Shape --data_set 123 --sysdb postgres://postgres:otitan123@192.168.1.149:31771/TitanCloud.System --metadb postgres://postgres:otitan123@192.168.1.149:31771/TitanCloud.Meta --datadb postgres://postgres:otitan123@192.168.1.149:31771/TitanCloud.Data --esdb http://192.168.1.149:9200 --workdir /sdc/grm/config/dataworker/load/load_bcecf4e2b7fc11e8b79806df3d2386c9

shp2pgsql -s 4267 -k -W UTF-8 /sdc/grm/data/disk1/Feature/airport/airports.shp cf34e857f9154e9696513553fe9ed8eb > /sdc/grm/config/dataworker/load/load_bcecf4e2b7fc11e8b79806df3d2386c9/airports.sql public.cf34e857f9154e9696513553fe9ed8eb | psql -d shp2pgsqldemo -U gisdb -W




