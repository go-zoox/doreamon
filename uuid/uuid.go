package uuid

import (
	"time"

	_uuid "github.com/google/uuid"
)

func UUID() string {
	return V4()
}

func V4() string {
	instance := _uuid.New()
	return instance.String()
}

func V5() string {
	instance := _uuid.NewSHA1(_uuid.NameSpaceURL, []byte(time.Now().String()))

	return instance.String()
}

func V1() string {
	instance, err := _uuid.NewUUID()
	if err != nil {
		return ""
	}

	return instance.String()
}

func V2() string {
	instance, err := _uuid.NewDCEPerson()
	if err != nil {
		return ""
	}

	return instance.String()
}

func V3() string {
	instance := _uuid.NewMD5(_uuid.NameSpaceURL, []byte(time.Now().String()))

	return instance.String()
}
