FROM alpine:latest
RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata
COPY output/address_rpc/address.yaml address.yaml
COPY output/address_rpc/address_rpc address_rpc

ENTRYPOINT ["/address_rpc", "-f", "address.yaml"]