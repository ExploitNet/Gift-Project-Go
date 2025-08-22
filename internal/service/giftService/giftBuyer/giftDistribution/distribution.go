package giftDistribution

import (
	"gift-buyer/internal/config"
	"sync/atomic"
)

type distributionImpl struct {
	accountCounters map[string]*atomic.Int32
}

func NewDistribution(allParams [][]*config.DistributionParams) *distributionImpl {
	accountCounters := make(map[string]*atomic.Int32, len(allParams))
	for _, giftParam := range allParams {
		for _, param := range giftParam {
			if existing, exists := accountCounters[param.Username]; exists {
				existing.Add(int32(param.Count))
			} else {
				accountCounters[param.Username] = &atomic.Int32{}
				accountCounters[param.Username].Store(int32(param.Count))
			}
		}
	}
	return &distributionImpl{accountCounters: accountCounters}
}

func (d *distributionImpl) CheckAvailableDistrubuit(username string) bool {
	count, exists := d.accountCounters[username]
	if !exists {
		return false
	}

	if count.Load() > 0 {
		count.Add(-1)
		return true
	}

	return false
}

func (d *distributionImpl) ReturnAllocation(username string) {
	if count, exists := d.accountCounters[username]; exists {
		count.Add(1)
	}
}
