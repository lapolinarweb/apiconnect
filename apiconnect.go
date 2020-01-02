// Edge Network
// (c) 2020 Edge Network technologies Ltd.

package apiconnect

import (
	"context"
	"encoding/json"
)

// Instance stores connection parameters.
type Instance struct {
	Hostname string
	Port     int
	Property string
	wallet   *Wallet
}

// Do executes a request.
func (i *Instance) Do(ctx context.Context, r Request, res interface{}) ([]byte, error) {
	bearer := i.wallet.GetBearer(ctx)
	if err := r.Initialize(i.Hostname, i.Property, bearer, i.Port); err != nil {
		return nil, err
	}

	resp, err := r.Do(ctx)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &res)
	return resp, err
}

// New returns a new instance of API.
func New(host, property, client, secret string, port int) *Instance {
	return &Instance{
		Hostname: host,
		Port:     port,
		Property: property,
		wallet:   NewWallet(host, client, secret, port),
	}
}