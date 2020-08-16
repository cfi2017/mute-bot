FROM alpine AS build

# install dependencies
RUN apk update && apk add --no-cache git ca-certificates tzdata upx && update-ca-certificates

# don't use root
ENV USER=hj
ENV UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

WORKDIR /code

COPY mute-bot .

# use upx to reduce binary size even further
RUN upx ./mute-bot

# we use "scratch" image to run go service
# the scratch image "doesn't contain anything"
FROM scratch

COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /etc/group /etc/group

WORKDIR /app
COPY --from=build /code/mute-bot /app/mute-bot
COPY --from=build /code/config/mute-bot.* .
USER hj:hj

ENTRYPOINT ["/app/mute-bot"]