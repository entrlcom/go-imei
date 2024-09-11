package imei_tac_rbi

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidRBI = errors.New("invalid RBI")

// RBI — Reporting Body Identifier.
//
// https://en.wikipedia.org/wiki/Reporting_Body_Identifier
type RBI string

func (x RBI) IsTestIMEI() bool {
	switch x {
	case "00", "02", "03", "04", "05", "06", "07", "08", "09": //nolint:goconst // OK.
		return true
	default:
		return false
	}
}

func (x RBI) IsValid() bool {
	return x.Validate() == nil
}

func (x RBI) String() string {
	return string(x)
}

//nolint:cyclop // OK.
func (x RBI) Validate() error {
	if len(x) != 2 || !strings.ContainsFunc(string(x), unicode.IsDigit) {
		return ErrInvalidRBI
	}

	switch x {
	case "00": // Test IMEI. Nations with 2-digit CCs.
		return nil
	case "01": // PTCRB. United States.
		return nil
	case "02", "03", "04", "05", "06", "07", "08", "09": // Test IMEI. Nations with 3-digit CCs.
		return nil
	case "10": // DECT devices.
		return nil
	case "30": // Iridium. United States (satellite phones).
		return nil
	case "33": // DGPT. France.
		return nil
	case "35", "44", "98": // BABT. United Kingdom.
		return nil
	case "45": // NTA. Germany.
		return nil
	case "49": // BZT / BAPT. Germany.
		return nil
	case "50": // BZT ETS. Germany,.
		return nil
	case "51": // Cetecom ICT. Germany.
		return nil
	case "52": // Cetecom. Germany.
		return nil
	case "53": // TÜV. Germany.
		return nil
	case "54": // Phoenix Test Lab. Germany.
		return nil
	case "86": // TAF. China.
		return nil
	case "91": // MSAI. India.
		return nil
	case "99": // GHA. For multi RAT 3GPP2/3GPP.
		return nil
	default:
		return ErrInvalidRBI
	}
}

func NewRBI(s string) (RBI, error) {
	x := RBI(s)

	if err := x.Validate(); err != nil {
		return "", err
	}

	return x, nil
}
