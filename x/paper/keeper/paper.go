package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"paper/x/paper/types"

	"fmt"
	"strconv"
)

// func (k Keeper) AppendPost() uint64 {
//   count := k.GetPostCount()
//   store.Set()
//   k.SetPostCount()
//   return count
// }

func (k Keeper) GetPaperCount(ctx sdk.Context) uint64 {
	// Get the store using storeKey (which is "blog") and PostCountKey (which is "Post-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PaperCountKey))

	// Convert the PostCountKey to bytes
	byteKey := []byte(types.PaperCountKey)

	// Get the value of the count
	bz := store.Get(byteKey)

	// Return zero if the count value is not found (for example, it's the first post)
	if bz == nil {
		return 0
	}

	// Convert the count into a uint64
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetPaperCount(ctx sdk.Context, count uint64) {
	// Get the store using storeKey (which is "blog") and PostCountKey (which is "Post-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PaperCountKey))

	// Convert the PostCountKey to bytes
	byteKey := []byte(types.PaperCountKey)

	// Convert count from uint64 to string and get bytes
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)

	// Set the value of Post-count- to count
	store.Set(byteKey, bz)
}

func (k Keeper) AppendPaper(ctx sdk.Context, paper types.Paper) uint64 {
	// Get the current number of posts in the store
	count := k.GetPaperCount(ctx)

	// Assign an ID to the post based on the number of posts in the store
	paper.Id = count

	// Get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PaperKey))

	// Convert the post ID into bytes
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, paper.Id)

	// Marshal the post into bytes
	appendedValue := k.cdc.MustMarshal(&paper)

	// Insert the post bytes using post ID as a key
	store.Set(byteKey, appendedValue)

	// Update the post count
	k.SetPaperCount(ctx, count+1)
	return count
}

func (k Keeper) GetPaper(ctx sdk.Context, id uint64) (val types.Paper, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PaperKey))

	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)

	b := store.Get(bz)
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetPaper(ctx sdk.Context, paper types.Paper) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PaperKey))

	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, paper.Id)

	newPaper := k.cdc.MustMarshal(&paper)

	store.Set(byteKey, newPaper)
}

func (k Keeper) ChangeOwnerAndPrice(ctx sdk.Context, id string, newOwner string, newPrice string) bool {
	targetId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	paper, found := k.GetPaper(ctx, targetId)
	if !found {
		return false
	}

	paper.Owner = newOwner
	paper.Price = newPrice

	k.SetPaper(ctx, paper)
	return true
}
