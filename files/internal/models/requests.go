package models

type UploadToStorageReq struct {
	UserID string
	File   []byte
	FileID string
}
