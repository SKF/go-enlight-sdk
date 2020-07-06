package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/pkg/errors"
)

const (
	MetaDataMaxKeyLength    = 128
	MetaDataMaxValueLength  = 256
	MetaDataMaxKeys         = 50
	MetaDataAllowedKeyChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+-=._:/@ "
)

func isValidMetaDataKeyChar(char rune) bool {
	for _, allowedChar := range MetaDataAllowedKeyChars {
		if char == allowedChar {
			return true
		}
	}
	return false
}

type NodeMetaData map[string]*string

func (n1 NodeMetaData) Equal(n2 NodeMetaData) bool {
	return reflect.DeepEqual(n1, n2)
}

func (n NodeMetaData) Value() (driver.Value, error) {
	return json.Marshal(n)
}

func (n *NodeMetaData) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, &n)
}

func (n *NodeMetaData) Validate() error {
	if n == nil {
		return nil
	}
	if len(*n) > MetaDataMaxKeys {
		return fmt.Errorf("maximum number of metadata keys is %d, got %d", MetaDataMaxKeys, len(*n))
	}
	for k, v := range *n {
		if len(k) > MetaDataMaxKeyLength {
			return fmt.Errorf("metadata key \"%s\" violates maximum key length %d", k, MetaDataMaxKeyLength)
		}
		if v != nil && len(*v) > MetaDataMaxValueLength {
			return fmt.Errorf("metadata value \"%s\" violates maximum value length %d", k, MetaDataMaxValueLength)
		}
		for _, c := range k {
			if !isValidMetaDataKeyChar(c) {
				return fmt.Errorf("metadata key %s contains invalid character %#U", k, c)
			}
		}
	}

	return nil
}
