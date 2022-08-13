package main

func GetUserCredentials(authoriser Authoriser) *UserCredentials {
	return authoriser.getUserCredentials()
}
