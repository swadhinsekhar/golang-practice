package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseGetCommand(t *testing.T) {
    cmd := &CommandGet {
        Key:    []byte("foo"),
    }

    r := bytes.NewReader(cmd.Bytes())
    pcmd := ParseCommand(r)

    assert.Equal(t, cmd, pcmd, "CommandGet should equal")
}

func TestParseSetCommand(t *testing.T) {
    cmd := &CommandSet {
        Key:    []byte("foo"),
        Value:  []byte("bar"),
        TTL:    2,
    }

    r := bytes.NewReader(cmd.Bytes())
    pcmd := ParseCommand(r)

    assert.Equal(t, cmd, pcmd, "CommandSet should equal")
}





