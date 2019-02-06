package data_collector

import (
	"git01.bravofly.com/N7/heimdall/pkg/client/zone"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func Test_zoneCollection(t *testing.T) {
	aggregates, _ := GetZones(zone.MockZones{
		Path: filepath.Join("..", "..", "test", "cloudflare_zone.json"),
	})

	assert.Equal(t, len(aggregates), 18)
	assert.Equal(t, aggregates[0].ZoneName, "bravofly.at")
	assert.Equal(t, aggregates[12].ZoneName, "rumbo.com")
	assert.Equal(t, aggregates[0].ZoneID, "d746c5cf71899095e42c691788c3ccb9")
	assert.Equal(t, aggregates[12].ZoneID, "5d3e4831360561bdfe9f498142f6d032")
}
