package keeper

import (
	"context"

	"todolist/x/todolist/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowItem(goCtx context.Context, req *types.QueryShowItemRequest) (*types.QueryShowItemResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	item, found := k.GetItem(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}
	return &types.QueryShowItemResponse{Item: &item}, nil
}
