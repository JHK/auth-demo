# Auth Plugin

Demo Plugin for Traefik. Requests will forward a cookie to a given
authentication service. The response will be sent sent as header to the
downstream service.

## Usage

This repository is set up in a way such that you can run the plugin in a Traefik
docker container:

```bash
# run traefik with the plugin enabled
make run

# run a opentelemetry process traefik will send data to
make otel

# perform a request against traefik
curl -I --header "Cookie: login=foo" localhost
```
