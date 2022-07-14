package sort

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestInsertion(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			start := time.Now()
			InsertionSortInt(tc.sl)
			fmt.Println(time.Since(start))

			if !reflect.DeepEqual(tc.sl, tc.want) {
				t.Errorf("got %d, want %d", tc.sl, tc.want)
			}
		})
	}
}
