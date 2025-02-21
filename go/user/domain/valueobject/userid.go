package valueobject

import "github.com/oklog/ulid/v2"

type UserID struct {
	v ulid.ULID
}

func NewUserID() UserID {
	id := ulid.Make()

	return UserID{
		v: id,
	}
}

func ParseUserID(v string) (*UserID, error) {
	id, err := ulid.Parse(v)

	if err != nil {
		return nil, err
	}

	return &UserID{v: id}, nil
}

func (u *UserID) Value() *ulid.ULID {
	return &u.v
}
