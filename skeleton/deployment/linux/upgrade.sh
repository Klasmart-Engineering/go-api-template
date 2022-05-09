#!/bin/bash

function go-api-template {
    echo "Upgrade go-api-template ..."

    cd /etc/go-api-template
    mv config.prod.yml config.back.yml

    LATEST_VERSION=$(curl --silent "https://api.github.com/repos/james wilson/go-api-template/releases/latest" | jq '.tag_name' | sed -E 's/.*"([^"]+)".*/\1/' | tr -d v)

    curl -sL https://github.com/james wilson/go-api-template/releases/download/v{$LATEST_VERSION}/go-api-template_{$LATEST_VERSION}_Linux_x86_64.tar.gz | tar xz

    rm config.prod.yml
    mv config.back.yml config.prod.yml

    systemctl restart go-api-template

    echo "go-api-template upgrade done!"
}

go-api-template
