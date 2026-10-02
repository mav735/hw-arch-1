package server

import (
	"os"
	"testing"
)

func TestAddressDefaultsTo8080(t *testing.T) {
	t.Setenv("HTTP_ADDR", "")

	if got := Address(os.Getenv); got != ":8080" {
		t.Fatalf("Address() = %q, want %q", got, ":8080")
	}
}

func TestAddressUsesEnvironment(t *testing.T) {
	t.Setenv("HTTP_ADDR", "127.0.0.1:9090")

	if got := Address(os.Getenv); got != "127.0.0.1:9090" {
		t.Fatalf("Address() = %q, want %q", got, "127.0.0.1:9090")
	}
}
