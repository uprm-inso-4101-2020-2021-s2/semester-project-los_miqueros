package Themis

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/keeper"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the vote
	for _, elem := range genState.VoteList {
		k.SetVote(ctx, *elem)
	}

	// Set vote count
	k.SetVoteCount(ctx, uint64(len(genState.VoteList)))

	// Set all the poll
	for _, elem := range genState.PollList {
		k.SetPoll(ctx, *elem)
	}

	// Set poll count
	k.SetPollCount(ctx, uint64(len(genState.PollList)))

	// Set all the group
	for _, elem := range genState.GroupList {
		k.SetGroup(ctx, *elem)
	}

	// Set group count
	k.SetGroupCount(ctx, uint64(len(genState.GroupList)))

	// Set all the account
	for _, elem := range genState.AccountList {
		k.SetAccount(ctx, *elem)
	}

	// Set account count
	k.SetAccountCount(ctx, uint64(len(genState.AccountList)))

}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all vote
	voteList := k.GetAllVote(ctx)
	for _, elem := range voteList {
		elem := elem
		genesis.VoteList = append(genesis.VoteList, &elem)
	}

	// Get all poll
	pollList := k.GetAllPoll(ctx)
	for _, elem := range pollList {
		elem := elem
		genesis.PollList = append(genesis.PollList, &elem)
	}

	// Get all group
	groupList := k.GetAllGroup(ctx)
	for _, elem := range groupList {
		elem := elem
		genesis.GroupList = append(genesis.GroupList, &elem)
	}

	// Get all account
	accountList := k.GetAllAccount(ctx)
	for _, elem := range accountList {
		elem := elem
		genesis.AccountList = append(genesis.AccountList, &elem)
	}

	return genesis
}
