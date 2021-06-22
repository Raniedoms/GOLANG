FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .

WORKDIR /dist

RUN cp -R /build/* .

EXPOSE 3000

#CMD ["ls"]
CMD ["/dist/ico"]
