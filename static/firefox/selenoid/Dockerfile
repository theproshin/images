ARG VERSION
FROM selenoid/dev_firefox:$VERSION

COPY geckodriver /usr/bin/
COPY selenoid /usr/bin/
COPY --chown=selenium:root browsers.json /home/selenium/
COPY entrypoint.sh /
COPY fonts/* /usr/share/fonts/truetype/

USER root
RUN mkdir -p /usr/lib/firefox/distribution
COPY policies.json /usr/lib/firefox/distribution/policies.json

ENV XDG_RUNTIME_DIR=/tmp/runtime-firefox

RUN mkdir -p /tmp/runtime-firefox && chmod 777 /tmp/runtime-firefox

RUN echo "cookie-file = ~/.config/pulse/cookie" >> /etc/pulse/client.conf

USER selenium

EXPOSE 4444
ENTRYPOINT ["/entrypoint.sh"]
