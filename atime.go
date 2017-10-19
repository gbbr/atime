package atime

import (
	"os"
	"sort"
)

type Ascending []os.FileInfo

var _ sort.Interface = (*Ascending)(nil)

func (fis Ascending) Len() int           { return len(fis) }
func (fis Ascending) Swap(i, j int)      { fis[i], fis[j] = fis[j], fis[i] }
func (fis Ascending) Less(i, j int) bool { return atime(fis[i]).Before(atime(fis[j])) }
