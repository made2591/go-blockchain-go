package main

import (
	"github.com/goadesign/goa"
	"github.com/made2591/go-blockchain-go/app"
	"github.com/gorilla/websocket"
	"time"
)

var HTTP_PORT = 3001
var P2P_PORT = 6001
var INITIAL_PEERS = []
var BLOCKCHAIN = append(app.GeAviationBlockCollection{}, getGenesisBlock())

// BlockController implements the block resource.
type BlockController struct {
	*goa.Controller
}

func getGenesisBlock() (*app.GeAviationBlock) {
	return &app.GeAviationBlock{
		"My genesis block data",
		"816534932c2b7154836da6afc367695e6337db8a921823784c14378abed4f7d7",
		0,
		"0",
		time.Now(),
	}
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
