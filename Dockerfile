FROM golang:1.18-alpine AS build
ARG APP_VERSION
WORKDIR /go/src/project/
COPY . /go/src/project/
RUN go build -ldflags "-X main.Version=${APP_VERSION}" -o /bin/whoami ./main.go
FROM alpine:3.5
COPY --from=build /bin/whoami /bin/whoami
EXPOSE 8080
CMD ["/bin/whoami"]
