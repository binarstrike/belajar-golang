package helper

type IF[Ttype any] bool

func (c IF[T]) If(a, b T) T {
	if c {
		return a
	}
	return b
}
