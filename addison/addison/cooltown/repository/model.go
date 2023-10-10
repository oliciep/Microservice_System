package repository

// Structure for metadata, contains Id of track along with audio stored as base64 json string and the artist of the track
type Metadata struct {
	Id      string
	Title   string
	Artist  string
}

// Structure for a sample, contains Audio of track along with Audio stored as base64 json string
type Sample struct {
	Audio   string
}

// Structure for a track, contains Id of track along with audio stored as base64 json string
type Track struct {
	Id 	string
	Audio	string
}

// Structure for a result, contains Id of track
type Result struct {
	Id	string
}
