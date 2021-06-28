#!/bin/bash
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJ_DIR="$(readlink -f "$SCRIPT_DIR/..")"
cd "$PROJ_DIR"

PROJECT_NAME="${PROJECT_NAME:-tars}"
VERSION="${VERSION:-$(cat "$PROJ_DIR/VERSION")}"

build() {
	image="wexel/$PROJECT_NAME-$1"
	dockerfile="./docker/Dockerfile.$1"

	docker build \
		-t "$image:$VERSION" \
		-f "$dockerfile" \
		.
}

build bot
build db-init
