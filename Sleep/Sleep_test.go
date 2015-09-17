package main

import ("testing"; "time")

func TestSleep(t *testing.T){

	startTime := time.Now().Second()

	sleep(time.Second* 5)

    endTime := (time.Now().Second())

    if (endTime-startTime) != 5 { 
    	t.Error("Expected time sleep is 5 second, got ", (endTime-startTime)) 
    	}
	}
