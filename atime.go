// Package atime provides some minimal utilities in working with file access times.
package atime

import (
	"os"
	"sort"
	"time"
)

type Ascending []os.FileInfo

var _ sort.Interface = (*Ascending)(nil)

func (fis Ascending) Len() int           { return len(fis) }
func (fis Ascending) Swap(i, j int)      { fis[i], fis[j] = fis[j], fis[i] }
func (fis Ascending) Less(i, j int) bool { return atime(fis[i]).Before(atime(fis[j])) }

// Get returns the access time from the given os.FileInfo.
func Get(fi os.FileInfo) time.Time { return atime(fi) }
