// Copyright (c) 2013-2014 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

/*
Package btcdb provides a database interface for the Bitcoin block chain.

As of May 2013, there are over 235,000 blocks in the Bitcoin block chain and
and over 17 million transactions (which turns out to be over 11GB of data).
btcdb provides a database layer to store and retrieve this data in a fairly
simple and efficient manner.  The use of this should not require specific
knowledge of the database backend.

Basic Design

The basic design of btcdb is to provide two classes of items in a
database; blocks and transactions (tx) where the block number
increases monotonically.  Each transaction belongs to a single block
although a block can have a variable number of transactions.  Along
with these two items, several convenience functions for dealing with
the database are provided as well as functions to query specific items
that may be present in a block or tx.

Usage

At the highest level, the use of this packages just requires that you
import it, setup a database, insert some data into it, and optionally,
query the data back.  The first block inserted into the database will be
treated as the genesis block.  Every subsequent block insert requires the
referenced parent block to already exist.  In a more concrete example:

	// Import packages.
	import (
		"github.com/conformal/btcdb"
		_ "github.com/conformal/btcdb/ldb"
		"github.com/conformal/btcutil"
		"github.com/conformal/btcwire"
	)

	// Create a database and schedule it to be closed on exit.
	dbName := "example.db"
	db, err := btcdb.CreateDB("leveldb", dbName)
	if err != nil {
		// Log and handle the error
	}
	defer db.Close()


	// Insert the main network genesis block.
	genesis := btcutil.NewBlock(&btcwire.GenesisBlock)
	newHeight, err := db.InsertBlock(genesis)
	if err != nil {
		// Log and handle the error
	}
*/
package btcdb
