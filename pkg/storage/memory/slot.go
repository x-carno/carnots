package memory

type Slot struct {
	Value     float64
	Timestamp int64

	// // first Slot's dod is the delta from MinTimestamp,
	// Dod []byte
}
