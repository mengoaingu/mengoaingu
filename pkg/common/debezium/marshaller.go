package debezium

// import (
// 	"encoding/json"

// 	"github.com/ThreeDotsLabs/watermill/message"
// )

// func UnmarshallFunc(msg *message.Message) (messaging.NamedMessage, error) {
// 	b := msg.Payload
// 	var e Event[any]
// 	err := json.Unmarshal(b, &e)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return e, nil
// }
