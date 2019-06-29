package Bitcoin

import "testing"

var BTCtest BtcAccount

const name = "abc"

func TestBtcAccount_ChangeIn(t *testing.T) {
	BTCtest.Cache.Hash = make(map[string]string)
	ok := BTCtest.ChangeIn(name, 1500)
	if ok {
		t.Log("ChangeIn test is successful")
	} else {
		t.Fatal("ChangeIn test is false")
	}
	BTCtest.ChangeOut(name, 1500)
}

func TestBtcAccount_ChangeOut(t *testing.T) {
	BTCtest.Cache.Hash = make(map[string]string)
	_ = BTCtest.ChangeIn(name, 1500)
	ok := BTCtest.ChangeOut(name, 1500)
	if ok {
		t.Log("ChangeIn test is successful")
	} else {
		t.Fatal("ChangeIn test is false")
	}
}

func TestBtcAccount_Check(t *testing.T) {
	BTCtest.Cache.Hash = make(map[string]string)
	_ = BTCtest.ChangeIn(name, 1500)
	account := BTCtest.Check()
	if account == 1500 {
		t.Log("Check test is successful")
	} else {
		t.Fatal("Check test is false")
	}
	BTCtest.ChangeOut(name, 1500)
}

func TestBtcAccount_InquiryTransfer(t *testing.T) {
	BTCtest.Cache.Hash = make(map[string]string)
	_ = BTCtest.ChangeIn(name, 1500)
	_, get := BTCtest.InquiryTransfer(name)
	if get {
		t.Log("InquiryTransfer test is successful")
	} else {
		t.Fatal("InquiryTransfer test is false")
	}
	BTCtest.ChangeOut(name, 1500)

}
