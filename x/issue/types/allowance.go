package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Allowance struct {
	Amount  sdk.Int `json:"amount"`
	Spender string  `json:"spender"`
}

func NewAllowance(amount sdk.Coin, spender sdk.AccAddress) Allowance {
	return Allowance{amount.Amount, spender.String()}
}

func (a Allowance) String() string {
	return fmt.Sprintf(`%s:%s`, a.Spender, a.Amount)
}
