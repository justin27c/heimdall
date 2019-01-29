package data_collector

import (
	"fmt"
	"git01.bravofly.com/n7/heimdall/cmd/client/zone"
	"git01.bravofly.com/n7/heimdall/cmd/model"
)

func GetZones(zoneClient zone.ZonesClient) ([]*model.Aggregate, error) {
	zones, err := zoneClient.GetZonesId()
	if err != nil {
		return nil, fmt.Errorf("ERROR ZoneName from CF Client %v", zones)
	}

	result := make([]*model.Aggregate, 0)
	for _, zone := range zones {
		result = append(result, model.NewAggregate(zone))
	}

	return result, nil
}
