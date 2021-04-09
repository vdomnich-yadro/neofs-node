package managers

import (
	"github.com/nspcc-dev/neofs-node/pkg/services/reputation"
	reputationroute "github.com/nspcc-dev/neofs-node/pkg/services/reputation/local/route"
	"github.com/pkg/errors"
)

// NextStage builds Manager list for node and returns it directly.
//
// If passed route has more than one point, then endpoint of the route is reached.
func (b *Builder) NextStage(epoch uint64, p reputation.PeerID, passed []reputationroute.ServerInfo) ([]reputationroute.ServerInfo, error) {
	if len(passed) > 1 {
		return nil, nil
	}

	route, err := b.managerBuilder.BuildManagers(epoch, p)
	if err != nil {
		return nil, errors.Wrapf(err, "could not build managers for epoch: %d", epoch)
	}

	return route, nil
}