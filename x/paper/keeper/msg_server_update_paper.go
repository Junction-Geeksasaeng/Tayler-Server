package keeper

import (
	"context"

	"paper/x/paper/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdatePaper(goCtx context.Context, msg *types.MsgUpdatePaper) (*types.MsgUpdatePaperResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.ChangeOwner(ctx, msg.Id, msg.NewOwner, msg.NewPrice)

	return &types.MsgUpdatePaperResponse{IsSuccess: true}, nil
}
