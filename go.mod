module github.com/lxbc135/go_rest_api_example

go 1.17

require (
	github.com/gorilla/mux v1.8.0
)

require	github.com/lxbc135/go_rest_api_example/handlers v0.0.0-00010101000000-000000000000
require github.com/lxbc135/go_rest_api_example/models v0.0.0-00010101000000-000000000000

replace github.com/lxbc135/go_rest_api_example/handlers => ./handlers

replace github.com/lxbc135/go_rest_api_example/models => ./models
