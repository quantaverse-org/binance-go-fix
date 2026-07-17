package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	fix "binance-go-fix"
	"binance-go-fix/message"

	"github.com/joho/godotenv"
	nanoid "github.com/matoous/go-nanoid/v2"
)

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyz"

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("load .env: %w", err)
	}

	var symbols []string
	flag.Func("symbol", "symbol to subscribe; may be specified multiple times", func(value string) error {
		symbol := strings.ToUpper(strings.TrimSpace(value))
		if symbol == "" {
			return errors.New("symbol cannot be empty")
		}
		symbols = append(symbols, symbol)
		return nil
	})
	streamFlag := flag.String("stream", "orderbook", "stream to test: orderbook or trade")
	depthFlag := flag.String("depth", "2", "order book depth: 1 for book ticker, 2-5000 for depth stream")
	flag.Parse()

	if len(symbols) == 0 {
		symbols = append(symbols, "BTCUSDT")
	}
	stream := strings.ToLower(strings.TrimSpace(*streamFlag))
	if stream != "orderbook" && stream != "trade" {
		return fmt.Errorf("unsupported stream %q: want orderbook or trade", *streamFlag)
	}

	apiKey, err := loadAPIKey()
	if err != nil {
		return err
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	config := fix.NewClientConfig(apiKey).WithEnableNotify()
	client, updating, err := fix.NewMarketClient(ctx, config)
	if err != nil {
		return fmt.Errorf("connect market data client: %w", err)
	}
	if updating == nil {
		return errors.New("market data updating channel is disabled")
	}

	requestID := nanoid.MustGenerate(alphabet, 16)
	request := message.NewMarketDataRequest(requestID, message.SubscriptionRequestTypeSubscribe)
	request.Symbols = symbols
	if stream == "orderbook" {
		aggregatedBook := true
		request.MarketDepth = strings.TrimSpace(*depthFlag)
		request.AggregatedBook = &aggregatedBook
		request.MDEntryTypes = []message.MDEntryType{
			message.MDEntryTypeBid,
			message.MDEntryTypeOffer,
		}
	} else {
		request.MDEntryTypes = []message.MDEntryType{message.MDEntryTypeTrade}
	}

	snapshots, err := client.MarketData(request)
	if err != nil {
		return fmt.Errorf("subscribe market data: %w", err)
	}
	fmt.Printf("subscribed request=%s stream=%s symbols=%s snapshots=%d\n",
		requestID,
		stream,
		strings.Join(symbols, ","),
		len(snapshots),
	)
	for _, snapshot := range snapshots {
		fmt.Printf("snapshot request=%s symbol=%s entries=%s last_update_id=%s\n",
			snapshot.MDReqID,
			snapshot.Symbol,
			snapshot.NoMDEntries,
			snapshot.LastBookUpdateID,
		)
	}

	defer func() {
		unsubscribe := message.NewMarketDataRequest(requestID, message.SubscriptionRequestTypeUnsubscribe)
		unsubscribe.MarketDepth = request.MarketDepth
		if _, unsubscribeErr := client.MarketData(unsubscribe); unsubscribeErr != nil {
			fmt.Fprintf(os.Stderr, "unsubscribe market data: %v\n", unsubscribeErr)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return nil
		case update, ok := <-updating.MarketData:
			if !ok {
				return ctx.Err()
			}
			switch update := update.(type) {
			case *message.MarketDataSnapshot:
				fmt.Printf("snapshot request=%s symbol=%s entries=%s last_update_id=%s\n",
					update.MDReqID,
					update.Symbol,
					update.NoMDEntries,
					update.LastBookUpdateID,
				)
			case *message.MarketDataIncrementalRefresh:
				fmt.Printf("incremental request=%s entries=%s\n", update.MDReqID, update.NoMDEntries)
			}
		}
	}
}

func loadAPIKey() (*fix.ApiKey, error) {
	username := strings.TrimSpace(os.Getenv("BINANCE_FIX_API_KEY"))
	if username == "" {
		return nil, errors.New("BINANCE_FIX_API_KEY is required")
	}

	privateKeyFile := strings.TrimSpace(os.Getenv("BINANCE_FIX_PRIVATE_KEY_FILE"))
	if privateKeyFile == "" {
		return nil, errors.New("BINANCE_FIX_PRIVATE_KEY_FILE is required")
	}
	privateKeyPEM, err := os.ReadFile(privateKeyFile)
	if err != nil {
		return nil, fmt.Errorf("read private key: %w", err)
	}
	privateKey, err := message.ParseLogonPrivateKeyPEM(privateKeyPEM)
	if err != nil {
		return nil, fmt.Errorf("parse private key: %w", err)
	}

	return &fix.ApiKey{
		UserName:   username,
		PrivateKey: privateKey,
	}, nil
}
