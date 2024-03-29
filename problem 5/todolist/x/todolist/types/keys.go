package types

const (
	// ModuleName defines the module name
	ModuleName = "todolist"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_todolist"

	// ItemKey defines the prefix to store each item
	ItemKey = "Item/value/"

	// ItemCountKey defines the key to store the count of items
	ItemCountKey = "Item/count/"
)

var (
	ParamsKey = []byte("p_todolist")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
