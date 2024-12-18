package pkgs

import "time"

// Process Name Constants
// process : identifier
const (
	ProcessEvents                 = "Validator: ProcessEvents"
	BuildMerkleTree               = "Validator: BuildMerkleTree"
	ContractSubmission            = "Validator: ContractSubmission"
	StartFetchingBlocks           = "Validator: StartFetchingBlocks"
	FetchBatchSubmission          = "Validator: FetchBatchSubmission"
	StoreBatchSubmission          = "Validator: StoreBatchSubmission"
	SendBatchAttestationToRelayer = "Validator: SendBatchAttestationToRelayer"
)

// General Key Constants
const (
	ValidatorKey    = "ValidatorKey"
	ValidatorSetKey = "ValidatorSetKey"
)

// General Constants
const (
	Day = 24 * time.Hour
)
