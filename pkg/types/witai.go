package types

type BenderContextKey string

type WitEntity struct {
	Confidence float64
	Value      string
}

type BenderWit struct {
	Intent    []WitEntity
	DbType    []WitEntity
	Artifact  []WitEntity
	Cluster   []WitEntity
	Sentiment []WitEntity
}
