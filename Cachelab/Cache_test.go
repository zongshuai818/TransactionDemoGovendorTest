package Cachelib

import "testing"

var c Cache

const keyTemp, valueTemp = "aaa", "a"

func TestCache_Set(t *testing.T) {
	c.Hash = make(map[string]string)
	ok := c.Set(keyTemp, valueTemp)
	if ok {
		t.Log("Set test is successful")
	} else {
		t.Fatal("Set test is false")
	}
}

func TestCache_Get(t *testing.T) {
	c.Hash = make(map[string]string)
	c.Hash[keyTemp] = valueTemp
	_, getValue := c.Get(keyTemp)
	if getValue {
		t.Log("Get test is successful")
	} else {
		t.Fatal("Get test is false")
	}
}

func TestCache_Del(t *testing.T) {
	c.Hash = make(map[string]string)
	c.Hash[keyTemp] = valueTemp
	ok := c.Del(keyTemp)
	if ok {
		t.Log("Del test is successful")
	} else {
		t.Fatal("Del test is false")
	}
}
