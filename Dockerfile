FROM gcr.io/distroless/base
COPY purslane /
ENTRYPOINT ["/purslane"] 