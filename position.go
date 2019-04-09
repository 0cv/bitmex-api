package bitmex

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Position Summary of Open and Closed Positions
// swagger:model Position
type Position struct {

	// account
	// Required: true
	Account *int64 `json:"account"`

	// avg cost price
	AvgCostPrice float64 `json:"avgCostPrice,omitempty"`

	// avg entry price
	AvgEntryPrice float64 `json:"avgEntryPrice,omitempty"`

	// bankrupt price
	BankruptPrice float64 `json:"bankruptPrice,omitempty"`

	// break even price
	BreakEvenPrice float64 `json:"breakEvenPrice,omitempty"`

	// commission
	Commission float64 `json:"commission,omitempty"`

	// cross margin
	CrossMargin bool `json:"crossMargin,omitempty"`

	// currency
	// Required: true
	Currency *string `json:"currency"`

	// current comm
	CurrentComm int64 `json:"currentComm,omitempty"`

	// current cost
	CurrentCost int64 `json:"currentCost,omitempty"`

	// current qty
	CurrentQty int64 `json:"currentQty,omitempty"`

	// current timestamp
	// Format: date-time
	CurrentTimestamp strfmt.DateTime `json:"currentTimestamp,omitempty"`

	// deleverage percentile
	DeleveragePercentile float64 `json:"deleveragePercentile,omitempty"`

	// exec buy cost
	ExecBuyCost int64 `json:"execBuyCost,omitempty"`

	// exec buy qty
	ExecBuyQty int64 `json:"execBuyQty,omitempty"`

	// exec comm
	ExecComm int64 `json:"execComm,omitempty"`

	// exec cost
	ExecCost int64 `json:"execCost,omitempty"`

	// exec qty
	ExecQty int64 `json:"execQty,omitempty"`

	// exec sell cost
	ExecSellCost int64 `json:"execSellCost,omitempty"`

	// exec sell qty
	ExecSellQty int64 `json:"execSellQty,omitempty"`

	// foreign notional
	ForeignNotional float64 `json:"foreignNotional,omitempty"`

	// gross exec cost
	GrossExecCost int64 `json:"grossExecCost,omitempty"`

	// gross open cost
	GrossOpenCost int64 `json:"grossOpenCost,omitempty"`

	// gross open premium
	GrossOpenPremium int64 `json:"grossOpenPremium,omitempty"`

	// home notional
	HomeNotional float64 `json:"homeNotional,omitempty"`

	// indicative tax
	IndicativeTax int64 `json:"indicativeTax,omitempty"`

	// indicative tax rate
	IndicativeTaxRate float64 `json:"indicativeTaxRate,omitempty"`

	// init margin
	InitMargin int64 `json:"initMargin,omitempty"`

	// init margin req
	InitMarginReq float64 `json:"initMarginReq,omitempty"`

	// is open
	IsOpen bool `json:"isOpen,omitempty"`

	// last price
	LastPrice float64 `json:"lastPrice,omitempty"`

	// last value
	LastValue int64 `json:"lastValue,omitempty"`

	// leverage
	Leverage float64 `json:"leverage,omitempty"`

	// liquidation price
	LiquidationPrice float64 `json:"liquidationPrice,omitempty"`

	// long bankrupt
	LongBankrupt int64 `json:"longBankrupt,omitempty"`

	// maint margin
	MaintMargin int64 `json:"maintMargin,omitempty"`

	// maint margin req
	MaintMarginReq float64 `json:"maintMarginReq,omitempty"`

	// margin call price
	MarginCallPrice float64 `json:"marginCallPrice,omitempty"`

	// mark price
	MarkPrice float64 `json:"markPrice,omitempty"`

	// mark value
	MarkValue int64 `json:"markValue,omitempty"`

	// open order buy cost
	OpenOrderBuyCost int64 `json:"openOrderBuyCost,omitempty"`

	// open order buy premium
	OpenOrderBuyPremium int64 `json:"openOrderBuyPremium,omitempty"`

	// open order buy qty
	OpenOrderBuyQty int64 `json:"openOrderBuyQty,omitempty"`

	// open order sell cost
	OpenOrderSellCost int64 `json:"openOrderSellCost,omitempty"`

	// open order sell premium
	OpenOrderSellPremium int64 `json:"openOrderSellPremium,omitempty"`

	// open order sell qty
	OpenOrderSellQty int64 `json:"openOrderSellQty,omitempty"`

	// opening comm
	OpeningComm int64 `json:"openingComm,omitempty"`

	// opening cost
	OpeningCost int64 `json:"openingCost,omitempty"`

	// opening qty
	OpeningQty int64 `json:"openingQty,omitempty"`

	// opening timestamp
	// Format: date-time
	OpeningTimestamp strfmt.DateTime `json:"openingTimestamp,omitempty"`

	// pos allowance
	PosAllowance int64 `json:"posAllowance,omitempty"`

	// pos comm
	PosComm int64 `json:"posComm,omitempty"`

	// pos cost
	PosCost int64 `json:"posCost,omitempty"`

	// pos cost2
	PosCost2 int64 `json:"posCost2,omitempty"`

	// pos cross
	PosCross int64 `json:"posCross,omitempty"`

	// pos init
	PosInit int64 `json:"posInit,omitempty"`

	// pos loss
	PosLoss int64 `json:"posLoss,omitempty"`

	// pos maint
	PosMaint int64 `json:"posMaint,omitempty"`

	// pos margin
	PosMargin int64 `json:"posMargin,omitempty"`

	// pos state
	PosState string `json:"posState,omitempty"`

	// prev close price
	PrevClosePrice float64 `json:"prevClosePrice,omitempty"`

	// prev realised pnl
	PrevRealisedPnl int64 `json:"prevRealisedPnl,omitempty"`

	// prev unrealised pnl
	PrevUnrealisedPnl int64 `json:"prevUnrealisedPnl,omitempty"`

	// quote currency
	QuoteCurrency string `json:"quoteCurrency,omitempty"`

	// realised cost
	RealisedCost int64 `json:"realisedCost,omitempty"`

	// realised gross pnl
	RealisedGrossPnl int64 `json:"realisedGrossPnl,omitempty"`

	// realised pnl
	RealisedPnl int64 `json:"realisedPnl,omitempty"`

	// realised tax
	RealisedTax int64 `json:"realisedTax,omitempty"`

	// rebalanced pnl
	RebalancedPnl int64 `json:"rebalancedPnl,omitempty"`

	// risk limit
	RiskLimit int64 `json:"riskLimit,omitempty"`

	// risk value
	RiskValue int64 `json:"riskValue,omitempty"`

	// session margin
	SessionMargin int64 `json:"sessionMargin,omitempty"`

	// short bankrupt
	ShortBankrupt int64 `json:"shortBankrupt,omitempty"`

	// simple cost
	SimpleCost float64 `json:"simpleCost,omitempty"`

	// simple pnl
	SimplePnl float64 `json:"simplePnl,omitempty"`

	// simple pnl pcnt
	SimplePnlPcnt float64 `json:"simplePnlPcnt,omitempty"`

	// simple qty
	SimpleQty float64 `json:"simpleQty,omitempty"`

	// simple value
	SimpleValue float64 `json:"simpleValue,omitempty"`

	// symbol
	// Required: true
	Symbol *string `json:"symbol"`

	// target excess margin
	TargetExcessMargin int64 `json:"targetExcessMargin,omitempty"`

	// tax base
	TaxBase int64 `json:"taxBase,omitempty"`

	// taxable margin
	TaxableMargin int64 `json:"taxableMargin,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`

	// underlying
	Underlying string `json:"underlying,omitempty"`

	// unrealised cost
	UnrealisedCost int64 `json:"unrealisedCost,omitempty"`

	// unrealised gross pnl
	UnrealisedGrossPnl int64 `json:"unrealisedGrossPnl,omitempty"`

	// unrealised pnl
	UnrealisedPnl int64 `json:"unrealisedPnl,omitempty"`

	// unrealised pnl pcnt
	UnrealisedPnlPcnt float64 `json:"unrealisedPnlPcnt,omitempty"`

	// unrealised roe pcnt
	UnrealisedRoePcnt float64 `json:"unrealisedRoePcnt,omitempty"`

	// unrealised tax
	UnrealisedTax int64 `json:"unrealisedTax,omitempty"`

	// var margin
	VarMargin int64 `json:"varMargin,omitempty"`
}

// MarshalBinary interface implementation
func (m *Position) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Position) UnmarshalBinary(b []byte) error {
	var res Position
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
