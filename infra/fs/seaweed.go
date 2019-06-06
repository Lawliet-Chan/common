package fs

import (
	"fmt"
	"github.com/CrocdileChan/common/logger"
	"github.com/CrocdileChan/goseaweedfs"
	"io"
	"net/http"
	"time"
)

type SeaweedFS struct {
	cli *goseaweedfs.Seaweed
}

func NewSeaweedFS(scheme, master string, filers []string, chunkSize int64, timeOut string) *SeaweedFS {
	timeout, err := time.ParseDuration(timeOut)
	if err != nil {
		logger.GetLogger().Fatalf("seaweedfs error : %v", err)
	}
	cli := goseaweedfs.NewSeaweed(scheme, master, filers, chunkSize, timeout)
	return &SeaweedFS{
		cli: cli,
	}
}

// @return fileName,fileData,error
func (sfs *SeaweedFS) DownloadFile(fid string) (string, []byte, error) {
	return sfs.cli.DownloadFile(fid, nil)
}

func (sfs *SeaweedFS) UploadFilePath(fpath, fname string) (string, error) {
	_, _, fid, err := sfs.cli.UploadFile(fpath, "", "")
	if err != nil {
		return "", err
	}
	return /*sfs.cli.Scheme + "://" + fpart.Server + "/" + */ fid, nil
}

func (sfs *SeaweedFS) UploadFileReader(reader io.Reader, fname string, fsize int64) (string, error) {
	_, fid, err := sfs.cli.Upload(reader, fname, fsize, "", "")
	if err != nil {
		return "", err
	}
	return /*sfs.cli.Scheme + "://" + fpart.Server + "/" + */ fid, nil
}

func (sfs *SeaweedFS) DeleteFile(fid string) error {
	return sfs.cli.DeleteFile(fid, nil)
}

func (sfs *SeaweedFS) UpdateFilePath(fpath, key string) (string, error) {
	panic("seaweedFS does not support update by file path")
}

//不建议用
func (sfs *SeaweedFS) UpdateFileReader(reader io.Reader, fid string, fsize int64) (string, error) {
	newFid, err := sfs.UploadFileReader(reader, "", fsize)
	if err != nil {
		return "", err
	}
	return newFid, sfs.DeleteFile(fid)
}

func (sfs *SeaweedFS) GetDownloadUrl(fid string) string {
	return fid
}

func (sfs *SeaweedFS) GetToken(typ TokenType, fid string) (string, error) {
	var getTokenUrl string
	switch typ {
	case UploadToken:
		getTokenUrl = fmt.Sprintf("%s://%s/dir/assign", sfs.cli.Scheme, sfs.cli.Master)
	case DownloadToken:
		getTokenUrl = fmt.Sprintf("%s://%s/dir/lookup?fileId=%s&read=yes", sfs.cli.Scheme, sfs.cli.Master, fid)
	case UpdateToken:
		getTokenUrl = fmt.Sprintf("%s://%s/dir/lookup?fileId=%s", sfs.cli.Scheme, sfs.cli.Master, fid)
	case DeleteToken:
		getTokenUrl = fmt.Sprintf("%s://%s/dir/lookup?fileId=%s", sfs.cli.Scheme, sfs.cli.Master, fid)
	}
	resp, err := http.Get(getTokenUrl)
	if err != nil {
		return "", err
	}
	return resp.Header.Get("Authorization"), nil
}
