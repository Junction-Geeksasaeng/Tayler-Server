package keeper

import (
	"context"

	"paper/x/paper/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreatePaper(goCtx context.Context, msg *types.MsgCreatePaper) (*types.MsgCreatePaperResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Create variable of type Post
	var paper = types.Paper{
		Creator:   msg.Creator,
		Host:      msg.Host,
		PaperName: msg.PaperName,
		Owner:     msg.Owner,
		Price:     msg.Price,
	}

	// Add a post to the store and get back the ID
	id := k.AppendPaper(ctx, paper)

	// Return the ID of the post
	return &types.MsgCreatePaperResponse{Id: id}, nil
}
