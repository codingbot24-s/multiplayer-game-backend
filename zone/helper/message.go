package zoneHelper

import "encoding/json"

type Message struct {
	Type string
	Data json.RawMessage `json:"data"` 
}

type MoveReq struct {
	X int
	Y int
}

/*
	send the json websocket message like this
	{
		"type": "something",
		"data " : {

		}
	}
*/
