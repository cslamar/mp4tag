package mp4tag

import "os"

type MP4File struct {
	f         *os.File
	trackPath string
}

type Tags struct {
	Album          string
	AlbumArtist    string
	AlbumSort      string
	Artist         string
	Comment        string
	Composer       string
	Copyright      string
	CoversData     [][]byte
	Custom         map[string]string
	Delete         []string
	DiskNumber     int
	DiskTotal      int
	Genre          string
	Label          string
	Title          string
	TitleSort      string
	TrackNumber    int
	TrackTotal     int
	UnsyncedLyrics string
	Year           int
	yearStr        string
	// ItunesAdvisory 	  int
	// itunesAdvisoryStr string
}
