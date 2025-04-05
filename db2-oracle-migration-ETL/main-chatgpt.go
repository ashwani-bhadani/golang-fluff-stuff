├── cmd
│   └── migrate.go
├── config
│   └── config.go
├── db
│   ├── db2.go
│   └── oracle.go
├── internal
│   ├── processor.go
│   ├── types.go
│   ├── checkpoints.go
│   └── writer.go
├── logger
│   └── logger.go
├── main.go
├── Dockerfile
├── output
│   └── failed_records.csv
├── test
│   └── processor_test.go

// ---------------- main.go ----------------
package main

import (
	"db2oracle/cmd"
)

func main() {
	cmd.Execute()
}

// ---------------- cmd/migrate.go ----------------
package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"db2oracle/internal"
	"db2oracle/logger"
	"db2oracle/config"
)

var rootCmd = &cobra.Command{
	Use:   "db2oracle",
	Short: "Migrates data from DB2 to Oracle with transformation logic",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Init()
		cfg := config.Load()
		internal.StartMigration(cfg)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// ---------------- config/config.go ----------------
package config

type Config struct {
	DB2DSN        string
	OracleDSN     string
	BatchSize     int
	Workers       int
	CheckpointFile string
	FailedLogPath string
}

func Load() Config {
	return Config{
		DB2DSN:        "db2user:pass@tcp(db2host:50000)/dbname",
		OracleDSN:     "oracleuser:pass@oraclehost:1521/sid",
		BatchSize:     1000,
		Workers:       10,
		CheckpointFile: "checkpoint.txt",
		FailedLogPath:  "output/failed_records.csv",
	}
}

// ---------------- logger/logger.go ----------------
package logger

import (
	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

func Init() {
	logger, _ := zap.NewProduction()
	Log = logger.Sugar()
	defer logger.Sync()
}

// ---------------- internal/types.go ----------------
package internal

type Record struct {
	ID     int
	FieldA string
	FieldB string
}

// ---------------- internal/checkpoints.go ----------------
package internal

import (
	"os"
	"strconv"
	"io/ioutil"
)

func SaveCheckpoint(file string, offset int) {
	_ = ioutil.WriteFile(file, []byte(strconv.Itoa(offset)), 0644)
}

func LoadCheckpoint(file string) int {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return 0
	}
	offset, _ := strconv.Atoi(string(data))
	return offset
}

// ---------------- internal/writer.go ----------------
package internal

import (
	"encoding/csv"
	"os"
	"strconv"
)

func WriteFailedRecord(path string, r Record, err error) {
	f, _ := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	writer := csv.NewWriter(f)
	defer writer.Flush()
	writer.Write([]string{strconv.Itoa(r.ID), r.FieldA, r.FieldB, err.Error()})
}

// ---------------- internal/processor.go ----------------
package internal

import (
	"db2oracle/db"
	"db2oracle/config"
	"db2oracle/logger"
	"sync"
	"database/sql"
	"fmt"
)

func StartMigration(cfg config.Config) {
	db2, err := db.ConnectDB2(cfg)
	if err != nil {
		logger.Log.Fatal("DB2 connection error:", err)
	}
	oracle, err := db.ConnectOracle(cfg)
	if err != nil {
		logger.Log.Fatal("Oracle connection error:", err)
	}

	offset := LoadCheckpoint(cfg.CheckpointFile)
	records := make(chan Record, cfg.BatchSize)
	var wg sync.WaitGroup

	// Worker pool
	for i := 0; i < cfg.Workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for rec := range records {
				if shouldMigrate(rec) {
					err := writeToOracle(oracle, rec)
					if err != nil {
						WriteFailedRecord(cfg.FailedLogPath, rec, err)
						logger.Log.Errorf("Failed to write record ID %d: %v", rec.ID, err)
					}
				}
			}
		}()
	}

	for {
		batch, err := fetchFromDB2(db2, offset, cfg.BatchSize)
		if err != nil || len(batch) == 0 {
			break
		}
		for _, r := range batch {
			records <- r
		}
		offset += cfg.BatchSize
		SaveCheckpoint(cfg.CheckpointFile, offset)
	}

	close(records)
	wg.Wait()
	logger.Log.Info("Migration completed.")
}

func fetchFromDB2(db *sql.DB, offset, limit int) ([]Record, error) {
	rows, err := db.Query("SELECT ID, COL_A, COL_B FROM TABLE WHERE ... FETCH FIRST ? ROWS ONLY OFFSET ?", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []Record
	for rows.Next() {
		var r Record
		err := rows.Scan(&r.ID, &r.FieldA, &r.FieldB)
		if err == nil {
			records = append(records, r)
		}
	}
	return records, nil
}

func writeToOracle(db *sql.DB, r Record) error {
	_, err := db.Exec("INSERT INTO ORACLE_TABLE (ID, FIELD_A, FIELD_B) VALUES (:1, :2, :3)", r.ID, r.FieldA, r.FieldB)
	return err
}

func shouldMigrate(r Record) bool {
	return r.FieldA != ""
}

// ---------------- Dockerfile ----------------
FROM golang:1.21-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o db2oracle main.go
CMD ["./db2oracle"]

// ---------------- test/processor_test.go ----------------
package test

import (
	"db2oracle/internal"
	"testing"
)

func TestShouldMigrate(t *testing.T) {
	rec := internal.Record{FieldA: "valid"}
	if !internal.ShouldMigrate(rec) {
		t.Errorf("Expected true, got false")
	}
} 
