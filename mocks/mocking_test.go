package mocks

import (
	"bytes"
	"reflect"
	"testing"
)

type SpySleeper struct{
	Call int64
}

func (s *SpySleeper) Sleep(){
	s.Call++
}

type SpyCountDownOperations struct{
	Call []string
}

func (s *SpyCountDownOperations) Sleep(){
	s.Call = append(s.Call, sleep)
}

func (s *SpyCountDownOperations) Write(bytes []byte) (n int, err error){
	s.Call = append(s.Call, write)
	return
}
const write = "Write"
const sleep = "Sleep"

func TestCountDown(t *testing.T){
	// Show 3
	
	t.Run("Sleep in loop", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}
		CountDown(buffer, spySleeper)

		got := buffer.String()
		expected := `3
2
1
go!`

		if got != expected{
			t.Errorf("got %s expected %s", got, expected)
		}

		if spySleeper.Call != 3{
			t.Errorf("The count of sleep is different of 3 time")
		}
	})

	t.Run("Sleep before of loop", func(t *testing.T) {
		spyCountDownOperations := &SpyCountDownOperations{}
		CountDown(spyCountDownOperations, spyCountDownOperations)

		expected := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(expected, spyCountDownOperations.Call){
			t.Errorf("wanted calls %v got %v", expected, spyCountDownOperations.Call)
		}
	})





	
}