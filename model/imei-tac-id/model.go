package imei_tac_id_model

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

type ID string

func (x ID) Validate() error {
	if len(x) != 6 || !strings.ContainsFunc(string(x), unicode.IsDigit) {
		return errors.New("invalid ID")
	}

	return nil
}

func NewID(s string) (ID, error) {
	x := ID(s)

	if err := x.Validate(); err != nil {
		return "", fmt.Errorf("x.Validate: %w", err)
	}

	return x, nil
}
