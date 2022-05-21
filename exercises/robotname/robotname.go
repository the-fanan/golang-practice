package robotname

import (
	"errors"
	"math/rand"
	"strconv"
)

var gn = map[string]bool{}

const maxAvailableNames int = 26 * 26 * 10 * 10 * 10

type Robot struct {
    name string
}

func RandomChar() string {
	max := int('Z')
	min := int('A')
	value := rand.Intn(max-min) + min

	return string(byte(value))
}

func RandomNum() string {
	return strconv.Itoa(rand.Intn(9-0))
}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	if len(gn) < maxAvailableNames {
		code := ""
		code += RandomChar()
		code += RandomChar()
		code += RandomNum()
		code += RandomNum()
		code += RandomNum()
		if _, ok := gn[code]; ok {
			return r.Name()
		}

		gn[code] = true
		r.name = code
		return r.name, nil 
	}

	return "", errors.New("Name space has been exhausted")
}

func (r *Robot) Reset() {
	r.name = ""
}
