package usecase

type SetEntryDTO struct {
	TypeID   string
	Name     string
	Metadata string
	Data     []byte
}
