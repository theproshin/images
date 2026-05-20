ARG VERSION
FROM selenoid/dev_firefox:$VERSION

COPY geckodriver /usr/bin/
COPY selenoid /usr/bin/
COPY --chown=selenium:root browsers.json /home/selenium/
COPY entrypoint.sh /
COPY fonts/* /usr/share/fonts/truetype/
COPY firefox-user.js /opt/firefox-user.js
RUN mv /usr/bin/firefox /usr/bin/firefox-real
COPY firefox-wrapper.sh /usr/bin/firefox

RUN chmod +x /usr/bin/firefox

RUN echo "cookie-file = ~/.config/pulse/cookie" >> /etc/pulse/client.conf

USER selenium

EXPOSE 4444
ENTRYPOINT ["/entrypoint.sh"]
