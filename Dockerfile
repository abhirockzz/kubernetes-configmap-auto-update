FROM golang as build-stage
WORKDIR /go/
COPY src/ /go/src
RUN cd /go/src && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o configmaptestapp

FROM alpine
COPY --from=build-stage /go/src/configmaptestapp /
#RUN mkdir -p /data/store
EXPOSE 8080
CMD ["/configmaptestapp"]