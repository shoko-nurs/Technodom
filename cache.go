package main

type Cache interface {
	Add(key, value string)
	Get(key string) (value string, ok bool)
	Len() int
}

type Container struct {
	MapObj map[string]string
	Max    int
}

func (b *Container) MakeBox() {
	b.MapObj = make(map[string]string, b.Max)
}

func (b *Container) Add(key, value string) {

	// Add if does not exist and does not exceed the limit
	if _, ok := b.MapObj[key]; !ok {

		if b.Len() < b.Max {
			b.MapObj[key] = value
		}

	} else {
		// update if necessary
		b.MapObj[key] = value
	}

}

func (b *Container) Get(key string) (value string, ok bool) {
	if value, ok = b.MapObj[key]; ok {
		return value, ok
	}

	return "", false
}

func (b *Container) Len() int {

	return len(b.MapObj)
}

func PerformGet(cache Cache, key string) (string, int) {
	if val, ok := cache.Get(key); ok {

		if val == key {
			return key, 200
		} else {
			return val, 301
		}

	} else {
		return "", -1
	}
}

func PerformAdd(cache Cache, key, value string) {
	cache.Add(key, value)
}
