//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2018] Last.Backend LLC
// All Rights Reserved.
//
// NOTICE:  All information contained herein is, and remains
// the property of Last.Backend LLC and its suppliers,
// if any.  The intellectual and technical concepts contained
// herein are proprietary to Last.Backend LLC
// and its suppliers and may be covered by Russian Federation and Foreign Patents,
// patents in process, and are protected by trade secret or copyright law.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from Last.Backend LLC.
//

package state

import (
	"github.com/lastbackend/lastbackend/pkg/distribution/types"
)

const logLevel = 3

type State struct {
	discovery *DiscoveryState
	networks  *NetworkState
}

type DiscoveryState struct {
	Info   types.DiscoveryInfo
	Status types.DiscoveryStatus
}

func (s *State) Discovery() *DiscoveryState {
	return s.discovery
}

func (s *State) Networks() *NetworkState {
	return s.networks
}


func New() *State {

	state := State{
		discovery: new(DiscoveryState),
		networks: &NetworkState{
			subnets: make(map[string]types.NetworkState, 0),
		},
	}

	return &state
}