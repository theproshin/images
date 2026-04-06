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
    printf '%s\n' \
'pref("app.normandy.enabled", false);' \
'pref("app.normandy.api_url", "");' \
'pref("app.normandy.first_run", false);' \
'pref("app.shield.optoutstudies.enabled", false);' \
'pref("app.shield.optoutstudies.enabled", false);' \
'' \
'pref("services.settings.server", "");' \
'pref("services.settings.poll_interval", 0);' \
'' \
'pref("browser.region.network.url", "");' \
'pref("browser.region.update.enabled", false);' \
'' \
'pref("browser.newtabpage.activity-stream.feeds.topsites", false);' \
'pref("browser.newtabpage.activity-stream.feeds.section.topstories", false);' \
'' \
'pref("browser.translations.enable", false);' \
'' \
'pref("datareporting.healthreport.uploadEnabled", false);' \
'pref("datareporting.policy.dataSubmissionEnabled", false);' \
'pref("toolkit.telemetry.enabled", false);' \
> /usr/lib/$PACKAGE/defaults/pref/disable-nimbus.js
