package db

type Oracle struct {
	conn Connection
}

func (db *Oracle) ValidateString(string string) error {
}
