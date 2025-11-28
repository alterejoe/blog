package interfaces

import (
	"context"
	"io"
	"net/http"

	"github.com/google/uuid"
)

type User interface {
	ID() uuid.UUID
	Name() string
	Email() string
	// GroupID() string
}

type Sanitizer any

type Logger interface {
	Debug(string, ...any)
	Info(string, ...any)
	Warn(string, ...any)
	Error(string, ...any)
}

type SessionManager interface {
	LoadAndSave(http.Handler) http.Handler
	Get(r context.Context, key string) any
	Put(r context.Context, key string, value any)
	Remove(r context.Context, key string)
	Destroy(ctx context.Context) error
	Token(ctx context.Context) string
	SetFlashMessage(r context.Context, t, msg string)
	GetFlashMessage(r context.Context) (string, string)

	SetAuthUser(ctx context.Context, user User)
	SetElectionID(ctx context.Context, electionID uuid.UUID)
	GetElectionID(ctx context.Context) uuid.UUID
	// GetAuthUser(ctx context.Context) User
	DeleteAuthUser(ctx context.Context)
	GetGroupName(ctx context.Context) string
	GetAdminGroupID(ctx context.Context) uuid.UUID
	GetClientGroupID(ctx context.Context) uuid.UUID
}

type Response interface {
	ServerSuccess(w http.ResponseWriter, r *http.Request)
	ServerCreated(w http.ResponseWriter, r *http.Request)
	ServerError(w http.ResponseWriter, r *http.Request, err error)
	ClientError(w http.ResponseWriter, r *http.Request, err error)
}

type S3 interface {
	UploadFile(ctx context.Context, key string, localPath string) error
	UploadReader(ctx context.Context, key string, r io.Reader) error
	DownloadFile(ctx context.Context, key string, localPath string) error
	ListGroupFiles(ctx context.Context, groupID string) ([]string, error)
	AddFileToGroup(ctx context.Context, groupID, fileName string, r io.Reader) error
	DeleteFileFromGroup(ctx context.Context, groupID, fileName string) error
	FileExists(ctx context.Context, groupID, filename string) (bool, error)
	GenerateKey(ctx context.Context, groupID string, filename string) string
	FileURL(ctx context.Context, key string) string
}
