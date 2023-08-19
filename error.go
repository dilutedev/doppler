package doppler

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidToken = errors.New("received invalid token")
)

type DopplerError struct {
	Status   int      `json:"status"`
	Messages []string `json:"message"`
}

func (err *DopplerError) Error() string {
	return fmt.Sprintf(`request failed with Code: %d, Message: %v`, err.Status, err.Messages[0])
}
