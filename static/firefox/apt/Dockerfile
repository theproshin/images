FROM dev-image-store.tensor.ru/sbis-testing/browsers/base:7.4.2

ARG VERSION
ARG PACKAGE=firefox
ARG PPA

LABEL browser=$PACKAGE:$VERSION

RUN  \
        ( [ "$PPA" != "" ] && \
            apt-get update && \
            apt-get install -y software-properties-common && \
            apt-get update && \
            add-apt-repository -y $PPA \
        ) || true && \
        apt-get update && \
        apt-get -y --no-install-recommends install iproute2 libavcodec60 $PACKAGE=$VERSION && \
        ( [ "$PACKAGE" != "firefox" ] && ln /usr/bin/$PACKAGE /usr/bin/firefox ) || true && \
        firefox --version && \
        rm -Rf /tmp/* && rm -Rf /var/lib/apt/lists/*

RUN mkdir -p /usr/lib/$PACKAGE/defaults/pref && \
echo -e '\
pref("app.normandy.enabled", false);\n\
pref("app.normandy.api_url", "");\n\
pref("app.normandy.first_run", false);\n\
pref("services.settings.server", "");\n\
pref("services.settings.poll_interval", 0);\n\
pref("browser.region.network.url", "");\n\
pref("browser.region.update.enabled", false);\n\
pref("browser.newtabpage.activity-stream.feeds.topsites", false);\n\
pref("browser.newtabpage.activity-stream.feeds.section.topstories", false);\n\
pref("browser.translations.enable", false);\n\
pref("datareporting.healthreport.uploadEnabled", false);\n\
pref("toolkit.telemetry.enabled", false);\
' > /usr/lib/$PACKAGE/defaults/pref/disable-nimbus.js
