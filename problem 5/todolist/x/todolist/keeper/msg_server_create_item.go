package keeper

import (
	"context"

	"todolist/x/todolist/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateItem(goCtx context.Context, msg *types.MsgCreateItem) (*types.MsgCreateItemResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Handling the message
	var item = types.Item{
		Creator:  msg.Creator,
		Desc:     msg.Desc,
		Priority: msg.Priority,
	}
	id := k.AppendItem(
		ctx,
		item,
	)
	return &types.MsgCreateItemResponse{
		Id: id,
	}, nil
}
