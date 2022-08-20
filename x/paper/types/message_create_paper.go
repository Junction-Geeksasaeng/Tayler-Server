package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreatePaper = "create_paper"

var _ sdk.Msg = &MsgCreatePaper{}

func NewMsgCreatePaper(creator string, host string, paperName string, owner string, price string) *MsgCreatePaper {
	return &MsgCreatePaper{
		Creator:   creator,
		Host:      host,
		PaperName: paperName,
		Owner:     owner,
		Price:     price,
	}
}

func (msg *MsgCreatePaper) Route() string {
	return RouterKey
}

func (msg *MsgCreatePaper) Type() string {
	return TypeMsgCreatePaper
}

func (msg *MsgCreatePaper) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreatePaper) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreatePaper) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
