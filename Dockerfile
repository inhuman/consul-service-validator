FROM alpine
LABEL maintainer="msgexec@gmail.com"
COPY ./bin/validator /usr/local/bin
RUN chmod +x /usr/local/bin/validator
RUN mkdir /to_validate
ENTRYPOINT /usr/local/bin/validator