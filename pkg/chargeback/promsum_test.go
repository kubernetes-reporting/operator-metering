package chargeback

import (
	"testing"
	"time"

	prom "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPromsumGetTimeRanges(t *testing.T) {
	janOne := time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC)
	tests := map[string]struct {
		startTime         time.Time
		endTime           time.Time
		chunkSize         time.Duration
		stepSize          time.Duration
		maxPromTimeRanges int64
		expectedRanges    []prom.Range
	}{
		"start and end are zero": {
			chunkSize:      time.Minute * 5,
			stepSize:       time.Minute,
			expectedRanges: nil,
		},
		"start and end are same": {
			startTime:      janOne,
			endTime:        janOne,
			chunkSize:      time.Minute * 5,
			stepSize:       time.Minute,
			expectedRanges: nil,
		},
		"period is exactly divisible by chunkSize": {
			startTime: janOne,
			endTime:   janOne.Add(2 * time.Hour),
			chunkSize: time.Hour,
			stepSize:  time.Minute,
			expectedRanges: []prom.Range{
				{
					Start: janOne,
					End:   janOne.Add(time.Hour),
					Step:  time.Minute,
				},
				// There is no second chunk, because it would be too small with
				// stepSize added
			},
		},
		"period is divisible by chunkSize with stepSize added": {
			startTime: janOne,
			endTime:   janOne.Add(2 * time.Hour).Add(time.Minute), // Add stepSize
			chunkSize: time.Hour,
			stepSize:  time.Minute,
			expectedRanges: []prom.Range{
				{
					Start: janOne,
					End:   janOne.Add(time.Hour),
					Step:  time.Minute,
				},
				{
					Start: janOne.Add(time.Hour + time.Minute),
					End:   janOne.Add(time.Hour + time.Minute).Add(time.Hour),
					Step:  time.Minute,
				},
			},
		},
	}

	for name, test := range tests {
		// Fix closure captures
		test := test
		t.Run(name, func(t *testing.T) {
			timeRanges, err := promsumGetTimeRanges(test.startTime, test.endTime, test.chunkSize, test.stepSize, test.maxPromTimeRanges)
			require.NoError(t, err)
			assert.Equal(t, timeRanges, test.expectedRanges)
		})
	}

}
