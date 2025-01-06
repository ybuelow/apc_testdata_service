package main

import (
	"testing"
)

func Test_TraversWorks(t *testing.T) {
	t.Run("should return 0,0,0 when the parameter is 0", func(t *testing.T) {
		if a,b,c := travers(0); a != 0 && b != 0 && c != 0 {
			t.Errorf("Error should be 0 in case the parameter is 0")
		}
	})
	t.Run(" it should split the numbers even and correnctly case %2", func(t *testing.T) {
		if q,w,e := travers(22); q != 11 && w != 11 && e != 0{
			t.Errorf("Tne number should habe been split in 11, 11 ,0")
		}
	})
	t.Run("it should test the case %3 correctly", func(t *testing.T) {
		if a,b,c := travers(9); a != 3 && b != 3 && c != 3 {
			t.Errorf("Error, return should be 3,3,3 in case the parameter is dividable by 3")
		}
	})
	t.Run("it should test default case , that the number is not dividable by 3 and 2", func(t *testing.T) {
		if a,b,c := travers(157); a != 79 && b != 79 && c != 0 {
			t.Errorf("Error, return should be 3,3,3 in case the parameter is dividable by 3")
		}
		if a,b,c := travers(5557); a != 2779 && b != 2779 && c != 0 {
			t.Errorf("Error, return should be 3,3,3 in case the parameter is dividable by 3")
		}
	})
}

func TestRandRange(t *testing.T) {
	t.Run("it should test if it the function returns a randomin int value inbetween min and max", func(t *testing.T) {
		if got:= randRange(1, 30); got < 1 && got > 30{
			t.Errorf("Error is out of bounds")
		}
	})
	t.Run("it should test if it the function returns a randomin int value inbetween min and max", func(t *testing.T) {
		if got:= randRange(0, 56); got < 0 && got > 56{
			t.Errorf("Error is out of bounds")
		}
	})
}

func TestAllocatePassengers(t *testing.T) {
	t.Run("it should get the correct bounds for passerngers in and out on the first station", func(t *testing.T) {
		if a,b :=allocatePassangers(0, "start"); a > 30 && a < 0 && b != 0 {
			t.Errorf("There should have been at least 1 person who got in... The driver and not more that 30 also there should't be anyone who got out")
		}
	})
	t.Run("it should get the correct bounds for passerngers in and out for the last station", func(t *testing.T) {
		if a,b :=allocatePassangers(34, "end"); a != 0 && b != 34 {
			t.Errorf("There should't be anyone going in since its the last stop and all passengers left should go out")
		}
	})
	t.Run("it should get the correct bounds for passerngers in and out on a random station", func(t *testing.T) {
		if a,b :=allocatePassangers(14, "normal"); a > 30 && a < 0 && b < 0 && b > 14{
			t.Errorf(".....")
		}
	})
}
func TestInsertTelemetry(t *testing.T) {
	t.Run("it should test wether the func correctly changes the Telemetry", func(t *testing.T) {
		//got:= insertTelemetry(0, 2)
		//assert.Equal(got)
	})
}