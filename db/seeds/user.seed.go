package main

import (
	"database/sql"
	"layerg-poi-da/global"

	"go.uber.org/zap"
)

func seedUsers(cdb *sql.DB) {

	// insert nodes
	query := "INSERT INTO nodes (id, name, ip) VALUES ('f2a12890-3300-4150-aa04-b25dc3fab2d5','node1','14.241.247.139')"
	_, err := cdb.Exec(query)
	if err != nil {
		global.Logger.Error("error insert nodes", zap.Error(err))
		panic(err)
	}

	// insert users
	query = "INSERT INTO users (id, wallet_address, node_id) VALUES ('a1c3d943-d6b8-42ef-9c75-71e55c9be358','0x821dAb5C6fffD8183d4E3e4A5C1725c847c36789','f2a12890-3300-4150-aa04-b25dc3fab2d5')"
	_, err = cdb.Exec(query)
	if err != nil {
		global.Logger.Error("error insert users", zap.Error(err))
		panic(err)
	}

}
