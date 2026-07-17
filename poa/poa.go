package poa

type StarchBlock struct {
  ID           uint        `cbor:"0,keyasint" json:"block_id"`
  Attendance   []string    `cbor:"1,keyasint" json:"attendance"`
  VRF          string      `cbor:"2,keyasint" json:"vrf"`
  Winner       ProofBlock  `cbor:"3,keyasint" json:"winner"`
  TeamConfig   string      `cbor:"4,keyasint" json:"team_config"`
  BlockReward  uint        `cbor:"5,keyasint" json:"block_reward"`
  StakeReward  StakeReward `cbor:"6,keyasint" json:"stake_reward"`
  Distribution []Reward    `cbor:"5,keyasint" json:"distribution"`
  Era          string      `cbor:"6,keyasint" json:"era"`
  Timestamp    uint        `cbor:"7,keyasint" json:"timestamp"`
  Version      uint        `cbor:"8,keyasint" json:"version"`
}

type ProofBlock struct {
  MinerID  string `cbor:"0,keyasint" json:"miner_id"`
  PrevHash string `cbor:"1,keyasint" json:"previous_hash"`
  Color    string `cbor:"2,keyasint" json:"color"`
}

type StakeReward struct {
  Epoch  uint   `cbor:"0,keyasint" json:"epoch"`
  Ticker string `cbor:"1,keyasint" json:"ticker"`
  Amount uint   `cbor:"2,keyasint" json:"amount"`
}

type Reward struct {
  To     string `cbor:"0,keyasint" json:"to"`
  Amount uint   `cbor:"1,keyasint" json:"amount"`
}
