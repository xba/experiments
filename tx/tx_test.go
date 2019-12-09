package tx

import (
	"testing"
)

func TestSupported(t *testing.T) {
	if !Supported() {
		t.Fatal("RTM/TSX not supported on this CPU")
	}
}

func TestTest(t *testing.T) {
	if Test() != uint8(0) {
		t.Fatal("Test called outside of transaction != 0")
	}
}

func TestTx(t *testing.T) {
	Begin()
	n := Test()
	End()
	if n != uint8(1) {
		t.Fatal("Test called within a transaction != 1")
	}
}

func TestTxAbort(t *testing.T) {
	n := 0
	if status := Begin(); status == Started {
		n = 1
		Abort()
		End()
		t.Fatal("Abort didn't stop the transaction")
	} else {
		if status&AbortExplicit != 1 {
			t.Fatal("Aborted transaction but status != AbortedExplicit")
		}
		if n != 0 {
			t.Fatal("Aborted transaction but mutation persisted")
		}
	}
}

func TestTxCommit(t *testing.T) {
	n := 0
	if status := Begin(); status == Started {
		n = 1
		End()
	} else {
		t.Fatal("Tx didn't Start/Commit")
	}
	if n != 1 {
		t.Fatal("Tx Committed but mutation didn't persist")
	}
}
