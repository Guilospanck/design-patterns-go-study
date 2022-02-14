package main

func main() {
	// first strategy
	lfuStrategy := NewLFUStrategy()

	// context
	cache := NewCache(lfuStrategy)

	// add some content
	cache.Add("key1", "value1")
	cache.Add("key2", "value2")
	cache.Add("key3", "value3")

	cache.PrintCache()

	// set another strategy
	cache.SetEvictStrategy(NewFIFOStrategy())

	cache.Add("key4", "value4")
	cache.Add("key5", "value5")

	cache.PrintCache()

	cache.SetEvictStrategy(NewLRUStrategy())

	cache.Add("key6", "value6")
	cache.Add("key7", "value7")

	cache.PrintCache()

}
