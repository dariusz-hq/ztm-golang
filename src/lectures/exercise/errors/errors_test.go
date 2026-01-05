package timeparse

import "testing"

func TestParseSuccessful(t *testing.T) {
	res, err := Parse("11:22:33")
	if err != nil {
		t.Errorf("Expected success, got error %v", err)
	}
	expectedHours := 11
	if res.hours != expectedHours {
		t.Errorf("Expected %v, got %v", expectedHours, res.hours)
	}
	expectedMinutes := 22
	if res.minutes != expectedMinutes {
		t.Errorf("Expected %v, got %v", expectedMinutes, res.minutes)
	}
	expectedSeconds := 33
	if res.seconds != expectedSeconds {
		t.Errorf("Expected %v, got %v", expectedSeconds, res.seconds)
	}
}

func TestParseFailsIncorrectString(t *testing.T) {
	res, err := Parse("11:22")
	if err == nil {
		t.Errorf("Expected error, got %v", res)
	}
}

func TestParseFailsIncorrectComponent(t *testing.T) {
	testCases := []struct {
		timeString        string
		expectedHourError bool
		expectedMinError  bool
		expectedSecError  bool
	}{
		{"xx:12:12", true, false, false},
		{"12:xx:12", false, true, false},
		{"12:12:xx", false, false, true},
	}

	for _, v := range testCases {
		_, err := Parse(v.timeString)
		var errMsg string
		if v.expectedHourError {
			errMsg = "Cannot convert xx to hours"
		}
		if v.expectedMinError {
			errMsg = "Cannot convert xx to minutes"
		}
		if v.expectedSecError {
			errMsg = "Cannot convert xx to seconds"
		}
		if errMsg != "" && err == nil {
			t.Errorf("Wanted error message %v, got nothing", errMsg)
		}
	}
}
