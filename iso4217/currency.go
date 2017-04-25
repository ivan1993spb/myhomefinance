package iso4217

import (
	"database/sql/driver"
	"errors"
)

type Currency byte

func (c Currency) String() string {
	if code, ok := codes[c]; ok {
		return code
	}
	return defaultCode
}

func (c Currency) Value() (driver.Value, error) {
	return c.GetCode(), nil
}

func (c *Currency) Scan(src interface{}) error {
	var err error
	*c, err = Parse(string(src))
	return err
}

func (c Currency) GetName() string {
	if name, ok := names[c]; ok {
		return name
	}
	return defaultName
}

func (c Currency) GetCode() string {
	if name, ok := codes[c]; ok {
		return name
	}
	return defaultCode
}

func Parse(code string) (Currency, error) {
	for currency, strCode := range codes {
		if strCode == code {
			return currency, nil
		}
	}
	return defaultCurrency, errors.New("cannot parse currency code")
}
