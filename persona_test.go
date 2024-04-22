package persona

import (
	"testing"
)

func TestRandomAddress(t *testing.T) {
	address := randomAddress()
	if len(address) != 42 {
		t.Errorf("Expected address to be 42 characters long, got %s", address)
	}
}

func TestNew(t *testing.T) {
	p := New("0x4D4dCA590b0929cEe04Bbea60420aFd21A723799")
	if p.Name != "Ryoichi Takasaki" {
		t.Errorf("Expected name to be Ryoichi Takasaki, got %s", p.Name)
	}
	if p.Zodiac != "scorpio" {
		t.Errorf("Expected zodiac to be scorpio, got %s", p.Zodiac)
	}
	if p.Sex != "male" {
		t.Errorf("Expected sex to be male, got %s", p.Sex)
	}
}

func TestNewRandom(t *testing.T) {
	p := New("")
	if p.Name == "" {
		t.Errorf("Expected name to not be empty")
	}
	if p.Zodiac == "" {
		t.Errorf("Expected zodiac to not be empty")
	}
	if p.Sex == "" {
		t.Errorf("Expected sex to not be empty")
	}
}
