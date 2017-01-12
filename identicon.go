package identicon

// A Data provides to construct identicon.
type Data struct {
	id string
}

// NewData returns a Data object.
func NewData(id string) Data {
	return Data{
		id: id,
	}
}
