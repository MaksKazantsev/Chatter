package converter

type Converter interface {
}

type converter struct {
}

func NewConverter() Converter {
	return &converter{}
}
