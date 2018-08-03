echo "start to build app_server"
go build
echo "build app_server success!"
cd bin/
rm -rf app_server

cd ../
cp app_server bin/
cd bin/

# kill old app_server
ps -efww|grep -w 'app_server'|grep -v grep|cut -c 9-15|xargs kill -9

# prase log name
DATEYMD=$(date +%Y%m%d)
DATEHMS=$(date +%H%M%S)
value1="alog_all_"
value2=$value1$DATEYMD"_"$DATEHMS
value3=$value2".log"

# run server
nohup ./app_server -alsologtostderr > $value3 2>&1 &

echo $value3
echo "press any key to continue"
