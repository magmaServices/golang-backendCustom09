FROM alpine:latest as runner
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY ./fesl-backend ./app
#  fesl/theater (client), fesl/theater (server)
EXPOSE 18270 18275 18051 18056
CMD ["./app"]
