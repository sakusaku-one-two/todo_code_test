package values

type IValue[valueType comparable] interface {
	GetValue() valueType
}
