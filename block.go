package main

import (
	"github.com/goadesign/goa"
	"github.com/made2591/go-blockchain-go/app"
)

var INITIAL_PEERS = []*app.GeAviationPeer{}
var BLOCKCHAIN = append(app.GeAviationBlockCollection{}, getGenesisBlock())

// BlockController implements the block resource.
type BlockController struct {
	*goa.Controller
}

// NewBlockController creates a block controller.
func NewBlockController(service *goa.Service) *BlockController {
	return &BlockController{Controller: service.NewController("BlockController")}
}

// List runs the list action.
func (c *BlockController) List(ctx *app.ListBlockContext) error {
	return ctx.OK(BLOCKCHAIN)
}

// MineBlock runs the mineBlock action.
func (c *BlockController) MineBlock(ctx *app.MineBlockBlockContext) error {
	// BlockController_MineBlock: start_implement

	// Put your logic here

	// BlockController_MineBlock: end_implement
	return nil
}

// Show runs the show action.
func (c *BlockController) Show(ctx *app.ShowBlockContext) error {
	// BlockController_Show: start_implement

	// Put your logic here

	// BlockController_Show: end_implement
	res := &app.GeAviationBlock{}
	return ctx.OK(res)
}
