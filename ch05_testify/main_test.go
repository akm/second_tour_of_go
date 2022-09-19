package main

import "testing"

func TestPeopleAverageAge(t *testing.T) {
	t.Run("0 people", func(t *testing.T) {
		people := People{}
		if people.AverageAge() != 0 {
			t.Error("expected 0, got", people.AverageAge())
		}
	})

	t.Run("1 person", func(t *testing.T) {
		age := 25
		people := People{
			{
				FirstName: "Libbie",
				LastName:  "Drisko",
				Birthday:  "1998-06-15",
				Age:       age,
			},
		}
		actual := people.AverageAge()
		if actual != age {
			t.Errorf("expected %d, got %d", age, actual)
		}
	})

	t.Run("3 person", func(t *testing.T) {
		people := People{
			{Age: 10},
			{Age: 30},
			{Age: 50},
		}
		expect := 30
		actual := people.AverageAge()
		if actual != expect {
			t.Errorf("expected %d, got %d", expect, actual)
		}
	})
}
