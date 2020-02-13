package types

// distribution module event types
const (
	EventTypeIssue             = "issue"
	EventTypeApprove           = "approve"
	EventTypeIncreaseAllowance = "increase_allowance"
	EventTypeDecreaseAllowance = "decrease_allowance"
	EventTypeTransfer          = "transfer"
	EventTypeTransferFrom      = "transfer_from"
	EventTypeTransferOwnership = "transfer_ownership"
	EventTypeMint              = "mint"
	EventTypeBurn              = "burn"
	EventTypeBurnFrom          = "burn_from"

	AttributeKeyIssuer    = "issuer"
	AttributeKeyRecipient = "recipient"
	AttributeKeyOwner     = "owner"
	AttributeKeyFrom      = "from"
	AttributeKeyTo        = "to"
	AttributeKeySpender   = "spender"
	AttributeKeyDenom     = "denom"
	AttributeKeyMinter    = "minter"
	AttributeKeyBurner    = "burner"

	AttributeValueCategory = ModuleName
)
