FROM alpine
LABEL maintainer="msgexec@gmail.com"
COPY ./bin/validator /
RUN chmod +x /validator
ENTRYPOINT /validator