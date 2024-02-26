package keeper

import (
	"context"
	"sort"

	"todolist/x/todolist/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListItem(ctx context.Context, req *types.QueryListItemRequest) (*types.QueryListItemResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ItemKey))

	var items []types.Item
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var item types.Item
		if err := k.cdc.Unmarshal(iterator.Value(), &item); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		items = append(items, item)
	}

	// Check if we have at least one item; if not, return an error or an empty response
	if len(items) == 0 {
		return nil, status.Error(codes.NotFound, "no items found")
	}

	// Sort the items by priority, highest first
	sort.Slice(items, func(i, j int) bool {
		return items[i].Priority > items[j].Priority
	})

	// Select the most important one - the first item after sorting
	mostImportantItem := items[0]

	// Return only the most important item
	return &types.QueryListItemResponse{Item: &mostImportantItem}, nil
}
