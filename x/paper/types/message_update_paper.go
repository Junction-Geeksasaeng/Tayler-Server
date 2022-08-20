package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdatePaper = "update_paper"

var _ sdk.Msg = &MsgUpdatePaper{}

func NewMsgUpdatePaper(creator string, id string, newOwner string, newPrice string) *MsgUpdatePaper {
	return &MsgUpdatePaper{
		Creator:  creator,
		Id:       id,
		NewOwner: newOwner,
		NewPrice: newPrice,
	}
}

func (msg *MsgUpdatePaper) Route() string {
	return RouterKey
}

func (msg *MsgUpdatePaper) Type() string {
	return TypeMsgUpdatePaper
}

func (msg *MsgUpdatePaper) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdatePaper) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdatePaper) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
