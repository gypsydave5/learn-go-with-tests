package concurrency

import "testing"

func TestWebsiteChecker(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	expectedResults := []bool{
		true,
		true,
		false,
	}

	actualResults := WebsiteChecker(websites)

	want := len(websites)
	got := len(actualResults)
	if want != got {
		t.Fatalf("Wanted %v, got %v", want, got)
	}

	if !sameResults(expectedResults, actualResults) {
		t.Fatalf("Wanted %v, got %v", expectedResults, actualResults)
	}
}

func sameResults(as, bs []bool) bool {
	for index, a := range as {
		if a != bs[index] {
			return false
		}
	}
	return true
}
