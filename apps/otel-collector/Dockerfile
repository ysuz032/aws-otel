FROM amazon/aws-otel-collector:latest

COPY src/config.yaml /etc/otel-agent-config.yaml

# Expose port 1777 for pprof extension
EXPOSE 1777

# Expose port 55679 for zpages extension
EXPOSE 55679

# Expose port 4317 for OTLP receiver
EXPOSE 4317

# Expose port 13133 for health_check
EXPOSE 13133

CMD ["--config=/etc/otel-agent-config.yaml"]