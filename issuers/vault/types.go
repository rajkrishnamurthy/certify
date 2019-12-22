package vault

import (
	"context"
	"strings"
	"time"

	"github.com/hashicorp/vault/api"
)

// AuthMethod defines the interface required to implement
// custom authentication against the Vault server.
type AuthMethod interface {
	SetToken(context.Context, *api.Client) error
}

// ConstantToken implements AuthMethod with a constant token
type ConstantToken string

// SetToken sets the clients token to the constant token value.
func (c ConstantToken) SetToken(_ context.Context, cli *api.Client) (error) {
	cli.SetToken(string(c))
	return nil
}


// https://www.vaultproject.io/api/secret/pki/index.html#parameters-14
type csrOpts struct {
	CSR               string    `json:"csr"`
	CommonName        string    `json:"common_name"`
	ExcludeCNFromSANS bool      `json:"exclude_cn_from_sans"`
	Format            string    `json:"format"`
	URISans           otherSans `json:"uri_sans,omitempty"`
	OtherSans         otherSans `json:"other_sans,omitempty"`
	TimeToLive        ttl       `json:"ttl,omitempty"`
}

type otherSans []string

func (o otherSans) MarshalJSON() ([]byte, error) {
	return []byte(`"` + strings.Join(o, ",") + `"`), nil
}

type ttl time.Duration

func (t ttl) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Duration(t).String() + `"`), nil
}
