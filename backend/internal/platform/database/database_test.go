package database

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"
)

func TestConnectInvalidURL(t *testing.T) {
	_, err := Connect(context.Background(), "://invalid")
	if err == nil {
		t.Fatal("Connect returned no error for an invalid URL")
	}
	if !strings.Contains(err.Error(), "creating the pool") {
		t.Errorf("error %q does not mention creating the pool", err)
	}
}

func TestConnectUnreachableDatabase(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := Connect(ctx, "postgres://user:password@127.0.0.1:1/database?sslmode=disable")
	if err == nil {
		t.Fatal("Connect returned no error for an unreachable database")
	}
	if !strings.Contains(err.Error(), "pinging the pool") {
		t.Errorf("error %q does not mention pinging the pool", err)
	}
}

func TestConnectDatabase(t *testing.T) {
	dbURL := os.Getenv("TEST_DATABASE_URL")
	if dbURL == "" {
		t.Skip("TEST_DATABASE_URL is not set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, err := Connect(ctx, dbURL)
	if err != nil {
		t.Fatalf("Connect returned an unexpected error: %v", err)
	}
	defer pool.Close()
}
