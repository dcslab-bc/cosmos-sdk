package baseapp

import (
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// SimCheck defines a CheckTx helper function that used in tests and simulations.
// func (app *BaseApp) SimCheck(txEncoder sdk.TxEncoder, tx sdk.Tx) (sdk.GasInfo, *sdk.Result, error) {
func (app *BaseApp) SimCheck(txEncoder sdk.TxEncoder, tx sdk.Tx) (sdk.GasInfo, error) {
	// runTx expects tx bytes as argument, so we encode the tx argument into
	// bytes. Note that runTx will actually decode those bytes again. But since
	// this helper is only used in tests/simulation, it's fine.
	bz, err := txEncoder(tx)
	if err != nil {
		// return sdk.GasInfo{}, nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "%s", err)
		return sdk.GasInfo{}, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "%s", err)
	}
	// gasInfo, result, _, _, err := app.runTx(runTxModeCheck, bz)
	// return gasInfo, result, err
	return app.checkTx(bz, tx, false)
}

// Simulate executes a tx in simulate mode to get result and gas info.
func (app *BaseApp) Simulate(txBytes []byte) (sdk.GasInfo, *sdk.Result, error) {
	// gasInfo, result, _, _, err := app.runTx(runTxModeSimulate, txBytes)
	tx, err := app.txDecoder(txBytes)
	gasInfo, result, _, _, err := app.runTx(txBytes, tx, true)
	return gasInfo, result, err
}

func (app *BaseApp) SimDeliver(txEncoder sdk.TxEncoder, tx sdk.Tx) (sdk.GasInfo, *sdk.Result, error) {
	// See comment for Check().
	bz, err := txEncoder(tx)
	if err != nil {
		return sdk.GasInfo{}, nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "%s", err)
	}
	// gasInfo, result, _, _, err := app.runTx(runTxModeDeliver, bz)
	gasInfo, result, _, _, err := app.runTx(bz, tx, false)
	return gasInfo, result, err
}

// Context with current {check, deliver}State of the app used by tests.
func (app *BaseApp) NewContext(isCheckTx bool, header tmproto.Header) sdk.Context {
	if isCheckTx {
		return sdk.NewContext(app.checkState.ms, header, true, app.logger).
			WithMinGasPrices(app.minGasPrices)
		// WithConsensusParams(app.consensusParams)
	}

	return sdk.NewContext(app.deliverState.ms, header, false, app.logger)
}

func (app *BaseApp) NewUncachedContext(isCheckTx bool, header tmproto.Header) sdk.Context {
	return sdk.NewContext(app.cms, header, isCheckTx, app.logger)
}

func (app *BaseApp) GetContextForDeliverTx(txBytes []byte) sdk.Context {
	return app.getContextForTx(app.deliverState, txBytes)
}
