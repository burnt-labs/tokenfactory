// Package utils preserves the one removed go-ethereum type referenced by
// cosmos/evm v0.6. The tokenfactory interchain tests do not enable Verkle trees,
// and cosmos/evm always returns a nil *PointCache, so no implementation is
// required after upgrading go-ethereum to the patched v1.17 API.
package utils

type PointCache struct{}
