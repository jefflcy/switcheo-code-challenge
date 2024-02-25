package keeper

import (
	"context"

	"todolist/x/todolist/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateItem(goCtx context.Context, msg *types.MsgUpdateItem) (*types.MsgUpdateItemResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUpdateItemResponse{}, nil
}
