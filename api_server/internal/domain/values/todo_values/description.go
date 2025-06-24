package todo_values

type Description struct {
	value string
}

func NewDescription(description_value string) (Description, error) {
	return Description{
		value: description_value,
	}, nil
}

func (d *Description) GetValue() string {
	return d.value
}
