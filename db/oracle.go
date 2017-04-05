package db

import (
	"github.com/jayderfranca/dbsamples/defs"
)

type Oracle struct {
	conn *defs.Connection
}

func (db *Oracle) ValidateString(string string) error {
}
