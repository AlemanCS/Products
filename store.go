package main

//Caches interface
type Caches interface {
	GetProducts(id string) ([]string, error)
	ForceCache(id string, value string) error
	cacheExists(id string) (bool, error)
	getCache(id string) (string, error)
	updateCache(id string, products []string) error
}
