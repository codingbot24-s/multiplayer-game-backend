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

type Greeting struct {
	Message string
}

/*
	send the json websocket message like this
	{
		"type": "something",
		"data " : {

		}
	}
*/
