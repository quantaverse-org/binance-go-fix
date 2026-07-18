# Market Data Example

This example connects to the Binance Spot FIX market-data endpoint and subscribes to either the incremental order book or trade stream.

Create a repository-root `config.yaml`:

```bash
cp examples/market_data/config.example.yaml config.yaml
```

The settings are:

- `stream`: `orderbook` or `trade`.
- `symbols`: symbols to subscribe.
- `depth`: `1` for book ticker, or `2..5000` for incremental depth. Ignored for trade.
- `api_key`: Binance Ed25519 API key.
- `api_secret`: path to the Ed25519 private key PEM.

Run the example:

```bash
go run ./examples/market_data
```

The example connects to the production endpoint `fix-md.binance.com:9000` and runs until interrupted.
