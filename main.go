package main

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/db/memdb"
)

func main() {
	db := memdb.NewDB()
	rwdb := db.ReadWriter()
	key := []byte{1}
	rwdb.Set(key, []byte{10})
	fmt.Println(rwdb.Get(key))
}
