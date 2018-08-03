echo "start to build server"
go build
echo "build server success!"

# kill old server
ps -efww|grep -w 'blockchain_server'|grep -v grep|cut -c 9-15|xargs kill -9

# prase log name
DATEYMD=$(date +%Y%m%d)
DATEHMS=$(date +%H%M%S)
value1="blockchain_server"
value2=$value1$DATEYMD"_"$DATEHMS
value3=$value2".log"

# run server
nohup ./blockchain_server -alsologtostderr > $value3 2>&1 &

echo $value3
echo "press any key to continue"
