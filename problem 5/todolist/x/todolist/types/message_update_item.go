package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateItem{}

func NewMsgUpdateItem(creator string, desc string, priority uint64, id uint64) *MsgUpdateItem {
	return &MsgUpdateItem{
		Creator:  creator,
		Desc:     desc,
		Priority: priority,
		Id:       id,
	}
}

func (msg *MsgUpdateItem) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
