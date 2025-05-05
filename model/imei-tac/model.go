package imei_tac_model

import (
	"errors"
	"fmt"

	imei_tac_id_model "flida.dev/imei/model/imei-tac-id"
	"flida.dev/imei/model/imei-tac-rbi"
)

// TAC — Type Allocation Code.
//
// https://en.wikipedia.org/wiki/Type_Allocation_Code
type TAC struct {
	id  imei_tac_id_model.ID
	rbi imei_tac_rbi_model.RBI
}

func (x TAC) GetID() imei_tac_id_model.ID {
	return x.id
}

func (x TAC) GetRBI() imei_tac_rbi_model.RBI {
	return x.rbi
}

func (x TAC) IsValid() bool {
	return x.Validate() == nil
}

func (x TAC) Validate() error {
	if err := x.id.Validate(); err != nil {
		return fmt.Errorf("x.id.Validate: %w", err)
	}

	if err := x.rbi.Validate(); err != nil {
		return fmt.Errorf("x.rbi.Validate: %w", err)
	}

	return nil
}

func ParseTAC(s string) (TAC, error) {
	if len(s) != 8 { //nolint:mnd // OK.
		return TAC{}, errors.New("invalid TAC")
	}

	id, err := imei_tac_id_model.NewID(s[2:])
	if err != nil {
		return TAC{}, fmt.Errorf("imei_tac_id_model.NewID: %w", err)
	}

	rbi, err := imei_tac_rbi_model.NewRBI(s[:2])
	if err != nil {
		return TAC{}, fmt.Errorf("imei_tac_rbi_model.NewRBI: %w", err)
	}

	x := TAC{
		id:  id,
		rbi: rbi,
	}

	if err = x.Validate(); err != nil {
		return TAC{}, err
	}

	return x, nil
}
