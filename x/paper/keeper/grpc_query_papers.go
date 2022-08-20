package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"paper/x/paper/types"
)

func (k Keeper) Papers(goCtx context.Context, req *types.QueryPapersRequest) (*types.QueryPapersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var papers []*types.Paper

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the key-value module store using the store key (in our case store key is "chain")
	store := ctx.KVStore(k.storeKey)

	// Get the part of the store that keeps posts (using post key, which is "Post-value-")
	paperStore := prefix.NewStore(store, []byte(types.PaperKey))

	// Paginate the posts store based on PageRequest
	pageRes, err := query.Paginate(paperStore, req.Pagination, func(key []byte, value []byte) error {
		var paper types.Paper
		if err := k.cdc.Unmarshal(value, &paper); err != nil {
			return err
		}

		papers = append(papers, &paper)

		return nil
	})

	// Throw an error if pagination failed
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Return a struct containing a list of posts and pagination info
	return &types.QueryPapersResponse{Paper: papers, Pagination: pageRes}, nil
}
