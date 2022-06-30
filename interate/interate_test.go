package interate

import "testing"

func TestRepeat(t *testing.T){

	

	t.Run("Should word repeat 5 times", func(t *testing.T) {
		got, countRepeat := Repeat("2")
		expected := "22222"
		countRepeatExpected := 5
		if countRepeat != countRepeatExpected{
			t.Errorf("countRepeat %d countRepeatExpected %d", countRepeat, countRepeatExpected)
		}

		if got != expected{
			t.Errorf("got %q expected %q", got, expected)
		}
	})

	


	
}

func BenchmarkRepeat(b *testing.B){
	for i:=0; i< b.N;i++{
		Repeat("a")
	}
}