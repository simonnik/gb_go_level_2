// Протестируйте производительность операций чтения и записи на множестве действительных чисел,
// безопасность которого обеспечивается sync.Mutex и sync.RWMutex для разных вариантов использования:
// 10% запись, 90% чтение; 50% запись, 50% чтение; 90% запись, 10% чтение
package main

import (
	"sync"
	"testing"
)

type Set struct {
	sync.Mutex
	mm map[interface{}]interface{}
}

func NewSet() *Set {
	return &Set{
		mm: make(map[interface{}]interface{}),
	}
}

func (s *Set) Add(i int) {
	s.Lock()
	s.mm[i] = struct{}{}
	s.Unlock()
}

func (s *Set) Has(i int) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.mm[i]
	return ok
}

type RWSet struct {
	sync.RWMutex
	mm map[interface{}]interface{}
}

func NewRWSet() *RWSet {
	return &RWSet{
		mm: make(map[interface{}]interface{}),
	}
}

func (s *RWSet) Add(i int) {
	s.Lock()
	s.mm[i] = struct{}{}
	s.Unlock()
}

func (s *RWSet) Has(i int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.mm[i]
	return ok
}

const (
	num     int = 1000
	num10pc     = (int)(float32(num) * 0.1)
	num90pc     = (int)(float32(num) * 0.9)
	num50pc     = (int)(float32(num) * 0.5)
)

func BenchmarkMutex_10_90(b *testing.B) {
	var (
		set = NewSet()
		i   = 1
	)

	b.Run("Mutex 10% write 90% read", func(b *testing.B) {
		b.SetParallelism(num)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if i <= num10pc {
					set.Add(1)
				} else {
					set.Has(1)
				}
				i++
			}
		})
	})
}

func BenchmarkRWMutex_10_90(b *testing.B) {
	var (
		set = NewRWSet()
		i   = 1
	)

	b.Run("Mutex 10% write 90% read", func(b *testing.B) {
		b.SetParallelism(num)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if i <= num10pc {
					set.Add(1)
				} else {
					set.Has(1)
				}
				i++
			}
		})
	})
}

func BenchmarkMutex_50_50(b *testing.B) {
	var (
		set = NewSet()
		i   = 1
	)

	b.Run("Mutex 50% write 50% read", func(b *testing.B) {
		b.SetParallelism(num)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if i <= num50pc {
					set.Add(1)
				} else {
					set.Has(1)
				}
				i++
			}
		})
	})
}

func BenchmarkRWMutex_50_50(b *testing.B) {
	var (
		set = NewRWSet()
		i   = 1
	)

	b.Run("Mutex 50% write 50% read", func(b *testing.B) {
		b.SetParallelism(num)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if i <= num50pc {
					set.Add(1)
				} else {
					set.Has(1)
				}
				i++
			}
		})
	})
}

func BenchmarkMutex_90_10(b *testing.B) {
	var (
		set = NewSet()
		i   = 1
	)

	b.Run("Mutex 90% write 10% read", func(b *testing.B) {
		b.SetParallelism(num)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if i <= num90pc {
					set.Add(1)
				} else {
					set.Has(1)
				}
				i++
			}
		})
	})
}

func BenchmarkRWMutex_90_10(b *testing.B) {
	var (
		set = NewRWSet()
		i   = 1
	)

	b.Run("Mutex 90% write 10% read", func(b *testing.B) {
		b.SetParallelism(num)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if i <= num90pc {
					set.Add(1)
				} else {
					set.Has(1)
				}
				i++
			}
		})
	})
}
