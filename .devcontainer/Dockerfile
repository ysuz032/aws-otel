FROM mcr.microsoft.com/devcontainers/base:jammy

WORKDIR /init
COPY setup.sh post-create-command ./
COPY export-aws-credentials /usr/local/bin/

RUN chmod 755 /usr/local/bin/export-aws-credentials
RUN chmod 755 ./post-create-command
RUN chmod 744 ./setup.sh && ./setup.sh