FROM titan
WORKDIR /titangrm
COPY ./titan-grm /usr/local/bin
COPY ./DataWorker /usr/local/bin

EXPOSE 8080

ENV GRM_SERVER=data-importer \
    GRM_SERVER_ADDRESS=:8080 \
    GRM_SERVER_NAMESPACE=titangrm \
    GRM_REGISTRY_ADDRESS=consul.titangrm:8500

ENTRYPOINT titan-grm data-importer -registry_address "$GRM_REGISTRY_ADDRESS" -server_address "$GRM_SERVER_ADDRESS" "-server_namespace" "$GRM_SERVER_NAMESPACE"
