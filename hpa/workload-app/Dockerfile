ARG IMAGE=rust
ARG TAG=1.65

###########
# Builder #
###########
FROM --platform=$BUILDPLATFORM ${IMAGE}:${TAG} AS builder

RUN apt-get update && apt-get install -y \
    build-essential \
    protobuf-compiler \
    libprotobuf-dev \
    g++-x86-64-linux-gnu libc6-dev-amd64-cross \
    g++-aarch64-linux-gnu libc6-dev-arm64-cross && \
    rm -rf /var/lib/apt/lists/*

RUN rustup target add x86_64-unknown-linux-musl aarch64-unknown-linux-musl
ENV CARGO_TARGET_X86_64_UNKNOWN_LINUX_MUSL_LINKER=x86_64-linux-gnu-gcc \
    CC_x86_64_unknown_linux_musl=x86_64-linux-gnu-gcc \
    CXX_x86_64_unknown_linux_musl=x86_64-linux-gnu-g++ \
    CARGO_TARGET_AARCH64_UNKNOWN_LINUX_MUSL_LINKER=aarch64-linux-gnu-gcc \
    CC_aarch64_unknown_linux_musl=aarch64-linux-gnu-gcc \
    CXX_aarch64_unknown_linux_musl=aarch64-linux-gnu-g++ \
    CARGO_INCREMENTAL=0

# Create appuser
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "1001" \
    "appuser"

#################
# Builder AMD64 #
#################

FROM --platform=$BUILDPLATFORM builder AS build-amd64

ARG APPNAME
WORKDIR /app
COPY . .
RUN cargo build --release --target x86_64-unknown-linux-musl --bin=${APPNAME}
RUN mv /app/target/x86_64-unknown-linux-musl/release/${APPNAME} /app/${APPNAME}

#################
# Builder ARM64 #
#################

# Thanks to Ilkka Anttonen https://gitlab.com/sirile/rust-multiarch-image-test

FROM --platform=$BUILDPLATFORM builder AS build-arm64

ARG APPNAME

WORKDIR /app
COPY . .
RUN cargo build --release --target aarch64-unknown-linux-musl
RUN mv /app/target/aarch64-unknown-linux-musl/release/${APPNAME} /app/${APPNAME}

###############
# Final AMD64 #
###############

FROM scratch AS final-amd64

ARG APPNAME

# Import from builder.
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy our build
COPY --from=build-amd64 /app/${APPNAME} /usr/local/bin/app

# Use an unprivileged user.
USER appuser:appuser

ENTRYPOINT ["app"]

###############
# Final ARM64 #
###############

FROM scratch AS final-arm64

ARG APPNAME

# Import from builder.
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy our build
COPY --from=build-amd64 /app/${APPNAME} /usr/local/bin/app

# Use an unprivileged user.
USER appuser:appuser

ENTRYPOINT ["app"]

#########
# Final #
#########

FROM final-${TARGETARCH}