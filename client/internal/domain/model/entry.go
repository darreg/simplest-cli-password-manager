package model

type Entry struct {
	ID       string
	TypeID   string
	Name     string
	Metadata string
	Data     []byte
}

func NewEntry(id, typeID, name, metadata string, data []byte) *Entry {
	return &Entry{
		ID:       id,
		TypeID:   typeID,
		Name:     name,
		Metadata: metadata,
		Data:     data,
	}
}
