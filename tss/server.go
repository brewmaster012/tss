package tss

import (
	"github.com/zeta-chain/go-tss-ctx/keygen"
	"github.com/zeta-chain/go-tss-ctx/keysign"
)

// Server define the necessary functionality should be provide by a TSS Server implementation
type Server interface {
	Start() error
	Stop()
	GetLocalPeerID() string
	Keygen(req keygen.Request) (keygen.Response, error)
	KeySign(req keysign.Request) (keysign.Response, error)
}
