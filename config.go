package hotels

import "time"

type Config struct {
	ServerAddr string

	FetchInterval   time.Duration
	ProducerTimeout time.Duration
	HTTPTimeout     time.Duration

	AcmeSupplierURL       string
	PatagoniaSupplierURL  string
	PaperfliesSupplierURL string
}

func NewConfig() Config {
	return Config{
		ServerAddr: ":8080",

		FetchInterval:   15 * time.Minute,
		ProducerTimeout: 30 * time.Second,
		HTTPTimeout:     10 * time.Second,

		AcmeSupplierURL:       "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/acme",
		PatagoniaSupplierURL:  "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/patagonia",
		PaperfliesSupplierURL: "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/paperflies",
	}
}
