FROM jaegertracing/all-in-one:latest

COPY src/jaeger-ui.json /etc/jaeger/jaeger-ui.json

CMD ["--query.ui-config","/etc/jaeger/jaeger-ui.json"]