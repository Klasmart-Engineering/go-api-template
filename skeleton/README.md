# ${{values.component_id}}

${{values.desciprtion}}


<p align="center">
    <img src="/assets/logo.png?v=1.0.0" width="110" />
    <h3 align="center">go api template</h3>
    <p align="center">go api template</p>
    <p align="center">
        <a href="https://github.com/james wilson/go-api-template/actions/workflows/build.yml">
            <img src="https://github.com/james wilson/go-api-template/actions/workflows/build.yml/badge.svg">
        </a>
        <a href="https://github.com/james wilson/go-api-template/actions">
            <img src="https://github.com/james wilson/go-api-template/workflows/Release/badge.svg">
        </a>
        <a href="https://github.com/james wilson/go-api-template/releases">
            <img src="https://img.shields.io/badge/Version-1.0.0-pink.svg">
        </a>
        <a href="https://goreportcard.com/report/github.com/james wilson/go-api-template">
            <img src="https://goreportcard.com/badge/github.com/james wilson/go-api-template?v=1.0.0">
        </a>
        <a href="https://godoc.org/github.com/james wilson/go-api-template">
            <img src="https://godoc.org/github.com/james wilson/go-api-template?status.svg">
        </a>
        <a href="https://github.com/james wilson/go-api-template/blob/master/LICENSE">
            <img src="https://img.shields.io/badge/LICENSE-MIT-orange.svg">
        </a>
    </p>
</p>

## Documentation

#### Linux Deployment

Download [the latest go-api-template binary](https://github.com/james wilson/go-api-template/releases). Make it executable from everywhere.

```zsh
$ export LATEST_VERSION=$(curl --silent "https://api.github.com/repos/james wilson/go-api-template/releases/latest" | jq '.tag_name' | sed -E 's/.*"([^"]+)".*/\1/' | tr -d v)

$ curl -sL https://github.com/james wilson/go-api-template/releases/download/v{$LATEST_VERSION}/go-api-template_{$LATEST_VERSION}_Linux_x86_64.tar.gz | tar xz
```

Create the configs file `config.yml` from `config.dist.yml`. Something like the following:

```yaml
# App configs
app:
  # App name
  name: ${APP_NAME:-go-api-template}
  # Env mode (dev or prod)
  mode: ${APP_MODE:-prod}
  # HTTP port
  port: ${API_PORT:-80}
  # Hostname
  hostname: ${API_HOSTNAME:-127.0.0.1}
  # TLS configs
  tls:
    status: ${API_TLS_STATUS:-off}
    crt_path: ${API_TLS_PEMPATH:-cert/server.crt}
    key_path: ${API_TLS_KEYPATH:-cert/server.key}

  # Global timeout
  timeout: ${API_TIMEOUT:-50}

  # Log configs
  log:
    # Log level, it can be debug, info, warn, error, panic, fatal
    level: ${LOG_LEVEL:-info}
    # Output can be stdout or abs path to log file /var/logs/go-api-template.log
    output: ${LOG_OUTPUT:-stdout}
    # Format can be json
    format: ${LOG_FORMAT:-json}
```

The run the `go-api-template` with `systemd`

```zsh
$ go-api-template server -c /path/to/config.yml
```

## Versioning

For transparency into our release cycle and in striving to maintain backward compatibility, Helmet is maintained under the [Semantic Versioning guidelines](https://semver.org/) and release process is predictable and business-friendly.

See the [Releases section of our GitHub project](https://github.com/james wilson/go-api-template/releases) for changelogs for each release version of Helmet. It contains summaries of the most noteworthy changes made in each release.

## Bug tracker

If you have any suggestions, bug reports, or annoyances please report them to our issue tracker at https://github.com/james wilson/go-api-template/issues

## Security Issues

If you discover a security vulnerability within Helmet, please send an email to [hello@clivern.com](mailto:hello@clivern.com)

## Contributing

We are an open source, community-driven project so please feel free to join us. see the [contributing guidelines](CONTRIBUTING.md) for more details.

## License

© 2022, james wilson. Released under [MIT License](https://opensource.org/licenses/mit-license.php).

**go api template** is authored and maintained by [@james wilson](http://github.com/james wilson).
