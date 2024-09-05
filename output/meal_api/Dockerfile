FROM alpine:latest
RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata
COPY output/meal_api/meal.yaml meal.yaml
COPY output/meal_api/meal_api meal_api

ENTRYPOINT ["/meal_api", "-f", "meal.yaml"]