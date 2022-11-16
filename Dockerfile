#https://fabianlee.org/2020/01/26/golang-using-multi-stage-builds-to-create-clean-docker-images/

# builder image
FROM golang:1.18-alpine as builder
RUN mkdir /build
COPY . /build
WORKDIR /build

WORKDIR /build/main
RUN CGO_ENABLED=0 GOOS=linux go build -a -o filewatcher .


# generate clean, final image for end users
FROM alpine:3.16

WORKDIR /opt/smartcam/

COPY --from=builder /build/main/filewatcher /usr/bin/filewatcher

# executable
ENTRYPOINT [ "filewatcher" ]

# arguments that can be overridden
#CMD [ ]
