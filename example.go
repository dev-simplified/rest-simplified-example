package main

import (
	"fmt"

	rest "github.com/dev-simplified/rest-simplified"
)

func main() {
	/*
		In this example, I will be showing how a rest API can be accessed using rest-simplified.
		For this example, I have chosen a simple github API API that will be returning a json object with the repos in dev-simplified organization
		The API requires no auth, but I will try and demonstrate the usage of auth functions defined in the package
		I will also try to add an additional request header
		Finally this example will also include mock testing of the API considered
		General Usage Steps:
			Create auth using auth functions
			create API client
			Add additional request headers if needed
			execute the api using the API client
	*/
	executeAPI()

}

func executeAPI() {
	//Step1: Creation of authentication (not needed here, using just to demonstrate the functions)

	/*
		if the API necessitates usage of basic auth use CreateBasicAuth function of rest simplified.
		This function will take username and password as input and url encodes them in necessary format for rest API calls
	*/
	auth := rest.CreateBasicAuth("username", "password@123")
	fmt.Println("basic auth test : " + auth)

	/*
		if the API necessitates usage of bearer token use CreateBearerAuth function of rest simplified.
		This function will take bearer token as input and returns the token in necessary format for rest API calls
	*/
	auth = rest.CreateBearerAuth("bearertoken")
	fmt.Println("Bearer auth test : " + auth)

	//in our example the API doesnt need any auth. So we can set auth to "none" or ""
	auth = "none"

	//STEP2 Create API client
	/*
		Once auth creation is completed, we can create the API client.
		The CreateAPIClient function takes in api URL, api method, auth and content type as input and returns an api client
	*/
	apiClient := rest.CreateAPIClient("https://api.github.com/orgs/dev-simplified/repos", "GET", auth, "application/json")

	//STEP3 Add additional headers if necessary
	/*
		Request headers for Authorization and content-type gets created by default
		If any additional request header needs to be added, the AddAdditionalRequestHeader function of rest-simplified can be used
		This method allows addition of one header key value pair for each invocation
		This method is designed as a receiver function of apiClient and takes in headerName and headerValue as inputs and returns updated apiClient
		For our example, I am adding one additional request header for Accept
	*/
	apiClient = apiClient.AddAdditionalRequestHeader("Accept", "application/json")

	//STEP4 Execute the API client and retrieve the results
	/*
		Once the API client is created and all necessary configurations are done, the client can be executed
		To execute the client use the ExecuteAPI function of rest simplified
		ExecuteAPI takes in the request payload as input and returns the response code(int), response(string), err(error)
		Here to execute the API we do not need any payload, so setting payload as ""
	*/
	responseCode, response, err := apiClient.ExecuteAPI("")
	fmt.Println(responseCode)
	fmt.Println(response)
	fmt.Println(err)
}
