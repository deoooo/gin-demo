FROM golang:1.19 as build
WORKDIR /app
RUN go install github.com/swaggo/swag/cmd/swag@latest
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gin-demo .
RUN swag init --parseDependency --parseInternal ./routers/*.go --output ./docs --outputTypes go

FROM scratch as prod
COPY --from=build /app/gin-demo /app/
COPY --from=build /app/docs /app/docs
COPY conf /app/conf
WORKDIR /app
EXPOSE 8080
CMD ["./gin-demo"]