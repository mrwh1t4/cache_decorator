package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/mrwh1t4/cache_decorator"
)

var ErrNotFound = errors.New("Not Found in Cache")

type mapCache[K comparable, V any] struct {
	dict map[K]V
}

func (d mapCache[K, V]) Set(ctx context.Context, key K, value V) {
	d.dict[key] = value
}

func (d mapCache[K, V]) Get(ctx context.Context, key K) (V, error) {
	if value, ok := d.dict[key]; ok {
		fmt.Println(key, "found in mapCache")
		return value, nil
	}
	fmt.Println(key, "not found in mapCache")
	return *new(V), ErrNotFound
}

type TestClass struct {
}

func (TestClass) Get(ctx context.Context, key uint32) (string, error) {
	fmt.Println(key, "found in object")
	return "result_1", nil
}

func main() {
	d := mapCache[uint32, string]{map[uint32]string{}}

	r := TestClass{}
	newObj := cache_decorator.NewDecorator[uint32, string](d)
	newObj.Decorate(r)
	ctx := context.Background()
	d.Set(ctx, 0, "result_0")
	fmt.Println("Try to Get key 0")
	newObj.Get(ctx, 0)
	fmt.Println("\nTry to Get key 1")
	newObj.Get(ctx, 1)
	fmt.Println("\nTry to Get key 1 again ")
	newObj.Get(ctx, 1)
}
