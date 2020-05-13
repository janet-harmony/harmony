package apiv1

import (
  "errors"
  "os/exec"
)

var (
  killCmd = exec.Command("pkill", "-9", "harmony")
  safeKillCmd = exec.Command("pkill", "harmony")
)

func kill(safe bool) {
  if safe {
    safeKillCmd.Run()
  }
  killCmd.Run()
}

/*
curl --location --request POST http://localhost:9502 \
--header 'Content-Type: application/json' \
--data-raw '{
    "jsonrpc": "2.0",
    "method": "hmy_killPrepare",
    "params": [true],
    "id": 1
}'
*/
func(s *PublicBlockChainAPI) KillNode(safe bool) {
  kill(safe)
}

func(s *PublicBlockChainAPI) KillLeader(safe bool) error {
  if node := s.b.GetNodeMetadata(); node.IsLeader {
    kill(safe)
  }
  return errors.New("Not leader")
}

func (s *PublicBlockChainAPI) KillLeaderAnnounce(safe bool) error {
  if node := s.b.GetNodeMetadata(); node.IsLeader {
    if phase := s.b.GetConsensusPhase(); phase == "Announce" {
      kill(safe)
    }
    return errors.New("Not in Announce phase")
  }
  return errors.New("Not leader")
}

func (s *PublicBlockChainAPI) KillLeaderPrepare(safe bool) error {
  if node := s.b.GetNodeMetadata(); node.IsLeader {
    if phase := s.b.GetConsensusPhase(); phase == "Prepare" {
      kill(safe)
    }
    return errors.New("Not in Prepare phase")
  }
  return errors.New("Not leader")
}

func (s *PublicBlockChainAPI) KillLeaderCommit(safe bool) error {
  if node := s.b.GetNodeMetadata(); node.IsLeader {
    if phase := s.b.GetConsensusPhase(); phase == "Commit" {
      kill(safe)
    }
    return errors.New("Not in Commit phase")
  }
  return errors.New("Not leader")
}

func (s *PublicBlockChainAPI) KillAnnounce(safe bool) error {
  if node := s.b.GetNodeMetadata(); !node.IsLeader {
    if phase := s.b.GetConsensusPhase(); phase == "Announce" {
      kill(safe)
    }
    return errors.New("Not in Announce phase")
  }
  return errors.New("Is leader")
}

func (s *PublicBlockChainAPI) KillPrepare(safe bool) error {
  if node := s.b.GetNodeMetadata(); !node.IsLeader {
    if phase := s.b.GetConsensusPhase(); phase == "Prepare" {
      kill(safe)
    }
    return errors.New("Not in Prepare phase")
  }
  return errors.New("Is leader")
}

func (s *PublicBlockChainAPI) KillCommit(safe bool) error {
  if node := s.b.GetNodeMetadata(); !node.IsLeader {
    if phase := s.b.GetConsensusPhase(); phase == "Commit" {
      kill(safe)
    }
    return errors.New("Not in Commit phase")
  }
  return errors.New("Is leader")
}
