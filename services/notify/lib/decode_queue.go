package lib

import (
	"encoding/json"

	"github.com/web-notify/api/monorepo/libs/utils/formats"
	"github.com/web-notify/api/monorepo/services/notify/models"
)

func DecodeRawMassageToInfo(rawMassage json.RawMessage) (models.Info, error) {

	// the message is stringified twice, so need to unmarshall twice
	var message string
	err := json.Unmarshal(rawMassage, &message)
	if err != nil {
		return models.Info{}, err
	}
	formats.Trace("rawMessage:", message)
	err = json.Unmarshal([]byte(message), &message)
	if err != nil {
		return models.Info{}, err
	}

	var info models.Info
	err = json.Unmarshal([]byte(message), &info)
	if err != nil {
		return models.Info{}, err
	}
	formats.Trace("info:", info)

	return info, nil
}
