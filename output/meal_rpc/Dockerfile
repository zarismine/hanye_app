FROM alpine:latest
RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata
COPY output/meal_rpc/meal.yaml meal.yaml
COPY output/meal_rpc/meal_rpc meal_rpc

ENTRYPOINT ["/meal_rpc", "-f", "meal.yaml"]