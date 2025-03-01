// BEGIN: 9c3d4f5g2hj6
package synchronizer

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestWorkerDecoratorLimitRetriesByTime_asyncRequestRollupInfoByBlockRange(t *testing.T) {
	// Create a new worker decorator with a minimum time between calls of 1 second
	workersMock := newWorkersMock(t)
	decorator := newWorkerDecoratorLimitRetriesByTime(workersMock, time.Second)

	// Create a block range to use for testing
	blockRange := blockRange{1, 10}

	// Test the case where there is no previous call to the block range
	ctx := context.Background()
	workersMock.On("asyncRequestRollupInfoByBlockRange", ctx, blockRange, noSleepTime).Return(nil, nil).Once()
	_, err := decorator.asyncRequestRollupInfoByBlockRange(ctx, blockRange, noSleepTime)
	assert.NoError(t, err)

	// Test the case where there is a previous call to the block range
	workersMock.On("asyncRequestRollupInfoByBlockRange", ctx, blockRange, mock.MatchedBy(func(sleep time.Duration) bool { return sleep > 0 })).Return(nil, nil).Once()
	_, err = decorator.asyncRequestRollupInfoByBlockRange(ctx, blockRange, noSleepTime)
	assert.NoError(t, err)
}
