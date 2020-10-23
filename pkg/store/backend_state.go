package store

import (
	"gitlab.com/infra.run/public/b3scale/pkg/bbb"
)

// The BackendState is shared across b3scale instances
// and encapsulates the list of meetings and recordings.
// The backend.ID should be used as identifier.
type BackendState struct {
	ID string

	NodeState  string
	AdminState string

	LastError string

	Host   string
	Secret string

	Tags []string

	// db storage
}

// GetBackendsOpts provides a set of possible filters
type GetBackendsOpts struct {
}

// GetBackends retrievs all backend states, filterable with opts
func GetBackends(opts *GetBackendsOpts) []*BackendState {
	return nil
}

// Meetings
func (s *BackendState) GetMeetings() (bbb.MeetingsCollection, error) {
	return nil, nil
}

// SetMeetings
func (s *BackendState) SetMeetings(bbb.MeetingsCollection) error {
	return nil
}

func (s *BackendState) AddMeeting(*bbb.Meeting) error {
	return nil
}

/*

	// Recordings
	GetRecordings(*cluster.Backend) (bbb.RecordingsCollection, error)
	SetRecordings(*cluster.Backend, bbb.RecordingsCollection) error

	// Forget about the backend
	Delete(*cluster.Backend)

*/
