FROM debian:bullseye-slim as tmp
RUN apt-get update && apt-get install -y wget
WORKDIR /app
# https://github.com/mozillazg/docker-credential-acr-helper
RUN wget https://github.com/mozillazg/docker-credential-acr-helper/releases/download/v0.3.0/docker-credential-acr-helper_0.3.0_Linux_x86_64.tar.gz && \
    tar zxvf docker-credential-acr-helper_0.3.0_Linux_x86_64.tar.gz && \
    chmod +x docker-credential-acr-helper

FROM gcr.io/kaniko-project/executor:v1.9.1

COPY --from=tmp /app/docker-credential-acr-helper /usr/local/bin/

WORKDIR /app
COPY app.Dockerfile ./Dockerfile
