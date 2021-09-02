package yookassa

import "testing"

func TestUUIDGenUniqueness (t *testing.T) {
	id1 := UUIDGen()
	id2 := UUIDGen()
	if id1 == id2{
		t.Fatalf("Got 2 same uuids.")
	}
}