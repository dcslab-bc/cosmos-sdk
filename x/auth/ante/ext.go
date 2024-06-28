package ante

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ExtensionOptionChecker is a function that returns true if the extension option is accepted.
type ExtensionOptionChecker func(*codectypes.Any) bool

type HasExtensionOptionsTx interface {
	GetExtensionOptions() []*codectypes.Any
	GetNonCriticalExtensionOptions() []*codectypes.Any
}

// rejectExtensionOption is the default extension check that reject all tx
// extensions.
func rejectExtensionOption(*codectypes.Any) bool {
	return false
}

// RejectExtensionOptionsDecorator is an AnteDecorator that rejects all extension
// options which can optionally be included in protobuf transactions. Users that
// need extension options should create a custom AnteHandler chain that handles
// needed extension options properly and rejects unknown ones.
type RejectExtensionOptionsDecorator struct {
	checker ExtensionOptionChecker
}

// NewRejectExtensionOptionsDecorator creates a new RejectExtensionOptionsDecorator
func NewRejectExtensionOptionsDecorator() RejectExtensionOptionsDecorator {
	return RejectExtensionOptionsDecorator{}
}

// NewExtensionOptionsDecorator creates a new antehandler that rejects all extension
// options which can optionally be included in protobuf transactions that don't pass the checker.
// Users that need extension options should pass a custom checker that returns true for the
// needed extension options.
func NewExtensionOptionsDecorator(checker ExtensionOptionChecker) sdk.AnteDecorator {
	if checker == nil {
		checker = rejectExtensionOption
	}

	return RejectExtensionOptionsDecorator{checker: checker}
}

var _ types.AnteDecorator = RejectExtensionOptionsDecorator{}

// AnteHandle implements the AnteDecorator.AnteHandle method
func (r RejectExtensionOptionsDecorator) AnteHandle(ctx types.Context, tx types.Tx, simulate bool, next types.AnteHandler) (newCtx types.Context, err error) {
	if hasExtOptsTx, ok := tx.(HasExtensionOptionsTx); ok {
		if len(hasExtOptsTx.GetExtensionOptions()) != 0 {
			return ctx, sdkerrors.ErrUnknownExtensionOptions
		}
	}

	return next(ctx, tx, simulate)
}
