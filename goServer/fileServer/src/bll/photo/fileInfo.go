package photo

import (
	"time"
)

type fileInfo struct {
	// FileName 文件名
	FileName string

	// DirName 文件夹名字
	DirName string

	// ModName 修改时间
	ModName time.Time
}
