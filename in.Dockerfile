FROM {ARG_FROM}

ADD bin/{ARG_OS}_{ARG_ARCH}/{ARG_BIN} /{ARG_BIN}
COPY ./config /config

EXPOSE 8080

ENTRYPOINT ["/{ARG_BIN}"]

# Used in Makefile - `make container` step
