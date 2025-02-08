package internal

import (
	"reflect"
	"testing"
)

func TestSortPostsByDate(t *testing.T) {
	t.Run("sort posts by date", func(t *testing.T) {
		testData := []PostMetadata{
			{
				ID:      "1",
				Title:   "Test1",
				Date:    "2022-01-01",
				Summary: "",
				Tags:    nil,
			},
			{
				ID:      "2",
				Title:   "Test2",
				Date:    "2025-02-01",
				Summary: "",
				Tags:    nil,
			},
			{
				ID:      "3",
				Title:   "Test3",
				Date:    "2028-01-01",
				Summary: "",
				Tags:    nil,
			},
		}
		want := []PostMetadata{
			{
				ID:      "3",
				Title:   "Test3",
				Date:    "2028-01-01",
				Summary: "",
				Tags:    nil,
			},
			{
				ID:      "2",
				Title:   "Test2",
				Date:    "2025-02-01",
				Summary: "",
				Tags:    nil,
			},
			{
				ID:      "1",
				Title:   "Test1",
				Date:    "2022-01-01",
				Summary: "",
				Tags:    nil,
			},
		}
		SortPostsByDate(testData)

		if !reflect.DeepEqual(testData, want) {
			t.Errorf("SortPostsByDate returned %v, want %v", testData, want)
		}
	})
}
