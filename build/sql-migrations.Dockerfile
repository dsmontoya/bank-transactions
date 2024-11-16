FROM alpine:3.18

RUN apk add \
    curl \
    jq

# Install Goose
COPY Makefile /Makefile
RUN curl -fsSL https://raw.githubusercontent.com/pressly/goose/master/install.sh |\
    sh -s $(cat Makefile| grep "GOOSE_VERSION := " | awk '{print $3}')

# Copy sql migrations
WORKDIR /sql
COPY sql/ /sql
COPY sql-migrations-entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
