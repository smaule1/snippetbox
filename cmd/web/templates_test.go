package main

import (
	"testing"
	"time"

	"snippetbox.samuel/internal/assert"
)

func TestHumanDate(t *testing.T) {
	// Create a slice of anonymous structs containing the test case name,
	// input to our humanDate() function (the tm field), and expected output
	// (the want field).
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2024, 3, 17, 10, 15, 0, 0, time.UTC),
			want: "17 Mar 2024 at 10:15",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2024, 3, 17, 10, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "17 Mar 2024 at 09:15",
		},
	}

	// Loop over the test cases.
	for _, subtest := range tests {
		t.Run(subtest.name, func(t *testing.T) {
			hd := humanDate(subtest.tm)

			assert.Equal(t, hd, subtest.want)
		})
	}
}

// Alternative test table
// func TestExample(t *testing.T) {
//     t.Run("Example sub-test 1", func(t *testing.T) {
//         // Do a test.
//     })

//     t.Run("Example sub-test 2", func(t *testing.T) {
//         // Do another test.
//     })

//     t.Run("Example sub-test 3", func(t *testing.T) {
//         // And another...
//     })
// }
