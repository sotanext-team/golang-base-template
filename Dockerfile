FROM public.ecr.aws/u0v8y0z7/golang:1.17.2 as builder
WORKDIR /app/
COPY . .
RUN rm -rf .git
# RUN git config --global --add url."git@github.com:".insteadOf "https://github.com/"
RUN export GOPRIVATE=github.com/es-hs
RUN go env -w GOPRIVATE=github.com/es-hs/*
ARG GITHUB_USER=es-hs
ENV GITHUB_USER=${GITHUB_USER}
ARG GITHUB_TOKEN
ENV GITHUB_TOKEN=${GITHUB_TOKEN}
RUN git config \
  --global \
  url."https://$GITHUB_USER:$GITHUB_TOKEN@github.com".insteadOf \
  "https://github.com"
RUN CGO_ENABLED=0 go build -o app-api .

FROM public.ecr.aws/u0v8y0z7/alpine
ARG PORT=80
ARG RUN_ARGS="-task service"
ENV RUN_ARGS=${RUN_ARGS}
EXPOSE ${PORT}
COPY --from=builder /app /app
COPY --from=builder /app/templates /templates
CMD ["sh", "-c", "/app/app-api $RUN_ARGS"]
