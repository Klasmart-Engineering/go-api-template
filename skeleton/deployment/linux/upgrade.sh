#!/bin/bash

function goapi {
    echo "Upgrade goapi ..."

    cd /etc/goapi
    mv config.prod.yml config.back.yml

    LATEST_VERSION=$(curl --silent "https://api.github.com/repos/james.wilson/goapi/releases/latest" | jq '.tag_name' | sed -E 's/.*"([^"]+)".*/\1/' | tr -d v)

    curl -sL https://github.com/james.wilson/goapi/releases/download/v{$LATEST_VERSION}/goapi_{$LATEST_VERSION}_Linux_x86_64.tar.gz | tar xz

    rm config.prod.yml
    mv config.back.yml config.prod.yml

    systemctl restart goapi

    echo "goapi upgrade done!"
}

goapi
