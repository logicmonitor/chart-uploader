FROM golang:1.9 as build
WORKDIR $GOPATH/src/github.com/logicmonitor/chart-uploader
RUN git clone https://github.com/logicmonitor/chart-uploader.git ./ \
    && go get \
    && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /chart-uploader

FROM golang:1.9 as helm
ENV HELM_VERSION="v2.6.1"
RUN curl -L https://storage.googleapis.com/kubernetes-helm/helm-${HELM_VERSION}-linux-amd64.tar.gz | tar -xz -C /tmp


FROM alpine:3.6
LABEL maintainer="Jeff Wozniak <jeff.wozniak@logicmonitor.com>"

RUN apk --update add ca-certificates \
    && rm -rf /var/cache/apk/* \
    && rm -rf /var/lib/apk/*

WORKDIR /app
COPY --from=build /chart-uploader /bin
COPY --from=helm /tmp/linux-amd64/helm /usr/local/bin/helm
RUN chmod +x /usr/local/bin/*

ENTRYPOINT ["chart-uploader"]
