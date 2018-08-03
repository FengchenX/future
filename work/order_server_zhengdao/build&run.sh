echo "start to build three_server"
go build
echo "build three_server success!"
cd bin/
rm -rf three_server

cd ../
cp three_server bin/
cd bin/


# kill old three_server

ps -efww|grep -w 'three_server'|grep -v grep|cut -c 9-15|xargs kill -9

DATEYMD=$(date +%Y%m%d)
DATEHMS=$(date +%H%M%S)
value1="log_three_server_all_"
value2=$value1$DATEYMD"_"$DATEHMS
value3=$value2".log"

nohup ./three_server -alsologtostderr > $value3 2>&1 &

echo $value3
echo "press any key to continue"
