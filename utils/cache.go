package utils

func NewCache[K comparable, V any]() (func(K, V) V, func(K) (V, bool)) {
	c := map[K]V{}
	cache := func(k K, v V) V {
		c[k] = v
		return v
	}
	getCache := func(k K) (V, bool) {
		v, ok := c[k]
		return v, ok
	}
	return cache, getCache
}
