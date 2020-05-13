package node

// GetConsensusPhase ..
func (n *Node) GetConsensusPhase() string {
	return n.Consensus.GetConsensusPhase()
}

// GetConsensusMode ..
func (n *Node) GetConsensusMode() string {
  return n.Consensus.GetConsensusMode()
}
