# Market Data Example

This example connects to the Binance Spot FIX market-data endpoint and subscribes to either the incremental order book or trade stream.

Create a repository-root `.env` file containing only the credentials:

```bash
cp examples/market_data/.env.example .env
```

The credential settings are:

- `BINANCE_FIX_API_KEY`: Binance Ed25519 API key.
- `BINANCE_FIX_PRIVATE_KEY_FILE`: absolute path to the Ed25519 private key PEM.

Order book (`MarketDepth=1` for book ticker, `2..5000` for incremental depth):

```bash
go run ./examples/market_data --stream orderbook --depth 10 --symbol BTCUSDT --symbol ETHUSDT
```

Trade stream:

```bash
go run ./examples/market_data --stream trade --symbol BTCUSDT
```

`--symbol` may be specified multiple times. When omitted, it defaults to `BTCUSDT`. `--depth` defaults to `2` and is ignored for the trade stream. The example connects to the production endpoint `fix-md.binance.com:9000` and runs until interrupted.
