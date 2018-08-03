BUILD=`date +%FT%T%z`
VERSION=`svn log| sed -n 2p | awk -F "|" '{print $1}'`

LDFLAGS= -ldflags "-X main._BUILD_=${BUILD} -X main._VERSION_=r541"

all:
	go build ${LDFLAGS} -o bin/finance
