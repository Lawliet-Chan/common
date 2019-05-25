package fs

import "io"

type FS interface {
	UploadFilePath(fpath, key string) (string, error)
	UploadFileReader(reader io.Reader, key string, fsize int64) (string, error)
	GetDownloadUrl(key string) string
	DownloadFile(fid string) (string, []byte, error)
	UpdateFilePath(fpath, key string) (string, error)
	UpdateFileReader(reader io.Reader, key string, fsize int64) (string, error)
	DeleteFile(key string) error
}
