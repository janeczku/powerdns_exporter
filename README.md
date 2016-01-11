# PowerDNS Exporter

[![Travis](https://img.shields.io/travis/janeczku/powerdns_exporter.svg)](https://travis-ci.org/janeczku/powerdns_exporter)

[PowerDNS](https://www.powerdns.com/) exporter for [Prometheus](http://prometheus.io/)

Periodically scrapes metrics via the [PowerDNS HTTP-API](https://doc.powerdns.com/md/httpapi/README/) and exports them via HTTP/JSON for consumption by Prometheus.

#### The following PowerDNS products are supported
* [Authoritative Server](https://www.powerdns.com/auth.html)
* [Recursor](https://www.powerdns.com/recursor.html)
* [Dnsdist](http://dnsdist.org/) (coming soon)

---

## Flags

Name | Description | Default
---- | ---- | ----
listen-address | Host:Port pair to run exporter on | `:9130`
metric-path | Path under which to expose metrics for Prometheus | `/metrics`
api-url | Base-URL of PowerDNS authoritative server/recursor API | `http://localhost:8001/`
api-key | PowerDNS API Key | `-`

## Installation

Typical way of installing in Go should work.

```
go install
```

A Makefile is provided in case you find a need for it.

## Getting Started

```bash
go run powerdns_exporter [flags]
```

Show help:

```bash
go run powerdns_exporter --help
```

The `api-url` flag value should have this format:

* PowerDNS server/recursor 3.x: `http://host:port/`
* PowerDNS server/recursor 4.x: `http://host:port/api/v1`

[See here](https://doc.powerdns.com/md/httpapi/README/) for the required configuration options to enable the built-in API server in PowerDNS.

## Docker

To run the PowerDNS exporter as a Docker container, run:

    $ docker run -p 9130:9130 janeczku/powerdns-exporter -api-url="http://host:port/" -api-key="YOUR_API_KEY"
