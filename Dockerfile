FROM golang:1.19 as build
ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE on
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gin_demo .

FROM scratch as prod
WORKDIR /app
COPY --from=build /app /app
COPY conf /app/conf
EXPOSE 8080
ENTRYPOINT ["./gin_demo"]