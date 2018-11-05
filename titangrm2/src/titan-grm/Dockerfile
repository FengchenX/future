FROM golang:alpine as grm-build
WORKDIR /titangrm
COPY ./third_party /titangrm
COPY ./titangrm2 /titangrm
RUN export GOPATH=$GOPATH:/titangrm \
        && cd /titangrm \
        && export GOPATH=$GOPATH:/titangrm \
        && GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build titan-grm

FROM alpine
WORKDIR /usr/local/bin
COPY --from=grm-build /titangrm/titan-grm /usr/local/bin/
COPY --from=grm-build /titangrm/src/titan-grm/translation /usr/local/bin/translation/
RUN apk add curl

EXPOSE 8080

ENV GRM_SERVER=data-manager \
    GRM_SERVER_ADDRESS=:8080 \
    GRM_SERVER_NAMESPACE=titangrm \
    GRM_REGISTRY_ADDRESS=consul.titangrm:8500

#ENTRYPOINT ["titan-grm"]
ENTRYPOINT titan-grm "$GRM_SERVER" -registry_address "$GRM_REGISTRY_ADDRESS" -server_address "$GRM_SERVER_ADDRESS" "-server_namespace" "$GRM_SERVER_NAMESPACE"
