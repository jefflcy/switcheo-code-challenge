package keeper

import (
	"context"

	"todolist/x/todolist/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListAllItems(ctx context.Context, req *types.QueryListAllItemsRequest) (*types.QueryListAllItemsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ItemKey))

	var items []types.Item
	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var post types.Item
		if err := k.cdc.Unmarshal(value, &post); err != nil {
			return err
		}

		items = append(items, post)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListAllItemsResponse{Item: items, Pagination: pageRes}, nil
}
