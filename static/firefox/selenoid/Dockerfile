ARG VERSION
FROM selenoid/dev_firefox:$VERSION

COPY geckodriver /usr/bin/
COPY entrypoint.sh /
COPY fonts/* /usr/share/fonts/truetype/

RUN echo "cookie-file = ~/.config/pulse/cookie" >> /etc/pulse/client.conf

USER selenium

EXPOSE 4444
ENTRYPOINT ["/entrypoint.sh"]
