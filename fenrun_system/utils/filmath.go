package utils

import (
	"encoding"
	"fmt"
	big2 "github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/lotus/build"
	"math/big"
	"strconv"
	"strings"
)

const (
	AttoFIL = 18
	NanoFIL = 9
)

func NanoOrAttoToFIL(fil string, filtype int) (res float64) {
	//大于18or9位
	if len(fil) > filtype {
		str := fil[0:len(fil)-filtype] + "." + fil[len(fil)-filtype:]
		res, _ = strconv.ParseFloat(str, 64)
		return res
	}
	//小于18or9位
	str := "0."
	for i := 0; i < filtype-len(fil); i++ {
		str += "0"
	}
	str = str + fil
	res, _ = strconv.ParseFloat(str, 64)
	return res
}

func NanoOrAttoToFILstr(fil string, filtype int) (res string) {
	//大于18or9位
	if len(fil) > filtype {
		str := fil[0:len(fil)-filtype] + "." + fil[len(fil)-filtype:]
		res = str
		return res
	}
	//小于18or9位
	str := "0."
	for i := 0; i < filtype-len(fil); i++ {
		str += "0"
	}
	str = str + fil
	return res
}

type BigInt = big2.Int

type FIL BigInt

func (f FIL) String() string {
	return f.Unitless()
}

func (f FIL) Unitless() string {
	r := new(big.Rat).SetFrac(f.Int, big.NewInt(int64(build.FilecoinPrecision)))
	if r.Sign() == 0 {
		return "0"
	}
	return strings.TrimRight(strings.TrimRight(r.FloatString(18), "0"), ".")
}

var unitPrefixes = []string{"a", "f", "p", "n", "μ", "m"}

func (f FIL) Short() string {
	n := BigInt(f).Abs()

	dn := uint64(1)
	var prefix string
	for _, p := range unitPrefixes {
		if n.LessThan(NewInt(dn * 1000)) {
			prefix = p
			break
		}
		dn *= 1000
	}

	r := new(big.Rat).SetFrac(f.Int, big.NewInt(int64(dn)))
	if r.Sign() == 0 {
		return "0"
	}

	return strings.TrimRight(strings.TrimRight(r.FloatString(3), "0"), ".") + " " + prefix + "FIL"
}

func (f FIL) Format(s fmt.State, ch rune) {
	switch ch {
	case 's', 'v':
		fmt.Fprint(s, f.String())
	default:
		f.Int.Format(s, ch)
	}
}

func (f FIL) MarshalText() (text []byte, err error) {
	return []byte(f.String()), nil
}

func (f FIL) UnmarshalText(text []byte) error {
	p, err := ParseFIL(string(text))
	if err != nil {
		return err
	}

	f.Int.Set(p.Int)
	return nil
}

func ParseFIL(s string) (FIL, error) {
	suffix := strings.TrimLeft(s, "-.1234567890")
	s = s[:len(s)-len(suffix)]
	var attofil bool
	if suffix != "" {
		norm := strings.ToLower(strings.TrimSpace(suffix))
		switch norm {
		case "", "fil":
		case "attofil", "afil":
			attofil = true
		default:
			return FIL{}, fmt.Errorf("unrecognized suffix: %q", suffix)
		}
	}

	if len(s) > 50 {
		return FIL{}, fmt.Errorf("string length too large: %d", len(s))
	}

	r, ok := new(big.Rat).SetString(s)
	if !ok {
		return FIL{}, fmt.Errorf("failed to parse %q as a decimal number", s)
	}

	if !attofil {
		r = r.Mul(r, big.NewRat(int64(build.FilecoinPrecision), 1))
	}

	if !r.IsInt() {
		var pref string
		if attofil {
			pref = "atto"
		}
		return FIL{}, fmt.Errorf("invalid %sFIL value: %q", pref, s)
	}

	return FIL{r.Num()}, nil
}

func MustParseFIL(s string) FIL {
	n, err := ParseFIL(s)
	if err != nil {
		panic(err)
	}

	return n
}

var _ encoding.TextMarshaler = (*FIL)(nil)
var _ encoding.TextUnmarshaler = (*FIL)(nil)

func NewInt(i uint64) BigInt {
	return BigInt{Int: big.NewInt(0).SetUint64(i)}
}
