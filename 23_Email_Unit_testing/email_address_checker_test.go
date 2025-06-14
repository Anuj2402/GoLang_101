package linkedinlearning

import(
	// "strings"
	"testing"
)

func TestIsLinkedInEmployee(t *testing.T) {
	// this for testing a single test case 
	/*
	if IsLinkedInEmployee("michael@google.com") == true {
		t.Errorf("expected false but got true")
	}
		*/

		//Now this is for testing the multiple test cases at once 
		testCases := []struct {
			inputEmail     string
			expectedOutput bool
		}{
			{
				inputEmail:     "anuj@linkedin.com",
				expectedOutput: true,
			},
			{
				inputEmail:     "abhi@google.com",
				expectedOutput: false,
			},
			{
				inputEmail:     "alex@microsoft.com",
				expectedOutput: false,
			},
		}
	
		for _, testCase := range testCases {
			t.Run(testCase.inputEmail, func(t *testing.T) {
				actualOutput := IsLinkedInEmployee(testCase.inputEmail)
				if actualOutput != testCase.expectedOutput {
					t.Errorf("expected IsLinkedInEmployee('%s') to be %t but got %t",
						testCase.inputEmail, testCase.expectedOutput, actualOutput)
				}
			})
		}
	}