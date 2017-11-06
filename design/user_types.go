package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var BlockPayload = Type("BlockPayload", func() {
	Attribute("index", Integer)
	Attribute("previousHash", String)
	Attribute("timestamp", DateTime)
	Attribute("data", Any)
	Attribute("hash", String)
})