FROM alpine:latest
RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata
COPY output/address_api/address.yaml address.yaml
COPY output/address_api/address_api address_api

ENTRYPOINT ["/address_api", "-f", "address.yaml"]