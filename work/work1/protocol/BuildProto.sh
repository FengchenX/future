#!/bin/bash
./protoc --go_out=plugins=grpc:../server/protocol ./*.proto
read -p "Press any key to continue." var