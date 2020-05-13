package hmy

// GetConsensusPhase ..
func (b *APIBackend) GetConsensusPhase() string {
  return b.hmy.nodeAPI.GetConsensusPhase()
}

// GetConsensusMode ..
func (b *APIBackend) GetConsensusMode() string {
	return b.hmy.nodeAPI.GetConsensusMode()
}
