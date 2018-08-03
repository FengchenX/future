package types

import (
	"encoding/json"
)

const (
	unknownModelStr    = "unknown"
	accountModelStr    = "account"
	issueModelStr      = "issue"
	orderModelInfoStr  = "orderInfo"
	settlementModelStr = "settlement"
	eventModelStr      = "event"
)

// ModelType : model type
type ModelType int

const (
	UnKnownModel ModelType = iota
	AccountModel
	IssueModel
	ScheduleModel
	QuotaModel
	OrderInfoModel
	SettlementModel
	EventModel
)

// String : Stringer interface
func (t ModelType) String() string {
	switch t {
	case AccountModel:
		return accountModelStr
	case IssueModel:
		return issueModelStr
	case OrderInfoModel:
		return orderModelInfoStr
	case SettlementModel:
		return settlementModelStr
	case EventModel:
		return eventModelStr
	default:
		return unknownModelStr
	}
}

// MarshalJSON : Marshaler interface
func (t ModelType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

// UnmarshalJSON : Marshaler interface
func (t *ModelType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case accountModelStr:
		*t = AccountModel
	case issueModelStr:
		*t = IssueModel
	case orderModelInfoStr:
		*t = OrderInfoModel
	case settlementModelStr:
		*t = SettlementModel
	case eventModelStr:
		*t = EventModel
	default:
		*t = UnKnownModel
	}
	return nil
}
