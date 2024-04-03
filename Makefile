.PHONY: lint test vendor clean

export GO111MODULE=on

default: run

run:
	@docker run --rm \
		--add-host host.docker.internal:host-gateway \
		-v ./dynamic.yml:/etc/traefik/dynamic.yml \
		-v ./:/plugins-local/src/github.com/JHK/auth-demo/ \
		-p 80:80 \
		traefik:v3.0.0-rc3 \
		--entryPoints.web.address=:80 \
		--experimental.localplugins.auth-demo.modulename=github.com/JHK/auth-demo \
		--providers.file.filename=/etc/traefik/dynamic.yml \
		--log.level=INFO \
		--tracing.serviceName=auth-demo \
		--tracing.otlp.http.endpoint=http://host.docker.internal:4318/v1/traces

otel:
	@docker run --rm \
		-v ./otel-collector-config.yml:/etc/otel-collector-config.yml \
		-p 4318:4318 \
		--name otel \
		otel/opentelemetry-collector:0.96.0 \
		--config=/etc/otel-collector-config.yml
