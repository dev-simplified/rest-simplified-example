package main

import (
	"fmt"
	"testing"

	rest "github.com/dev-simplified/rest-simplified"
)

func TestExecuteAPI(t *testing.T) {
	/*
		One of the main challanges while writing code to access rest APIs is to mock test the API for unit testing
		rest-simplified is created with a built in mock framework
		To use the inbuilt mock framework, follow the below steps
			Enable mock in the package
			Create Mock response
			test the API function
	*/

	//STEP1 Enabling Mock
	/*
		To enable mock invoke EnableMock function of rest simplifed. This returns a string cinfirming mock enablement
	*/
	status := rest.EnableMock()
	fmt.Println(status)

	//STEP2 Create Mock response
	/*
		Once mock is enabled, a custom mock response, status code and error needs to be set to the package
		to create a mock response, use CreateMockResponse function of the package.
		This function takes in responseCode, response and err as input and returns a string confirming mock creation
		In case a mock response is not set, the package will return a default response
	*/

	status = rest.CreateMockResponse(200, `{"response": "github mock response"}`, nil)
	fmt.Println(status)

	//STEP3 execute the mocked API function
	executeAPI()
}
