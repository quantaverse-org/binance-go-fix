# Market Data Example

This example sends FIX requests to the Binance Spot FIX market-data endpoint,
receives SBE responses, and subscribes to either the incremental order book or
trade stream.

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

The example connects to the production endpoint `fix-md.binance.com:9001` and runs until interrupted.

## Profiling

While the example is running, inspect the average mutex contention delay during a 30-second window with:

```bash
go tool pprof -top -mean_delay \
  'http://127.0.0.1:6060/debug/pprof/mutex?seconds=30'
```

Inspect live heap usage after forcing a GC cycle with:

```bash
go tool pprof -top -sample_index=inuse_space \
  'http://127.0.0.1:6060/debug/pprof/heap?gc=1'
```

Inspect allocations made during a 30-second window with:

```bash
go tool pprof -top -sample_index=alloc_space \
  'http://127.0.0.1:6060/debug/pprof/allocs?seconds=30'
```
