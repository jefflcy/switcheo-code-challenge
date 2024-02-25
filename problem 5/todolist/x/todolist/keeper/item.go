package keeper

import (
	"encoding/binary"

	"todolist/x/todolist/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Get item by ID
func (k Keeper) GetItem(ctx sdk.Context, id uint64) (val types.Item, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ItemKey))
	b := store.Get(GetItemIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// Set item with item to update the item in store
func (k Keeper) SetItem(ctx sdk.Context, post types.Item) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ItemKey))
	b := k.cdc.MustMarshal(&post)
	store.Set(GetItemIDBytes(post.Id), b)
}

// Delete item by ID
func (k Keeper) RemoveItem(ctx sdk.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ItemKey))
	store.Delete(GetItemIDBytes(id))
}

// Append items to the store
func (k Keeper) AppendItem(ctx sdk.Context, item types.Item) uint64 {
	count := k.GetItemCount(ctx)
	item.Id = count
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ItemKey))
	appendedValue := k.cdc.MustMarshal(&item)
	store.Set(GetItemIDBytes(item.Id), appendedValue)
	k.SetItemCount(ctx, count+1)
	return count
}

// Gettting the item count
func (k Keeper) GetItemCount(ctx sdk.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.ItemCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

// Converts item ID to bytes
func GetItemIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// Update the item count
func (k Keeper) SetItemCount(ctx sdk.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.ItemCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}
