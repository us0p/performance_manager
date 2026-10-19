FROM golang:1.27.2 AS builder
# Sets the working directory for any RUN, CMD, ENTRYPOINT, COPY and ADD instructions that follows it.
# If a relative path is provided, it will be relative to the path of the previous WORKDIR instruction.
WORKDIR /app

# Copies from <src outside container ...> to <dest inside container, in this case /app, must be represented as directory since we have multiple files from src>
COPY go.mod go.sum ./

# Execute any commands to create a new layer on top of the current image.
# Download dependencies first so this layer is cached until go.mod/go.sum change
RUN --mount=type=cache,target=/go/pkg/mod go mod download

# Two important concepts are in play here:
# Docker build system work in layers. Whenever there's a change in a layer, that layer is rebuilt.
# Go already store caches of the downloaded modules at /go/pkg/mod.

# If go.mod or go.sum changes, the layer will be rebuilt and the contents of /go/pkg/mod from an earlier build will be discarded.
# By adding --mount option to RUN we can create filesystem mounts making BuildKit(Docker) leave that folder out of the image and preserve it between builds.
# This way, we can give the build access to go's existing package cache and only download new packages that were added.

COPY . .

# CGO_ENABLED=0 builds a fully static binary. Build image uses glibc and alpine uses musl.
# ldflags passes options to the linker, the step that assembles the final executable:
# 	-s removes the symbol table, the list of function and variable names that debugging tools use.
#   	-w removes the DWARF debug data, which maps machine code back to 
# 	source lines for debuggers like delve. By removing this you only 
# 	lose the ability to attach a debugger to that particular binary, 
# 	which you wouldn't do in production anyway.
#   	- Together they usually cut the binary size by about 20–30%.
# 	-trimpath removes the build machine's folder paths from the binary 
# 	and replaces them with module paths. The binary then doesn't reveal
# 	your build setup.
# /root/.cache/go-build is where the compiled cache is stored, giving access to this speeds go build process as only changed packages get rebuilt.
RUN --mount=type=cache,target=/go/pkg/mod  \
	--mount=type=cache,target=/root/.cache/go-build \
	CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /bin/performance_manager .

# Multi-stage build, copies only the binary once ready
FROM alpine:3.24.2

# Needed for outbound HTTPS/TLS connections
RUN apk add --no-cache ca-certificates

# By default, the process inside a container runs as root (user ID 0). 
# Containers share the host's Linux kernel. If a mistake ever gives the 
# container access to host files or the Docker socket, or a container-escape bug exists, they're effectively root on your machine.
# -D: no password, -H: no home directory; creates an unprivileged user called "app"
RUN adduser -D -H app
COPY --from=builder /bin/performance_manager /bin/performance_manager

# change to the created user. Every later instruction is executed as this user.
USER app
WORKDIR /app
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/templates ./templates

# Informs Docker that the container listens on the specified port at runtime.
# Doesn't actually publish the port. It serves as documentation between the person who builds the image and the person who runs the container.
EXPOSE 8000

# Sets the command to be executed when running a container from an image.
# Can only be one CMD instruction in a Dockerfile. If you list more than one, only the last takes effect.
CMD ["/bin/performance_manager"]
