FROM golang:1.18-stretch as builder
ENV GO111MODULE=on
COPY . /banklo-authorizer-adapter
WORKDIR /banklo-authorizer-adapter/cmd
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/application
RUN ls bin

#s Run Image
FROM scratch
COPY --from=builder /banklo-authorizer-adapter/cmd/bin/application application
EXPOSE 9999
ENTRYPOINT ["./application"]