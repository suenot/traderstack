package grpcsrv

import (
	"github.com/liderman/traderstack/internal/domain"
	infov1 "github.com/liderman/traderstack/internal/grpcsrv/gen/go/liderman/traderstack/info/v1"
)

type InfoToProtobufMapper struct {
}

func (t *InfoToProtobufMapper) MapInstruments(in []*domain.Share) []*infov1.Instrument {
	var ret []*infov1.Instrument
	for _, v := range in {
		ret = append(ret, t.MapInstrument(v))
	}

	return ret
}

func (t *InfoToProtobufMapper) MapInstrument(in *domain.Share) *infov1.Instrument {
	if in == nil {
		return nil
	}
	return &infov1.Instrument{
		Figi:     in.Figi,
		Ticker:   in.Ticker,
		Isin:     in.Isin,
		Lot:      in.Lot,
		Currency: in.Currency,
		Name:     in.Name,
		Exchange: in.Exchange,
	}
}

func (t *InfoToProtobufMapper) MapMoneys(in []*domain.MoneyValue) []*infov1.Money {
	var ret []*infov1.Money
	for _, v := range in {
		ret = append(ret, t.MapMoney(v))
	}
	return ret
}

func (t *InfoToProtobufMapper) MapMoney(in *domain.MoneyValue) *infov1.Money {
	return &infov1.Money{
		Currency: in.Currency,
		Value:    in.Value.String(),
	}
}

func (t *InfoToProtobufMapper) MapAccounts(in []*domain.Account) []*infov1.Account {
	var ret []*infov1.Account
	for _, v := range in {
		ret = append(ret, &infov1.Account{
			Id:   v.Id,
			Name: v.Name,
		})
	}

	return ret
}
