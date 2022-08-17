package consumer_handler

import (
	"encoding/json"

	"github.com/eugbyte/monorepo/services/webnotify/models"
)

func decodeRawMassageToInfo(rawMassage json.RawMessage) (models.MessageInfo, error) {

	// the message is stringified twice, so need to unmarshall twice
	var message string
	err := json.Unmarshal(rawMassage, &message)
	if err != nil {
		return models.MessageInfo{}, err
	}
	err = json.Unmarshal([]byte(message), &message)
	if err != nil {
		return models.MessageInfo{}, err
	}

	var info models.MessageInfo
	err = json.Unmarshal([]byte(message), &info)
	if err != nil {
		return models.MessageInfo{}, err
	}

	return info, nil
}
