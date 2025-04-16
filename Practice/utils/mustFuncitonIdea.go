package utils

// It is hard to check everywhere for errors in the code
// Hence Must utility function is a simple Idea to Seprate the error handling logic from the main code.
// NOTE: This should be use carefully as there can be scenarious where error need to handle in different way (NOT the common way done in the Must func)
func Must[T any](input T, err error) T {
	if err != nil {
		// Error handling logic
		panic(err)
	}
	return input
}
