package main

import (
	"github.com/goadesign/goa"
	"github.com/made2591/go-blockchain-go/app"
)

// PeerController implements the peer resource.
type PeerController struct {
	*goa.Controller
}

// NewPeerController creates a peer controller.
func NewPeerController(service *goa.Service) *PeerController {
	return &PeerController{Controller: service.NewController("PeerController")}
}

// AddPeer runs the addPeer action.
func (c *PeerController) AddPeer(ctx *app.AddPeerPeerContext) error {
	// PeerController_AddPeer: start_implement

	// Put your logic here

	// PeerController_AddPeer: end_implement
	return nil
}

// List runs the list action.
func (c *PeerController) List(ctx *app.ListPeerContext) error {
	// PeerController_List: start_implement

	// Put your logic here

	// PeerController_List: end_implement
	res := app.GeAviationPeerCollection{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *PeerController) Show(ctx *app.ShowPeerContext) error {
	// PeerController_Show: start_implement

	// Put your logic here

	// PeerController_Show: end_implement
	return nil
}
