package proto

// Wire types duplicated per repo. JSON tags support REST stubs;
// binary encoding for the gateway wire is still TBD.

type Hello struct {
	Version    string `json:"version"`
	PlayerName string `json:"playerName"`
}

type Auth struct {
	Token string `json:"token"`
}

type ChunkRequest struct {
	GameID string `json:"gameId"`
	ChunkX int32  `json:"chunkX"`
	ChunkY int32  `json:"chunkY"`
	ChunkZ int32  `json:"chunkZ"`
}

type ChunkData struct {
	GameID string `json:"gameId"`
	ChunkX int32  `json:"chunkX"`
	ChunkY int32  `json:"chunkY"`
	ChunkZ int32  `json:"chunkZ"`
	Voxels []byte `json:"voxels,omitempty"`
}

type Edit struct {
	X        int32  `json:"x"`
	Y        int32  `json:"y"`
	Z        int32  `json:"z"`
	OldBrick uint8  `json:"oldBrick"`
	NewBrick uint8  `json:"newBrick"`
}

type ChunkDiff struct {
	GameID string `json:"gameId"`
	ChunkX int32  `json:"chunkX"`
	ChunkY int32  `json:"chunkY"`
	ChunkZ int32  `json:"chunkZ"`
	Edits  []Edit `json:"edits"`
}

type PlayerInput struct {
	InputSeq uint64  `json:"inputSeq"`
	MoveX    float32 `json:"moveX"`
	MoveY    float32 `json:"moveY"`
	Jump     bool    `json:"jump"`
	Tick     uint64  `json:"tick"`
}

type EntityState struct {
	EntityID  string  `json:"entityId"`
	X         float32 `json:"x"`
	Y         float32 `json:"y"`
	Z         float32 `json:"z"`
	Yaw       float32 `json:"yaw"`
	Pitch     float32 `json:"pitch"`
	BrickType uint8   `json:"brickType"`
}

type Snapshot struct {
	Seq          uint64        `json:"seq"`
	LastInputSeq uint64        `json:"lastInputSeq"`
	Entities     []EntityState `json:"entities"`
}

type OccurrenceLoc struct {
	GameID  string `json:"gameId"`
	ChunkX  int32  `json:"chunkX"`
	ChunkY  int32  `json:"chunkY"`
	ChunkZ  int32  `json:"chunkZ"`
	OffsetX int32  `json:"offsetX"`
	OffsetZ int32  `json:"offsetZ"`
}

type Occurrence struct {
	Type      string        `json:"type"`
	Location  OccurrenceLoc `json:"location"`
	Magnitude float32       `json:"magnitude"`
	Duration  uint32        `json:"duration"`
	StartTick uint64        `json:"startTick"`
}

type StateSummary struct {
	PlayerCount   int32    `json:"playerCount"`
	Regions       []string `json:"regions"`
	ActiveWeather string   `json:"activeWeather"`
	Tick          uint64   `json:"tick"`
}

type AuthToken struct {
	Token     string `json:"token"`
	PlayerID  string `json:"playerId"`
	ExpiresAt int64  `json:"expiresAt"`
}

type WorldMeta struct {
	WorldID     string `json:"worldId"`
	Name        string `json:"name"`
	Seed        int64  `json:"seed"`
	PlayerCount int32  `json:"playerCount"`
}
