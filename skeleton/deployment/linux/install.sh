#!/bin/bash

function goapi {
    echo "Installing goapi ..."

    apt-get install jq -y

    cd /etc/goapi

    LATEST_VERSION=$(curl --silent "https://api.github.com/repos/james.wilson/goapi/releases/latest" | jq '.tag_name' | sed -E 's/.*"([^"]+)".*/\1/' | tr -d v)

    curl -sL https://github.com/james.wilson/goapi/releases/download/v{$LATEST_VERSION}/goapi_{$LATEST_VERSION}_Linux_x86_64.tar.gz | tar xz


    echo "[Unit]
Description=goapi
Documentation=https://github.com/james.wilson/goapi

[Service]
ExecStart=/etc/goapi/goapi server -c /etc/goapi/config.prod.yml
Restart=on-failure
RestartSec=2

[Install]
WantedBy=multi-user.target" > /etc/systemd/system/goapi.service

    systemctl daemon-reload
    systemctl enable goapi.service
    systemctl start goapi.service

    echo "goapi installation done!"
}

goapi
