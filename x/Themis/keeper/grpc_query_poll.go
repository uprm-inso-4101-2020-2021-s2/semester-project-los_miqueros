package keeper

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PollAll(c context.Context, req *types.QueryAllPollRequest) (*types.QueryAllPollResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var polls []*types.Poll
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	pollStore := prefix.NewStore(store, types.KeyPrefix(types.PollKey))

	pageRes, err := query.Paginate(pollStore, req.Pagination, func(key []byte, value []byte) error {
		var poll types.Poll
		if err := k.cdc.UnmarshalBinaryBare(value, &poll); err != nil {
			return err
		}

		polls = append(polls, &poll)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPollResponse{Poll: polls, Pagination: pageRes}, nil
}

func (k Keeper) PollsWithName(c context.Context, req *types.QueryAllPollsWithNameRequest) (*types.QueryAllPollResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var polls []*types.Poll
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	pollStore := prefix.NewStore(store, types.KeyPrefix(types.PollNameKey))

	pageRes, err := types.PrefixPaginate(pollStore, types.KeyPrefix(req.Name), req.Pagination, func(key []byte, value []byte) error {
		var poll types.Poll
		if err := k.cdc.UnmarshalBinaryBare(value, &poll); err != nil {
			return err
		}

		polls = append(polls, &poll)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPollResponse{Poll: polls, Pagination: pageRes}, nil
}

func (k Keeper) PollsInGroup(c context.Context, req *types.QueryAllPollsInGroupRequest) (*types.QueryAllPollResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var polls []*types.Poll
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	pollStore := prefix.NewStore(store, types.KeyPrefix(types.PollGroupKey))

	pageRes, err := types.PrefixPaginate(pollStore, types.KeyPrefix(strconv.FormatUint(req.Group, 10)), req.Pagination, func(key []byte, value []byte) error {
		var poll types.Poll
		if err := k.cdc.UnmarshalBinaryBare(value, &poll); err != nil {
			return err
		}

		polls = append(polls, &poll)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPollResponse{Poll: polls, Pagination: pageRes}, nil
}

func (k Keeper) PollsInGroupWithName(c context.Context, req *types.QueryAllPollsInGroupWithNameRequest) (*types.QueryAllPollResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var polls []*types.Poll
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	pollStore := prefix.NewStore(store, types.KeyPrefix(types.PollGroupKey))

	pageRes, err := types.PrefixPaginate(pollStore, types.KeyPrefix(strconv.FormatUint(req.Group, 10)+"-"+req.Name), req.Pagination, func(key []byte, value []byte) error {
		var poll types.Poll
		if err := k.cdc.UnmarshalBinaryBare(value, &poll); err != nil {
			return err
		}

		polls = append(polls, &poll)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPollResponse{Poll: polls, Pagination: pageRes}, nil
}

func (k Keeper) Poll(c context.Context, req *types.QueryGetPollRequest) (*types.QueryGetPollResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var poll types.Poll
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasPoll(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetPollIDBytes(req.Id)), &poll)

	return &types.QueryGetPollResponse{Poll: &poll}, nil
}
