package mlib

import "testing"

func TestOps(t *testing.T) {
	mm := NewMusicManager()
	if mm == nil {
		t.Error("NewMusicManager failed")
	}
	if mm.Len() != 0 {
		t.Error("Manager failed, not empty")
	}
	m0 := &MusicEntry{
		"1", "Nevada (Original Mix)", "Vicetone", "/Users/Dylan/Music/QQ音乐/Vicetone-Nevada (Original Mix)", "EDM",
	}
	mm.Add(m0)
}
