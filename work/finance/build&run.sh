echo "start to build web_server"
go build
echo "build web_server success!"
cd bin/
rm -rf web_server

cd ../
cp web_server bin/
cd bin/


# kill old web_server

ps -efww|grep -w 'web_server'|grep -v grep|cut -c 9-15|xargs kill -9

DATEYMD=$(date +%Y%m%d)
DATEHMS=$(date +%H%M%S)
value1="web_server_log_all_"
value2=$value1$DATEYMD"_"$DATEHMS
value3=$value2".log"

nohup ./web_server -alsologtostderr > $value3 2>&1 &

echo $value3
echo "press any key to continue"