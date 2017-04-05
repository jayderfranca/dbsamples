package db

type Oracle struct {
	conn *defs.Connection
}

func (db *Oracle) ValidateString(string string) error {
}
