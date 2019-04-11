package devicecheck

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/google/uuid"
)

const queryTwoBitsPath = "/query_two_bits"

type queryTwoBitsRequestBody struct {
	DeviceToken   string `json:"device_token"`
	TransactionID string `json:"transaction_id"`
	Timestamp     int64  `json:"timestamp"`
}

type QueryTwoBitsResult struct {
	Bit0           bool `json:"bit0"`
	Bit1           bool `json:"bit1"`
	LastUpdateTime Time `json:"last_update_time"`
}

type Time struct {
	time.Time
}

const timeFormat = "2006-01"

func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Format(timeFormat))
}

func (t *Time) UnmarshalJSON(b []byte) (err error) {
	t.Time, err = time.Parse(timeFormat, strings.Trim(string(b), `"`))
	return
}

func (api api) queryTwoBits(deviceToken, jwt string) (int, []byte, error) {
	b := queryTwoBitsRequestBody{
		DeviceToken:   deviceToken,
		TransactionID: uuid.New().String(),
		Timestamp:     time.Now().UTC().UnixNano() / int64(time.Millisecond),
	}

	return api.do(jwt, queryTwoBitsPath, b)
}

func (client Client) QueryTwoBits(deviceToken string, result *QueryTwoBitsResult) error {
	key, err := client.cred.Key()
	if err != nil {
		return err
	}

	jwt, err := client.jwt.generate(key)
	if err != nil {
		return err
	}

	code, body, err := client.api.queryTwoBits(deviceToken, jwt)
	if err != nil {
		return err
	}

	if err := newError(code, body); err != nil {
		return err
	}

	return json.Unmarshal(body, result)
}
