package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("block", func() {

	DefaultMedia(Block)
	BasePath("")

	Action("list", func() {
		Routing(
			GET("/blocks"),
		)
		Description("Retrieve all blocks.")
		Response(OK, CollectionOf(Block))
	})

	Action("show", func() {
		Routing(
			GET("/blocks/:blockHASH"),
		)
		Description("Retrieve block with given hash.")
		Params(func() {
			Param("blockHASH", String, "Block hash")
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("mineBlock", func() {
		Routing(
			POST("/mineBlock"),
		)
		Description("Create new block.")
		Payload(func() {
			Member("data")
			Required("data")
		})
		Response(Created, "/blocks/[a-z;0-9]*")
		Response(BadRequest, ErrorMedia)
	})

})

var _ = Resource("peer", func() {

	Action("list", func() {
		Routing(
			GET("/peers"),
		)
		Description("Retrieve all peers.")
		Response(OK, CollectionOf(Peer))
	})

	Action("show", func() {
		Routing(
			GET("/peers/:peerHASH"),
		)
		Description("Retrieve peer with given hash.")
		Params(func() {
			Param("peerHASH", String, "Peer hash")
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("addPeer", func() {
		Routing(
			POST("/addPeer"),
		)
		Description("Create new peer.")
		Payload(func() {
			Member("host")
			Member("port")
			Required("host")
			Required("port")
		})
		Response(Created, "/peers/[a-z;0-9]*")
		Response(BadRequest, ErrorMedia)
	})

})

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/swagger.json", "public/swagger/swagger.json")
})
