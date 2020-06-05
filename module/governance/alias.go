package governance

import "github.com/okex/okchain-go-sdk/module/governance/types"

// const
const (
	ModuleName = types.ModuleName
)

type (
	// Proposal is the type alias of the one under governance/types
	Proposal = types.Proposal
	// ProposalStatus is the type alias of the one under governance/types
	ProposalStatus = types.ProposalStatus
	// TallyResult is the type alias of the one under governance/types
	TallyResult = types.TallyResult
	// TextProposal is the type alias of the one under governance/types
	TextProposal = types.TextProposal
)