FROM golang:1.19 as build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gin-demo .

FROM scratch as prod
COPY --from=build /app/gin-demo /app/
COPY conf /app/conf
WORKDIR /app
EXPOSE 8080
CMD ["./gin-demo"]