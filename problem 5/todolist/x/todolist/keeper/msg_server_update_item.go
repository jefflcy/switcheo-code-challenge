package keeper

import (
	"context"
	"fmt"

	"todolist/x/todolist/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateItem(goCtx context.Context, msg *types.MsgUpdateItem) (*types.MsgUpdateItemResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var item = types.Item{
		Creator: msg.Creator,
		Id:      msg.Id,
		Desc:    msg.Desc,
	}
	val, found := k.GetItem(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	k.SetItem(ctx, item)
	return &types.MsgUpdateItemResponse{}, nil
}
