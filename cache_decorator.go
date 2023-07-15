package cache_decorator

import (
	"context"
)

type DecoratorCache[K comparable, V any] interface {
	Get(context.Context, K) (V, error)
	Set(context.Context, K, V)
}

type decoratedClass[K comparable, V any] interface {
	Get(context.Context, K) (V, error)
}

type decorator[K comparable, V any] struct {
	decObj decoratedClass[K, V]
	cache  DecoratorCache[K, V]
}

func (d *decorator[K, V]) Get(ctx context.Context, key K) (V, error) {
	value, err := d.cache.Get(ctx, key)
	if err == nil {
		return value, nil
	}

	v, err := d.decObj.Get(ctx, key)
	if err != nil {
		return *new(V), err
	}

	d.cache.Set(ctx, key, v)
	return v, nil
}

func NewDecorator[K comparable, V any](cache DecoratorCache[K, V]) *decorator[K, V] {
	a := decorator[K, V]{
		cache: cache,
	}
	return &a
}

func (d *decorator[K, V]) Decorate(g decoratedClass[K, V]) {
	d.decObj = g
}
