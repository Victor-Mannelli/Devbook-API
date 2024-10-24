package routes

import (
	"api/src/controllers"
	"net/http"
)

var followerRoutes = []Route{
	{
		URI:      "/followers/{userId}",
		Method:   http.MethodGet,
		Function: controllers.FindFollowers,
		Auth:     true,
	},
	{
		URI:      "/followers/{userId}/following",
		Method:   http.MethodGet,
		Function: controllers.FindFollowing,
		Auth:     true,
	},
	{
		URI:      "/followers/{userId}",
		Method:   http.MethodPost,
		Function: controllers.Follow,
		Auth:     true,
	},
	{
		URI:      "/followers/{userId}",
		Method:   http.MethodDelete,
		Function: controllers.UnFollow,
		Auth:     true,
	},
}
