package hash_map

import (
	"fmt"
	"hash/fnv"
	"reflect"
)

const loadFactor = 0.75
const capacity = 2

type KeyValue[K any, V any] struct {
	Key   K
	Value V
}

type HashMap[K any, V any] struct {
	items    [][]KeyValue[K, V]
	capacity int
	size     int
}

func (hashMap *HashMap[K, V]) hashCode(key K) int {
	bytes := []byte(fmt.Sprint(key))
	h := fnv.New32a()
	h.Write(bytes)
	return int(h.Sum32())
}

func (hashMap *HashMap[K, V]) indexOf(key K) int {
	return hashMap.hashCode(key) % hashMap.capacity
}

func (hashMap *HashMap[K, V]) Put(key K, value V) {
	if float32(hashMap.size)/float32(hashMap.capacity) > loadFactor {
		hashMap.resize()
	}

	index := hashMap.indexOf(key)
	values := hashMap.items[index]
	if values != nil {
		for i := 0; i < len(values); i++ {
			if reflect.DeepEqual(values[i].Key, key) {
				values[i] = KeyValue[K, V]{key, value}
				return
			}
		}
	}
	hashMap.items[index] = append(values, KeyValue[K, V]{key, value})
	hashMap.size++
}

func (hashMap *HashMap[K, V]) resize() {
	hashMap.capacity *= 2
	items := make([][]KeyValue[K, V], hashMap.capacity)

	for i := 0; i < len(hashMap.items); i++ {
		for j := 0; j < len(hashMap.items[i]); j++ {
			keyValue := hashMap.items[i][j]
			values := items[hashMap.indexOf(keyValue.Key)]
			if values != nil {
				for k := 0; k < len(values); k++ {
					values[k] = keyValue
				}
			}
			items[hashMap.indexOf(keyValue.Key)] = append(values, keyValue)
		}
	}

	hashMap.items = items
}

func (hashMap *HashMap[K, V]) Get(key K) (*V, error) {
	items := hashMap.items[hashMap.indexOf(key)]

	if items == nil {
		return nil, fmt.Errorf("key not found")
	}

	for i := 0; i < len(items); i++ {
		if reflect.DeepEqual(items[i].Key, key) {
			return &items[i].Value, nil
		}
	}
	return nil, fmt.Errorf("key not found")
}

func (hashMap *HashMap[K, V]) Size() int {
	return hashMap.size
}

func New[K any, V any]() *HashMap[K, V] {
	return &HashMap[K, V]{items: make([][]KeyValue[K, V], capacity), capacity: capacity, size: 0}
}
