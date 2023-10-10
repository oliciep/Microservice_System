package repository

// Structure for a track, contains Id of track along with audio stored as base64 json string
type Track struct {
	Id      string
	Audio   string
}
