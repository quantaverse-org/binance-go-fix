package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"strings"
	"time"

	fix "github.com/quantaverse-org/binance-go-fix"
	"github.com/quantaverse-org/binance-go-fix/message"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
	nanoid "github.com/matoous/go-nanoid/v2"
)

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyz"

func main() {
	ctx := gctx.GetInitCtx()
	startProfiling(ctx)
	mainCmd.Run(ctx)
}

func startProfiling(ctx context.Context) {
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)

	server := &http.Server{
		Addr:              "127.0.0.1:6060",
		ReadHeaderTimeout: 3 * time.Second,
	}

	go func() {
		g.Log().Infof(ctx, "pprof listening on http://%s/debug/pprof/", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			g.Log().Errorf(ctx, "pprof server: %v", err)
		}
	}()

	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		if err := server.Shutdown(shutdownCtx); err != nil {
			g.Log().Errorf(ctx, "shutdown pprof server: %v", err)
		}
	}()
}

var mainCmd = gcmd.Command{
	Name:   "market-data",
	Usage:  "market-data",
	Brief:  "subscribe to Binance Spot FIX market data",
	Strict: true,
	Func: func(ctx context.Context, _ *gcmd.Parser) (err error) {
		symbolsConfig, err := g.Cfg().Get(ctx, "symbols")
		if err != nil {
			return fmt.Errorf("load symbols: %w", err)
		}
		if symbolsConfig == nil {
			return errors.New("symbols is required")
		}
		symbols := symbolsConfig.Strings()
		if len(symbols) == 0 {
			return errors.New("symbols cannot be empty")
		}
		for i, symbol := range symbols {
			symbols[i] = strings.ToUpper(strings.TrimSpace(symbol))
			if symbols[i] == "" {
				return errors.New("symbol cannot be empty")
			}
		}

		streamConfig, err := g.Cfg().Get(ctx, "stream")
		if err != nil {
			return fmt.Errorf("load stream: %w", err)
		}
		if streamConfig == nil {
			return errors.New("stream is required")
		}
		stream := strings.ToLower(strings.TrimSpace(streamConfig.String()))
		if stream != "orderbook" && stream != "trade" {
			return fmt.Errorf("unsupported stream %q: want orderbook or trade", stream)
		}

		apiKey, err := loadAPIKey(ctx)
		if err != nil {
			return err
		}

		config := fix.NewClientConfig(apiKey, "EXAMPLE").
			WithEnableNotify().
			WithEncodingMode(fix.EncodingModeSBE)
		client, subscription, err := fix.NewMarketClient(config)
		if err != nil {
			return fmt.Errorf("connect market data client: %w", err)
		}
		if subscription == nil {
			return errors.New("market data subscription channel is disabled")
		}
		if err = client.Run(ctx); err != nil {
			return fmt.Errorf("client failed to run: %w", err)
		}

		requestID := nanoid.MustGenerate(alphabet, 16)
		request := message.NewMarketDataRequest(requestID, message.SubscriptionRequestTypeSubscribe)
		request.Symbols = symbols
		if stream == "orderbook" {
			depthConfig, err := g.Cfg().Get(ctx, "depth")
			if err != nil {
				return fmt.Errorf("load depth: %w", err)
			}
			if depthConfig == nil {
				return errors.New("depth is required for orderbook stream")
			}
			request.MarketDepth = depthConfig.Int64()
			request.AggregatedBook = new(true)
			request.MDEntryTypes = []message.MDEntryType{
				message.MDEntryTypeBid,
				message.MDEntryTypeOffer,
			}
		} else {
			request.MDEntryTypes = []message.MDEntryType{message.MDEntryTypeTrade}
		}

		if err = client.MarketData(ctx, request); err != nil {
			return fmt.Errorf("subscribe market data: %w", err)
		}
		g.Log().Infof(ctx, "subscribed request=%s stream=%s symbols=%s",
			requestID,
			stream,
			strings.Join(symbols, ","),
		)

		defer func() {
			unsubscribe := message.NewMarketDataRequest(requestID, message.SubscriptionRequestTypeUnsubscribe)
			unsubscribe.MarketDepth = request.MarketDepth
			if unsubscribeErr := client.MarketData(ctx, unsubscribe); unsubscribeErr != nil {
				g.Log().Errorf(ctx, "unsubscribe market data: %v", unsubscribeErr)
			}
		}()

		for {
			select {
			case <-ctx.Done():
				return nil
			case update, ok := <-subscription.MarketData:
				if !ok {
					return ctx.Err()
				}
				now := time.Now()
				switch update := update.(type) {
				case *message.MarketDataSnapshot:
					g.Log().Infof(ctx, "snapshot latency=%s", now.Sub(update.SendingTime))
				case *message.MarketDataIncrementalRefresh:
					if stream == "trade" {
						for _, entry := range update.Entries {
							g.Log().Infof(ctx, "incremental latency=%s", now.Sub(entry.TransactTime))
						}
					} else {
						g.Log().Infof(ctx, "incremental latency=%s", now.Sub(update.SendingTime))
					}
				}
			}
		}
	},
}

func loadAPIKey(ctx context.Context) (*fix.ApiKey, error) {
	apiKeyConfig, err := g.Cfg().Get(ctx, "api_key")
	if err != nil {
		return nil, fmt.Errorf("load api_key: %w", err)
	}
	if apiKeyConfig == nil || strings.TrimSpace(apiKeyConfig.String()) == "" {
		return nil, errors.New("api_key is required")
	}

	apiSecretConfig, err := g.Cfg().Get(ctx, "api_secret")
	if err != nil {
		return nil, fmt.Errorf("load api_secret: %w", err)
	}
	if apiSecretConfig == nil {
		return nil, errors.New("api_secret is required")
	}
	privateKeyPEM := strings.TrimSpace(apiSecretConfig.String())
	if privateKeyPEM == "" {
		return nil, errors.New("api_secret is required")
	}
	privateKey, err := message.ParseLogonPrivateKeyPEM([]byte(privateKeyPEM))
	if err != nil {
		return nil, fmt.Errorf("parse private key: %w", err)
	}

	return &fix.ApiKey{
		UserName:   strings.TrimSpace(apiKeyConfig.String()),
		PrivateKey: privateKey,
	}, nil
}
