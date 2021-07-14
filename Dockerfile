FROM alpine:latest as build

# Install dependencies
RUN apk add --no-cache git npm go && \
	git clone https://github.com/xenobyter/xbVorrat && \
	echo VUE_APP_API=/api > xbVorrat/app/.env

# Build app
WORKDIR /xbVorrat/app
RUN npm install && npm run build && rm -rf /app/node_modules/


#Build go api
WORKDIR /xbVorrat
RUN go build


FROM alpine:latest
COPY --from=build /xbVorrat /

EXPOSE 8081

CMD ["./xbVorrat"]
