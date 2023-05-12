package debezium

// import (
// 	"encoding/json"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"
// )

// type Inventory struct {
// 	ID              int64     `json:"id"`
// 	UserID          string    `json:"user_id"`
// 	Name            string    `json:"name"`
// 	URL             string    `json:"url"`
// 	MediaChannelID  int64     `json:"media_channel_id"`
// 	MinViews        int64     `json:"min_views"`
// 	MaxViews        int64     `json:"max_views"`
// 	MinInteractions int64     `json:"min_interactions"`
// 	MaxInteractions int64     `json:"max_interactions"`
// 	CreatedAt       time.Time `json:"created_at"`
// 	UpdatedAt       time.Time `json:"updated_at"`
// 	DeletedAt       time.Time `json:"deleted_at"`
// 	Members         int64     `json:"members"`
// 	Followers       int64     `json:"followers"`
// }

// func TestEvent_Unmarshall(t *testing.T) {
// 	jsonData := `{
//     "schema": {
//         "type": "struct",
//         "fields": [
//             {
//                 "type": "struct",
//                 "fields": [
//                     {
//                         "type": "int64",
//                         "optional": false,
//                         "field": "id"
//                     },
//                     {
//                         "type": "string",
//                         "optional": false,
//                         "field": "user_id"
//                     },
//                     {
//                         "type": "string",
//                         "optional": false,
//                         "field": "name"
//                     },
//                     {
//                         "type": "string",
//                         "optional": true,
//                         "field": "url"
//                     },
//                     {
//                         "type": "int64",
//                         "optional": true,
//                         "field": "media_channel_id"
//                     },
//                     {
//                         "type": "int64",
//                         "optional": true,
//                         "default": 0,
//                         "field": "min_views"
//                     },
//                     {
//                         "type": "int64",
//                         "optional": true,
//                         "default": 0,
//                         "field": "max_views"
//                     },
//                     {
//                         "type": "int64",
//                         "optional": true,
//                         "default": 0,
//                         "field": "min_interactions"
//                     },
//                     {
//                         "type": "int64",
//                         "optional": true,
//                         "default": 0,
//                         "field": "max_interactions"
//                     },
//                     {
//                         "type": "string",
//                         "optional": false,
//                         "name": "io.debezium.time.ZonedTimestamp",
//                         "version": 1,
//                         "default": "1970-01-01T00:00:00Z",
//                         "field": "created_at"
//                     },
//                     {
//                         "type": "string",
//                         "optional": false,
//                         "name": "io.debezium.time.ZonedTimestamp",
//                         "version": 1,
//                         "default": "1970-01-01T00:00:00Z",
//                         "field": "updated_at"
//                     },
//                     {
//                         "type": "string",
//                         "optional": true,
//                         "name": "io.debezium.time.ZonedTimestamp",
//                         "version": 1,
//                         "field": "deleted_at"
//                     },
//                     {
//                         "type": "int64",
//                         "optional": true,
//                         "default": 0,
//                         "field": "members"
//                     },
//                     {
//                         "type": "int64",
//                         "optional": true,
//                         "default": 0,
//                         "field": "followers"
//                     }
//                 ],
//                 "optional": true,
//                 "name": "bmarketplace.profile.inventory.GetValue",
//                 "field": "before"
//             },
//             {
//                 "type": "struct",
//                 "fields": [
//                     {
//                         "type": "int64",
//                         "optional": false,
//                         "field": "id"
//                     },
//                     {
//                         "type": "string",
//                         "optional": false,
//                         "field": "user_id"
//                     },
//                     {
//                         "type": "string",
//                         "optional": false,
//                         "field": "name"
//                     },
//                     {
//                         "type": "string",
//                         "optional": true,
//                         "field": "url"
//                     },
//                     {
//                         "type": "int64",
//                         "optional": true,
//                         "field": "media_channel_id"
//                     },
//                     {
//                         "type": "int64",
//                         "optional": true,
//                         "default": 0,
//                         "field": "min_views"
//                     },
//                     {
//                         "type": "int64",
//                         "optional": true,
//                         "default": 0,
//                         "field": "max_views"
//                     },
//                     {
//                         "type": "int64",
//                         "optional": true,
//                         "default": 0,
//                         "field": "min_interactions"
//                     },
//                     {
//                         "type": "int64",
//                         "optional": true,
//                         "default": 0,
//                         "field": "max_interactions"
//                     },
//                     {
//                         "type": "string",
//                         "optional": false,
//                         "name": "io.debezium.time.ZonedTimestamp",
//                         "version": 1,
//                         "default": "1970-01-01T00:00:00Z",
//                         "field": "created_at"
//                     },
//                     {
//                         "type": "string",
//                         "optional": false,
//                         "name": "io.debezium.time.ZonedTimestamp",
//                         "version": 1,
//                         "default": "1970-01-01T00:00:00Z",
//                         "field": "updated_at"
//                     },
//                     {
//                         "type": "string",
//                         "optional": true,
//                         "name": "io.debezium.time.ZonedTimestamp",
//                         "version": 1,
//                         "field": "deleted_at"
//                     },
//                     {
//                         "type": "int64",
//                         "optional": true,
//                         "default": 0,
//                         "field": "members"
//                     },
//                     {
//                         "type": "int64",
//                         "optional": true,
//                         "default": 0,
//                         "field": "followers"
//                     }
//                 ],
//                 "optional": true,
//                 "name": "bmarketplace.profile.inventory.GetValue",
//                 "field": "after"
//             },
//             {
//                 "type": "struct",
//                 "fields": [
//                     {
//                         "type": "string",
//                         "optional": false,
//                         "field": "version"
//                     },
//                     {
//                         "type": "string",
//                         "optional": false,
//                         "field": "connector"
//                     },
//                     {
//                         "type": "string",
//                         "optional": false,
//                         "field": "name"
//                     },
//                     {
//                         "type": "int64",
//                         "optional": false,
//                         "field": "ts_ms"
//                     },
//                     {
//                         "type": "string",
//                         "optional": true,
//                         "name": "io.debezium.data.Enum",
//                         "version": 1,
//                         "parameters": {
//                             "allowed": "true,last,false,incremental"
//                         },
//                         "default": "false",
//                         "field": "snapshot"
//                     },
//                     {
//                         "type": "string",
//                         "optional": false,
//                         "field": "db"
//                     },
//                     {
//                         "type": "string",
//                         "optional": true,
//                         "field": "sequence"
//                     },
//                     {
//                         "type": "string",
//                         "optional": true,
//                         "field": "table"
//                     },
//                     {
//                         "type": "int64",
//                         "optional": false,
//                         "field": "server_id"
//                     },
//                     {
//                         "type": "string",
//                         "optional": true,
//                         "field": "gtid"
//                     },
//                     {
//                         "type": "string",
//                         "optional": false,
//                         "field": "file"
//                     },
//                     {
//                         "type": "int64",
//                         "optional": false,
//                         "field": "pos"
//                     },
//                     {
//                         "type": "int32",
//                         "optional": false,
//                         "field": "row"
//                     },
//                     {
//                         "type": "int64",
//                         "optional": true,
//                         "field": "thread"
//                     },
//                     {
//                         "type": "string",
//                         "optional": true,
//                         "field": "query"
//                     }
//                 ],
//                 "optional": false,
//                 "name": "io.debezium.connector.mysql.Source",
//                 "field": "source"
//             },
//             {
//                 "type": "string",
//                 "optional": false,
//                 "field": "op"
//             },
//             {
//                 "type": "int64",
//                 "optional": true,
//                 "field": "ts_ms"
//             },
//             {
//                 "type": "struct",
//                 "fields": [
//                     {
//                         "type": "string",
//                         "optional": false,
//                         "field": "id"
//                     },
//                     {
//                         "type": "int64",
//                         "optional": false,
//                         "field": "total_order"
//                     },
//                     {
//                         "type": "int64",
//                         "optional": false,
//                         "field": "data_collection_order"
//                     }
//                 ],
//                 "optional": true,
//                 "field": "transaction"
//             }
//         ],
//         "optional": false,
//         "name": "bmarketplace.profile.inventory.Envelope"
//     },
//     "payload": {
//         "before": null,
//         "after": {
//             "id": 1,
//             "user_id": "353qh8qorefy",
//             "name": "I2",
//             "url": "https://example.com/247.png",
//             "media_channel_id": 1,
//             "min_views": 0,
//             "max_views": 0,
//             "min_interactions": 0,
//             "max_interactions": 0,
//             "created_at": "2022-06-23T18:47:11Z",
//             "updated_at": "2022-06-23T18:47:11Z",
//             "deleted_at": null,
//             "members": 0,
//             "followers": 0
//         },
//         "source": {
//             "version": "1.8.0.Final",
//             "connector": "mysql",
//             "name": "bmarketplace",
//             "ts_ms": 1656010257901,
//             "snapshot": "true",
//             "db": "profile",
//             "sequence": null,
//             "table": "inventory",
//             "server_id": 0,
//             "gtid": null,
//             "file": "binlog.000002",
//             "pos": 91166,
//             "row": 0,
//             "thread": null,
//             "query": null
//         },
//         "op": "r",
//         "ts_ms": 1656010257902,
//         "transaction": null
//     }
// }`
// 	type InventoryEvent struct {
// 		Event[Inventory]
// 	}
// 	var e InventoryEvent
// 	err := json.Unmarshal([]byte(jsonData), &e)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, e.Payload)
// 	assert.Equal(t, e.Payload.After.ID, int64(1))
// }
