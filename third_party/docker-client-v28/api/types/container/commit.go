package container

import "github.com/strangelove-ventures/tokenfactory/third_party/docker-client/api/types/common"

// CommitResponse response for the commit API call, containing the ID of the
// image that was produced.
type CommitResponse = common.IDResponse
