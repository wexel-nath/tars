#!/bin/bash
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJ_DIR="$(readlink -f "$SCRIPT_DIR/..")"
cd "$PROJ_DIR"

PROJECT_NAME="${PROJECT_NAME:-tars}"
export VERSION="${VERSION:-$(cat "$PROJ_DIR/VERSION")}"

# common
export FEE_PERCENTAGE='0.0025'
export START_DATE='2018-01-01T01:00:00Z'
export END_DATE='2018-01-03T01:00:00Z'
export POSITION_SIZE='1000'

# SimpleBot
export POSITION_ENTER='0.985'
export POSITION_TARGET='1.015'

# local fake-btc-market values
#export MARKET_BASE_URL='https://172.21.0.1/fake-btc-markets'
#export MARKET_ID='BTC-USD'
#export TICKER_DELTA='60'

# prod fake-btc-market values
export MARKET_BASE_URL='https://api.getwexel.com/fake-btc-markets'
export MARKET_ID='ETH-USD'
export TICKER_DELTA='10'

compose() {
	docker-compose \
		--file "$PROJ_DIR/docker/docker-compose.yml" \
		--project-directory "$PROJ_DIR/docker" \
		--project-name "$PROJECT_NAME" \
		--compatibility \
		"$@"
}

compose up --remove-orphans --no-build db-init
compose run --rm bot
