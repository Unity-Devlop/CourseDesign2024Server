package game

import (
	"database/sql/driver"
	"encoding/json"
)

func (vec *Vector3) Scan(value interface{}) error {
	b, _ := value.([]byte)
	return json.Unmarshal(b, &vec)
}

func (vec Vector3) Value() (driver.Value, error) {
	return json.Marshal(vec)
}
