package utils

import (
	"github.com/gofrs/uuid"
	"github.com/jxncyjq/lib_stardust/core/errors"
)

// ErrGeneratingID indicates error in generating UUID
var ErrGeneratingID = errors.New("generating id failed", 0)

var _ IDProvider = (*uuidProvider)(nil)

// IDProvider specifies an API for generating unique identifiers.
type IDProvider interface {
	// ID generates the unique identifier.
	ID() (string, error)
}

type uuidProvider struct {
	option UuidOption
}

type UuidOption struct {
	IDType int
	Domain byte      //uuid2 option
	Ns     uuid.UUID //uuid3,uuid5 option
	Name   string    //uuid3,uuid5 option
}

// New instantiates a UUID provider.
func NewUUID(opt UuidOption) IDProvider {
	return &uuidProvider{option: UuidOption{}}
}

func (up *uuidProvider) ID() (string, error) {
	var id uuid.UUID
	var err error
	switch up.option.IDType {
	case 1:
		id, err = uuid.NewV1()
	case 2:
		id, err = uuid.NewV2(up.option.Domain)
	case 3:
		id = uuid.NewV3(up.option.Ns, up.option.Name)
	case 4:
		id, err = uuid.NewV4()
	case 5:
		id = uuid.NewV5(up.option.Ns, up.option.Name)
	default:
		id, err = uuid.NewV1()
	}
	if err != nil {
		return "", errors.WithStack(ErrGeneratingID, 0)
	}

	return id.String(), nil
}
