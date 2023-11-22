package mtproto

type Option[T any] struct {
	IsSet bool
	Value T
}

func (o *Option[T]) Set(val T) {
	o.Value = val
	o.IsSet = true
}

func (o *Option[T]) Clear() {
	var blank T
	o.Value = blank
	o.IsSet = false
}

func Some[T any](val T) Option[T] {
	return Option[T]{IsSet: true, Value: val}
}

func None[T any]() Option[T] {
	return Option[T]{IsSet: false}
}
