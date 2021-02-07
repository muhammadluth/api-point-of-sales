FROM golang:alpine as build-env

LABEL maintainer="Muhammad Luthfi <muhammadluthfi059@gmail.com>"

ARG APP_NAME=api-point-of-sales

RUN mkdir /app
ADD . /app/

COPY .env /app

RUN apk add --no-cache tzdata
ENV TZ Asia/Jakarta

WORKDIR /app
RUN go build -mod=vendor -o ${APP_NAME} .

FROM alpine
WORKDIR /app
COPY --from=build-env /app/${APP_NAME}  /app/${APP_NAME}
COPY --from=build-env /app/.env         /app/.env
EXPOSE 8081

RUN apk add --no-cache tzdata
ENV TZ Asia/Jakarta

ENTRYPOINT ["/app/api-point-of-sales"]
