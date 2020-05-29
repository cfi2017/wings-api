package model

import (
	"os"
	"sync"
)

type Backup interface {
	// Returns the UUID of this backup as tracked by the panel instance.
	Identifier() string

	// Generates a backup in whatever the configured source for the specific
	// implementation is.
	Generate(*IncludedFiles, string) (*ArchiveDetails, error)

	// Returns the ignored files for this backup instance.
	Ignored() []string

	// Returns a SHA256 checksum for the generated backup.
	Checksum() ([]byte, error)

	// Returns the size of the generated backup.
	Size() (int64, error)

	// Returns the path to the backup on the machine. This is not always the final
	// storage location of the backup, simply the location we're using to store
	// it until it is moved to the final spot.
	Path() string

	// Returns details about the archive.
	Details() *ArchiveDetails

	// Removes a backup file.
	Remove() error
}

type ArchiveDetails struct {
	Checksum string `json:"checksum"`
	Size     int64  `json:"size"`
}

type IncludedFiles struct {
	sync.RWMutex
	files map[string]*os.FileInfo
}

// Pushes an additional file or folder onto the struct.
func (i *IncludedFiles) Push(info *os.FileInfo, p string) {
	i.Lock()
	defer i.Unlock()

	if i.files == nil {
		i.files = make(map[string]*os.FileInfo)
	}

	i.files[p] = info
}

// Returns all of the files that were marked as being included.
func (i *IncludedFiles) All() map[string]*os.FileInfo {
	i.RLock()
	defer i.RUnlock()

	return i.files
}
