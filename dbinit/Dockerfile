
FROM postgres:latest as postgres
ENV POSTGRES_PASSWORD=${DB_PASSWORD} \
    POSTGRES_USER=${DB_USER} \
    POSTGRES_DB=${DB_NAME}
FROM builder
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 3333
CMD ["./main"]