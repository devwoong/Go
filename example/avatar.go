package example

import "errors"

var ErrorNoAvatar = errors.New("Unable to get an avatar URL.")

type Avatar interface {
	GetAvatarURL() (string, error)
}
