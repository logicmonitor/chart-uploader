FROM golang:1.8-onbuild

ENV HELM_VERSION="v2.6.1"

# Helm
RUN curl -L https://storage.googleapis.com/kubernetes-helm/helm-${HELM_VERSION}-linux-amd64.tar.gz | tar -xz -C /tmp \
    && mv /tmp/linux-amd64/helm /usr/local/bin/helm \
    && chmod +x /usr/local/bin/* \
    && rm -rf /tmp/linux-amd64
