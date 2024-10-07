package workspaces

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWSLManager_Fetch(t *testing.T) {
	manager, err := NewWSLManagerWorkspace()
	assert.NoError(t, err)

	for _, dist := range manager.Distributions {
		fmt.Println(dist)
	}

	if json, err := json.MarshalIndent(manager.Distributions, "", "\t"); err != nil {
		assert.NoError(t, err)
	} else {
		fmt.Printf("%s", string(json))
	}

	assert.Greater(t, len(manager.Distributions), 0)
}

func TestWSLManager_GetOnlineDistributions(t * testing.T) {
	workspace, err := NewWSLManagerWorkspace()
	assert.NoError(t, err)

	dist, err := workspace.GetOnlineDistributions()
	assert.NoError(t, err)
	fmt.Println(dist)
}