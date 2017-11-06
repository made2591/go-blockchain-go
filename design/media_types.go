package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Block is the block resource media type.
var Block = MediaType("application/ge.aviation.block+json", func() {
	Description("A tenant block")
	Attributes(func() {
		Attribute("index", Integer, "ID of block", func() {
			Example(1)
		})
		Attribute("previousHash", String, "Link to previous block", func() {
			Example("/blocks/1")
		})
		Attribute("timestamp", DateTime, "Timestamp of block")
		Attribute("data", Any, "Data in the block")
		Attribute("hash", String, "Hash of the block")

		Required("index", "previousHash", "timestamp", "data", "hash")
	})

	View("default", func() {
		Attribute("index")
		Attribute("previousHash")
		Attribute("timestamp")
		Attribute("data")
		Attribute("hash")
	})

	View("tiny", func() {
		Description("tiny is the view used to list blocks")
		Attribute("index")
		Attribute("previousHash")
		Attribute("data")
	})

	View("link", func() {
		Attribute("index")
		Attribute("previousHash")
	})
})

// Peer is the peer resource media type.
var Peer = MediaType("application/ge.aviation.peer+json", func() {
	Description("A tenant peer")
	Attributes(func() {
		Attribute("host", String, "Peer host endpoint")
		Attribute("port", Integer, "Peer port endpoint")

		Required("host", "port")
	})

	View("default", func() {
		Attribute("host")
		Attribute("port")
	})

	View("tiny", func() {
		Description("tiny is the view used to list peers")
		Attribute("host")
		Attribute("port")
	})

	View("link", func() {
		Attribute("host")
		Attribute("port")
	})
})