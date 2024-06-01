package database

import (
	"database/sql"
	"fmt"
)

func NewDBConfig(username, databaseName, host, password string) dbConfig {
	return dbConfig{
		Username: username,
		DatabaseName: databaseName,
		Host: host,
		Password: password,
	} 
}

type PostgresDB	struct {
	dbClient *sql.DB
	db Idb
	dbConfig dbConfig
}

func (dbInstance *PostgresDB) InitDB(dbConfig dbConfig) (*PostgresDB, error) {

	if dbInstance.dbClient != nil {
		return dbInstance ,fmt.Errorf("DB already initialized")
	}

	dbInstance.dbConfig = dbConfig

	if dbInstance.dbConfig == nil {
		return nil, fmt.Errorf("DB Config not provided")
	}

	connStr := fmt.Sprintf("user=%s dbname = %s host = %s password = %s sslmode=disable", dbConfig.Username, dbConfig.DatabaseName, dbConfig.Host, dbConfig.Password)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	dbInstance.dbClient = db
	

	return dbInstance, nil
}


func (dbInstance *PostgresDB) GetDB() (*PostgresDB, error) {

	if dbInstance.dbClient == nil {
		dbInstance.InitDB(dbInstance.dbConfig)
	}

	return dbInstance, nil
	
}