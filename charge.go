package coinbase

import "time"

type ChargeRequest struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	PricingType string            `json:"pricing_type"`
	LocalPrice  *LocalPrice       `json:"local_price,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
	RedirectURL string            `json:"redirect_url,omitempty"`
	CancelURL   string            `json:"cancel_url,omitempty"`
}

type StatusUpdate struct {
	Time    time.Time `json:"time"`
	Status  string    `json:"status"`
	Context string    `json:"context,omitempty"`
}

type CurrencyAmount struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}
type PricingInformation struct {
	Local    CurrencyAmount `json:"local"`
	Bitcoin  CurrencyAmount `json:"bitcoin"`
	Ethereum CurrencyAmount `json:"ethereum"`
}

type Payment struct {
	Network       string `json:"network"`
	TransactionID string `json:"transaction_id"`
	Status        string `json:"status"`
	Value         struct {
		Local  CurrencyAmount `json:"local"`
		Crypto CurrencyAmount `json:"crypto"`
	} `json:"value"`
	Block struct {
		Height                   int    `json:"height"`
		Hash                     string `json:"hash"`
		ConfirmationsAccumulated int    `json:"confirmations_accumulated"`
		ConfirmationsRequired    int    `json:"confirmations_required"`
	} `json:"block"`
}

type Charge struct {
	ID          string             `json:"id"`
	Resource    string             `json:"resource"` // <-- this will be "charge"
	Code        string             `json:"code"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	LogoURL     string             `json:"logo_url"`
	HostedURL   string             `json:"hosted_url"`
	CreatedAt   time.Time          `json:"created_at"`
	ExpiresAt   time.Time          `json:"expires_at"`
	ConfirmedAt time.Time          `json:"confirmed_at"`
	Checkout    Checkout           `json:"checkout"`
	Timeline    []StatusUpdate     `json:"timeline"`
	Metadata    map[string]string  `json:"metadata"`
	PricingType string             `json:"pricing_type"`
	Pricing     PricingInformation `json:"pricing"`
	Payments    []Payment          `json:"payments"`
	Addresses   struct {
		Bitcoin  string `json:"bitcoin"`
		Ethereum string `json:"ethereum"`
	} `json:"addresses"`
}

type ListChargesParams struct {
	Pagination *Pagination
}

type ListChargesResponse struct {
	Pagination *Pagination `json:"pagination"`
	Data       []Charge    `json:"data"`
}

type RetrieveChargeResponse struct {
	Data Charge `json:"data"`
}

type CreateChargeResponse struct {
	Data Charge `json:"data"`
}

type UpdateChargeResponse struct {
	Data Charge `json:"data"`
}
