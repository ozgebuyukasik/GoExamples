package main

import "testing"

func TestMigratory(t *testing.T) {
	t.Run("Most frequently sighted type: ", func(t *testing.T) {
		birds := []int32{1, 4, 4, 4, 5, 3}

		received := MigratoryBirds(birds)
		want := int32(4)

		if received != want {
			t.Errorf("You got %d but we want %d", received, want)
		}
	})
}
