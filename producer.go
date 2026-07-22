package hotels

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type DataProducer interface {
	Name() string
	Fetch(ctx context.Context) ([]Hotel, error)
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func fetchJSON[T any](ctx context.Context, client HTTPClient, url string) (T, error) {
	var out T

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return out, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return out, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return out, fmt.Errorf("unexpected status code %d from %s", resp.StatusCode, url)
	}

	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return out, err
	}

	return out, nil
}

type AcmeProducer struct {
	url    string
	client HTTPClient
}

func NewAcmeProducer(url string, client HTTPClient) *AcmeProducer {
	return &AcmeProducer{url: url, client: client}
}

func (p *AcmeProducer) Name() string { return "acme" }

func (p *AcmeProducer) Fetch(ctx context.Context) ([]Hotel, error) {
	raw, err := fetchJSON[[]AcmeHotel](ctx, p.client, p.url)
	if err != nil {
		return nil, err
	}

	return Transform(raw, func(h AcmeHotel) Hotel { return *h.Hotel() }), nil
}

type PatagoniaProducer struct {
	url    string
	client HTTPClient
}

func NewPatagoniaProducer(url string, client HTTPClient) *PatagoniaProducer {
	return &PatagoniaProducer{url: url, client: client}
}

func (p *PatagoniaProducer) Name() string { return "patagonia" }

func (p *PatagoniaProducer) Fetch(ctx context.Context) ([]Hotel, error) {
	raw, err := fetchJSON[[]PatagoniaHotel](ctx, p.client, p.url)
	if err != nil {
		return nil, err
	}

	return Transform(raw, func(h PatagoniaHotel) Hotel { return *h.Hotel() }), nil
}

type PaperfliesProducer struct {
	url    string
	client HTTPClient
}

func NewPaperfliesProducer(url string, client HTTPClient) *PaperfliesProducer {
	return &PaperfliesProducer{url: url, client: client}
}

func (p *PaperfliesProducer) Name() string { return "paperflies" }

func (p *PaperfliesProducer) Fetch(ctx context.Context) ([]Hotel, error) {
	raw, err := fetchJSON[[]PaperfliesHotel](ctx, p.client, p.url)
	if err != nil {
		return nil, err
	}

	return Transform(raw, func(h PaperfliesHotel) Hotel { return *h.Hotel() }), nil
}
