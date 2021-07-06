#!/bin/bash
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJ_DIR="$(readlink -f "$SCRIPT_DIR/..")"
cd "$PROJ_DIR"

PROJECT_NAME="${PROJECT_NAME:-tars}"
VERSION="${VERSION:-$(cat "$PROJ_DIR/VERSION")}"

# local fake-btc-market values
marketID='BTC-USD'
marketBaseURL='https://172.21.0.1/fake-btc-markets'
initialDate='2018-01-01T01:00:00Z'
tickerDelta='60'
positionSize='1000'

# prod fake-btc-market values
#marketID='ETH-USD'
#marketBaseURL='https://api.getwexel.com/fake-btc-markets'
#initialDate='2016-01-01T00:10:00Z'
#tickerDelta='10'
#positionSize='1000'

run() {
	image="wexel/$PROJECT_NAME-$1"
	name="tars_$1"

	docker run \
		-e "MARKET_ID=$marketID" \
		-e "MARKET_BASE_URL=$marketBaseURL" \
		-e "INITIAL_DATE=$initialDate" \
		-e "TICKER_DELTA=$tickerDelta" \
		-e "POSITION_SIZE=$positionSize" \
		--rm \
		--name "$name" \
		"$image:$VERSION"
}

run bot
