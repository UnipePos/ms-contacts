# --- Base ----
FROM golang:1.20-alpine AS base
WORKDIR $GOPATH/src/github.com/UnipePos/ms-contacts

# ---- Dependencies ----
FROM base AS dependencies
COPY . .

# ---- Build ----
FROM dependencies AS build
RUN go build -o /ms-contacts

# --- Release ----
FROM build AS image
EXPOSE 8000
CMD [ "/ms-contacts" ]
