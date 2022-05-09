#!/bin/bash

function go-api-template {
    echo "Installing go-api-template ..."

    apt-get install jq -y

    cd /etc/go-api-template

    LATEST_VERSION=$(curl --silent "https://api.github.com/repos/james wilson/go-api-template/releases/latest" | jq '.tag_name' | sed -E 's/.*"([^"]+)".*/\1/' | tr -d v)

    curl -sL https://github.com/james wilson/go-api-template/releases/download/v{$LATEST_VERSION}/go-api-template_{$LATEST_VERSION}_Linux_x86_64.tar.gz | tar xz


    echo "[Unit]
Description=go api template
Documentation=https://github.com/james wilson/go-api-template

[Service]
ExecStart=/etc/go-api-template/go-api-template server -c /etc/go-api-template/config.prod.yml
Restart=on-failure
RestartSec=2

[Install]
WantedBy=multi-user.target" > /etc/systemd/system/go-api-template.service

    systemctl daemon-reload
    systemctl enable go-api-template.service
    systemctl start go-api-template.service

    echo "go-api-template installation done!"
}

go-api-template
