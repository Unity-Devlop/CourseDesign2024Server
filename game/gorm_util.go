package game

import (
	"database/sql/driver"
	"encoding/json"
)

type Uint32Array []uint32

func (c Uint32Array) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *Uint32Array) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

type ItemArray []Item

func (c Item) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *Item) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

func (c ItemArray) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *ItemArray) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

func (vec *Vector3) Scan(value interface{}) error {
	b, _ := value.([]byte)
	return json.Unmarshal(b, &vec)
}

func (vec Vector3) Value() (driver.Value, error) {
	return json.Marshal(vec)
}
