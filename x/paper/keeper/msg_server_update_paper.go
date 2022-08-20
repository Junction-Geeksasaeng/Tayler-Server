package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"paper/x/paper/types"
)

func (k msgServer) UpdatePaper(goCtx context.Context, msg *types.MsgUpdatePaper) (*types.MsgUpdatePaperResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUpdatePaperResponse{}, nil
}
