FROM ubuntu:22.04 as build

ARG WEBKIT_TAG="main"

RUN \
    apt-get update && \
    apt-get -y install --no-install-recommends ca-certificates git && \
    git clone --depth 1 --single-branch --branch "$WEBKIT_TAG" https://github.com/WebKit/WebKit.git webkit && \
    mkdir -p /opt/webkit && \
    cd webkit && \
    yes | DEBIAN_FRONTEND=noninteractive Tools/gtk/install-dependencies && \
    cmake -DPORT=GTK -DCMAKE_BUILD_TYPE=Release -DCMAKE_INSTALL_PREFIX=/opt/webkit -DUSE_WPE_RENDERER=OFF -DENABLE_MINIBROWSER=ON -DENABLE_BUBBLEWRAP_SANDBOX=OFF -DENABLE_GAMEPAD=OFF -DENABLE_SPELLCHECK=OFF -DENABLE_WAYLAND_TARGET=OFF -DUSE_OPENJPEG=OFF -GNinja && \
    ninja && \
    ninja install && \
    rm -Rf /var/lib/apt/lists/*

FROM golang:1.22-bullseye as go

COPY cmd/prism /prism

RUN \
    apt-get update && \
    apt-get install --no-install-recommends -y upx-ucl libx11-dev && \
    cd /prism && \
    GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" && \
    upx /prism/prism && \
    rm -Rf /var/lib/apt/lists/*

FROM browsers/base:7.4.2

COPY --from=build /opt/webkit /opt/webkit
COPY --from=go /prism/prism /usr/bin/

ENV LD_LIBRARY_PATH /opt/webkit/lib/:${LD_LIBRARY_PATH}

RUN \
    apt-get update && \
    apt-get -y install --no-install-recommends \
        libsoup-3.0-0 \
        libgtk-3-0 \
        libwebp7 \
        libwebpdemux2 \
        libsecret-1-0 \
        libhyphen0 \
        libwoff1 \
        libharfbuzz-icu0 \
        libgstreamer-gl1.0-0 \
        libopenjp2-7 \
        libnotify4 \
        libxslt1.1 \
        libegl1 \
        libatomic1 \
        libopengl0 && \
    ldconfig && \
    apt-get clean && \
    rm -Rf /tmp/* && rm -Rf /var/lib/apt/lists/* && \
    chmod 777 /etc/ssl/certs && \
    mkdir /tmp/ca-certificates && \
    chmod 777 /tmp/ca-certificates

COPY entrypoint.sh /

USER selenium

EXPOSE 4444
ENTRYPOINT ["/entrypoint.sh"]
