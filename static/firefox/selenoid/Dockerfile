ARG VERSION
FROM selenoid/dev_firefox:$VERSION

COPY geckodriver /usr/bin/
COPY selenoid /usr/bin/
COPY --chown=selenium:root browsers.json /home/selenium/
COPY entrypoint.sh /
COPY fonts/* /usr/share/fonts/truetype/
COPY selenoid-wrapper /usr/bin/
RUN chmod +x /usr/bin/selenoid-wrapper

RUN echo "cookie-file = ~/.config/pulse/cookie" >> /etc/pulse/client.conf

USER selenium

EXPOSE 4444
ENTRYPOINT ["/entrypoint.sh"]
