package fs

import (
	"context"
	"fmt"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"io"
	"time"
)

type QiniuFS struct {
	isPrivate bool
	domain    string
	bucket    string
	cfg       *storage.Config
	mac       *qbox.Mac
}

func NewQiniuFS(accessKey, secretKey, domain, bucket string, zone *storage.Zone, isPrivate bool) FS {
	return &QiniuFS{
		isPrivate: isPrivate,
		domain:    domain,
		bucket:    bucket,
		mac:       qbox.NewMac(accessKey, secretKey),
		cfg:       &storage.Config{Zone: zone},
	}
}

func (q *QiniuFS) UploadFilePath(fpath, key string) (string, error) {
	return q.uploadFilePath(q.bucket, fpath, key)
}

func (q *QiniuFS) UploadFileReader(reader io.Reader, key string, fsize int64) (string, error) {
	return q.uploadFileReader(q.bucket, reader, key, fsize)
}

func (q *QiniuFS) UpdateFilePath(fpath, key string) (string, error) {
	return q.uploadFilePath(fmt.Sprintf("%s:%s", q.bucket, key), fpath, key)
}

func (q *QiniuFS) UpdateFileReader(reader io.Reader, key string, fsize int64) (string, error) {
	return q.uploadFileReader(fmt.Sprintf("%s:%s", q.bucket, key), reader, key, fsize)
}

func (q *QiniuFS) DeleteFile(key string) error {
	bucketManager := storage.NewBucketManager(q.mac, q.cfg)
	return bucketManager.Delete(q.bucket, key)
}

func (q *QiniuFS) GetDownloadUrl(key string) string {
	if q.isPrivate {
		return storage.MakePrivateURL(q.mac, q.domain, key, time.Now().Add(time.Minute*2).Unix())
	}
	return storage.MakePublicURL(q.domain, key)
}

func (q *QiniuFS) GetToken(typ TokenType, fid string) (string, error) {
	panic("qiniuFS does not implement yet !")
}

func (q *QiniuFS) DownloadFile(fid string) (string, []byte, error) {
	panic("qiniuFS does not implement yet!")
}

func (q *QiniuFS) uploadFilePath(scope, fpath, key string) (string, error) {
	putPolicy := storage.PutPolicy{
		Scope: scope,
	}
	formUploader := storage.NewFormUploader(q.cfg)
	err := formUploader.PutFile(context.Background(), &storage.PutRet{}, putPolicy.UploadToken(q.mac), key, fpath, nil)
	return key, err
}

func (q *QiniuFS) uploadFileReader(scope string, reader io.Reader, key string, fsize int64) (string, error) {
	putPolicy := storage.PutPolicy{
		Scope: scope,
	}
	formUploader := storage.NewFormUploader(q.cfg)
	err := formUploader.Put(context.Background(), &storage.PutRet{}, putPolicy.UploadToken(q.mac), key, reader, fsize, nil)
	return key, err
}
