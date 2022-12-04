package result

import (
	"fmt"
)

type Resolver[T any] interface {
	Ok() T
	IsOk() bool
	Err() error
	Expect(string) T
	Result() T
	Then(func(T) (T, error)) Resolver[T]
	Error() string
}

type Result[T any] struct {
	err error
	t   T
}

func Ok[T any](t T) Result[T] {
	return Result[T]{err: nil, t: t}
}

func Error[T any](err error) Result[T] {
	return Result[T]{err: err}
}

func (r Result[T]) And(res Resolver[T]) Resolver[T] {
	if !r.IsOk() {
		return Error[T](r.err)
	}
	return res
}

// will return the first result that is ok.
func Or[D any](res ...Resolver[D]) Resolver[D] {
	panic("not implemented") // TODO: Implement
}

func (r Result[T]) Ok() T {
	return r.t
}

func (r Result[D]) IsOk() bool {
	return r.err == nil
}

func (r Result[D]) Err() error {
	return r.err
}

func (r Result[D]) Error() string {
	return r.err.Error()
}

func (r Result[D]) Expect(msg string) D {
	if r.err != nil {
		panic(fmt.Sprintf("result: %s, error:%s", msg, r.err.Error()))
	}
	return r.t
}

func (r Result[T]) Result() T {
	if r.err != nil {
		r.Expect("error occured")
	}
	return r.t
}

func (r Result[T]) Then(f func(T) (T, error)) Resolver[T] {
	if !r.IsOk() {
		return r
	}
	t, err := f(r.t)
	if err != nil {
		return Error[T](err)
	}
	return Ok(t)
}
