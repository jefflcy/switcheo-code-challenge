package keeper

import (
	"context"

	"todolist/x/todolist/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DeleteItem(goCtx context.Context, msg *types.MsgDeleteItem) (*types.MsgDeleteItemResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgDeleteItemResponse{}, nil
}
