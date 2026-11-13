package auth

import (
	"net/http"
	"testing"
)

func TestNoheader(t *testing.T) {
	data := http.Header{}
	_, err := GetAPIKey(data)
	if err == nil || err != ErrNoAuthHeaderIncluded {
		t.Fatalf("GetAPIKey error = %v, want %v", err, ErrNoAuthHeaderIncluded)
	}
}

func TestBearer(t *testing.T) {
	data := http.Header{}
	data.Add("Authorization", "Bearer jsklfjaklvnl")
	_, err := GetAPIKey(data)
	if err == nil || err.Error() != "malformed authorization header" {
		t.Fatalf("GetAPIKey error = %v, want malformed authorization header", err)
	}
}

func TestAuthMultiWord(t *testing.T) {
	data := http.Header{}
	data.Add("Authorization", "ApiKey jsklfjaklvnl klkll lslsls")
	_, err := GetAPIKey(data)
	if err == nil || err.Error() != "malformed authorization header" {
		t.Fatalf("GetAPIKey error = %v, want malformed authorization header", err)
	}
}

func TestValid(t *testing.T) {
	data := http.Header{}
	data.Add("Authorization", "ApiKey HelloWorld")
	key, err := GetAPIKey(data)
	if err != nil {
		t.Fatalf("GetAPIKey error = %v, want nil", err)
	}
	if key != "HelloWorld" {
		t.Fatalf("GetAPIKey key = %q, want %q", key, "HelloWorld")
	}
}