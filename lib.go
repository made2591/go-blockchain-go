package main

import (

	"os"
	"time"
	"strconv"
	"strings"
	"crypto/sha256"
	"encoding/base64"

	log "github.com/sirupsen/logrus"

	"github.com/made2591/go-blockchain-go/app"

)

func init() {

	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)
	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)

}

func calculateHash(i int, p string, t time.Time, d string) (string){

	hasher := sha256.New()
	bytear := []byte(strconv.Itoa(i)+p+t.String()+d)
	hasher.Write(bytear)
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))

}

func generateNextBlock(d string) (*app.GeAviationBlock) {

	var pb = getLatestBlock()
	var ni = pb.Index + 1
	var nt = time.Now()
	var nh = calculateHash(ni, pb.Hash, nt, d)
	return &app.GeAviationBlock{d, nh, ni, pb.Hash, nt}

}

func getLatestBlock() (*app.GeAviationBlock) {

	return &app.GeAviationBlock{}

}

func getGenesisBlock() (*app.GeAviationBlock) {

	return &app.GeAviationBlock{"my genesis block!!", "816534932c2b7154836da6afc367695e6337db8a921823784c14378abed4f7d7", 0, "0", time.Now()}

}

func isValidNewBlock(nb *app.GeAviationBlock, pb *app.GeAviationBlock) (bool) {

	if pb.Index + 1 != nb.Index {

		log.Error("Invalid index")
		return false

	} else {

		if strings.Compare(pb.Hash, nb.PreviousHash) != 0 {

			log.Error("Invalid previous hash")
			return false

		} else {

			hnb := calculateHashForBlock(nb)

			if strings.Compare(hnb, nb.Hash) != 0 {

				log.Error("Invalid hash: " + hnb + " " + nb.Hash)
				return false

			}
		}
	}

	return true

}

func calculateHashForBlock(b *app.GeAviationBlock) (string) {

	return calculateHash(b.Index, b.PreviousHash, b.Timestamp, b.Data)

}

func replaceChain(newBlocks []*app.GeAviationBlock) (*app.GeAviationBlock) {

	if isValidChain(newBlocks) && len(newBlocks) > len(BLOCKCHAIN) {

		log.Info("Received blockchain is valid. Replacing current blockchain with received blockchain")

		BLOCKCHAIN := newBlocks

		broadcast(responseLatestMsg())

	} else {

		log.Info("Received blockchain invalid")

	}

}

func isValidChain(b []*app.GeAviationBlock) (bool) {

	return true

}