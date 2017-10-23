package photo

import (
	"time"
)

type fileInfo struct {
	// FileName 文件名
	FileName string

	// ModName 修改时间
	ModName time.Time
}
