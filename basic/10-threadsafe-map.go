package main

import (
	"fmt"
	"sync"
	"time"
)

// ------------------------------- threadsafe map of countries and capitals
type capitalMap struct {
	mutex sync.Mutex
	capitalOf map[string]string
}

func (c *capitalMap) set(country, capital string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.capitalOf[country] = capital
}

func (c *capitalMap) get(country string) (string, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	capital, ok := c.capitalOf[country]
	if ok {
		return capital, nil
	}
	return "", fmt.Errorf("no known capital for %q", country)
}

// --------------------------------------------- adding and retrieving info
func fillMap(c *capitalMap) {
	c.set("Netherlands", "Amsterdam")
	c.set("Denmark", "Copenhagen")
	time.Sleep(2 * time.Millisecond)
	c.set("Sweden", "Stockholm")
	c.set("Norway", "Oslo")
	time.Sleep(2 * time.Millisecond)
	c.set("Finland", "Helsinki")
}

func readMap(c *capitalMap, done chan bool) {
	for {
		fmt.Printf("\n-----------------------\n");
		errorSeen := false
		for _, country := range []string{
			"Netherlands", "Denmark", "Sweden", "Norway", "Finland",
		} {
			capital, err := c.get(country)
			fmt.Printf("%q: ", country)
			if err != nil {
				fmt.Printf("%v\n", err)
				errorSeen = true
			} else {
				fmt.Printf("%q\n", capital)
			}
		}
		if errorSeen == false {
			done <- true
			return
		}
	}
}

func main() {
	doneChannel := make(chan bool)
	testMap := &capitalMap{
		capitalOf: map[string]string{},
	}

	go fillMap(testMap)
	go readMap(testMap, doneChannel)

	<- doneChannel
}
		

	
