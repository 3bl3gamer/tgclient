package mtproto

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"

	"github.com/ansel1/merry"
)

var ErrNoSessionData = merry.New("no session data")

type SessionStore interface {
	Save(*SessionInfo) error
	Load(*SessionInfo) error
}

type SessNoopStore struct{}

func (s *SessNoopStore) Save(sess *SessionInfo) error { return nil }
func (s *SessNoopStore) Load(sess *SessionInfo) error { return ErrNoSessionData.Here() }

type SessFileStore struct {
	FPath string
}

func (s *SessFileStore) Save(sess *SessionInfo) (err error) {
	f, err := os.Create(s.FPath + ".temp")
	if err != nil {
		return merry.Wrap(err)
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "\t")
	if err := encoder.Encode(sess); err != nil {
		return merry.Wrap(err)
	}
	if err := f.Close(); err != nil {
		return merry.Wrap(err)
	}

	if err := os.Rename(s.FPath+".temp", s.FPath); err != nil {
		return merry.Wrap(err)
	}
	return nil
}

func (s *SessFileStore) Load(sess *SessionInfo) error {
	f, err := os.Open(s.FPath)
	if errors.Is(err, fs.ErrNotExist) {
		return ErrNoSessionData.Here().WithCause(err)
	}
	if err != nil {
		return merry.Wrap(err)
	}
	defer f.Close()

	if err := json.NewDecoder(f).Decode(sess); err != nil {
		return merry.Wrap(err)
	}
	return nil
}

type SessFileStoreExt struct {
	SessFileStore
	RequirePresent   bool // Session file must be already present
	IgnoreUnreadable bool // Session file can be unreadable and will be overwritten, if unreadable and writeable
}

func (s *SessFileStoreExt) Load(sess *SessionInfo) error {
	err := s.SessFileStore.Load(sess)
	if s.RequirePresent && errors.Is(err, ErrNoSessionData) {
		causeErr := merry.Cause(err)
		if causeErr != nil {
			return merry.Wrap(causeErr)
		}
		return merry.Wrap(err)
	}
	if s.IgnoreUnreadable && err != nil {
		return ErrNoSessionData.Here()
	}
	return merry.Wrap(err)
}
