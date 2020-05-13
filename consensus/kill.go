package consensus


// GetConsensusPhase ..
func (c *Consensus) GetConsensusPhase() string {
  return c.phase.String()
}

// GetConsensusMode ..
func (c *Consensus) GetConsensusMode() string {
  return c.current.mode.String()
}
