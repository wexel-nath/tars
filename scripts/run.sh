#!/bin/bash
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJ_DIR="$(readlink -f "$SCRIPT_DIR/..")"
cd "$PROJ_DIR"

PROJECT_NAME="${PROJECT_NAME:-tars}"
export VERSION="${VERSION:-$(cat "$PROJ_DIR/VERSION")}"

# local fake-btc-market values
export MARKET_ID='BTC-USD'
export MARKET_BASE_URL='https://172.21.0.1/fake-btc-markets'
export INITIAL_DATE='2018-01-01T01:00:00Z'
export TICKER_DELTA='60'
export POSITION_SIZE='1000'

# prod fake-btc-market values
#export MARKET_ID='ETH-USD'
#export MARKET_BASE_URL='https://api.getwexel.com/fake-btc-markets'
#export INITIAL_DATE='2016-01-01T00:10:00Z'
#export TICKER_DELTA='10'
#export POSITION_SIZE='1000'

deploy() {
	docker stack deploy \
		--compose-file 'docker/docker-stack.yml' \
		"$PROJECT_NAME"
}

run() {
	image="wexel/$PROJECT_NAME-$1"
	name="tars_$1"

	docker run \
		-e "MARKET_ID=$MARKET_ID" \
		-e "MARKET_BASE_URL=$MARKET_BASE_URL" \
		-e "INITIAL_DATE=$INITIAL_DATE" \
		-e "TICKER_DELTA=$TICKER_DELTA" \
		-e "POSITION_SIZE=$POSITION_SIZE" \
		--rm \
		--name "$name" \
		"$image:$VERSION"
}

#run bot
deploy
