#!/bin/bash
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJ_DIR="$(readlink -f "$SCRIPT_DIR/..")"
cd "$PROJ_DIR"

PROJECT_NAME="${PROJECT_NAME:-tars}"
VERSION="${VERSION:-$(cat "$PROJ_DIR/VERSION")}"

run() {
	image="wexel/$PROJECT_NAME-$1"
	name="tars_$1"

	docker run \
		--rm \
		--name "$name" \
		"$image:$VERSION"
}

run bot
