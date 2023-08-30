package percent_segments

type Cache struct {
	elements map[string]Segment
}

func NewCache() *Cache {
	return &Cache{
		elements: make(map[string]Segment),
	}
}

func (c *Cache) Get() map[string]Segment {
	return c.elements
}

func (c *Cache) Update(segments map[string]Segment) {
	c.elements = segments
}

func (c *Cache) Delete(key string) {
	delete(c.elements, key)
}
