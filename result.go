package result

import (
	"fmt"
	// "runtime"
)

type Resolver[T any] interface {
	Ok() T
	IsOk() bool
	Err() error
	Expect(string) T
	Result() T
	Then(func(T) Resolver[T]) Resolver[T]
	Error() string
	And(Resolver[T]) Resolver[T]
	Or(Resolver[T]) Resolver[T]
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

func (r Result[T]) And(r2 Resolver[T]) Resolver[T] {
	if !r.IsOk() {
		return Error[T](r.err)
	}
	return r2
}

func (r Result[T]) Or(r2 Resolver[T]) Resolver[T] {
	if r.IsOk() {
		return r
	}
	return r2
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
	// runtime.Caller(1)
	if r.err != nil {
		panic(fmt.Sprintf("%s, %s", r.err.Error(), msg))
	}
	return r.t
}

func (r Result[T]) Result() T {
	if r.err != nil {
		r.Expect("error occured")
	}
	return r.t
}

func (r Result[T]) Then(f func(T) Resolver[T]) Resolver[T] {
	if !r.IsOk() {
		return Error[T](r.Err())
	}
	return f(r.t)
}
