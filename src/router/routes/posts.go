package routes

import (
	"api/src/controllers"
	"net/http"
)

var postsRoutes = []Route{
	{
		URI:      "/posts",
		Method:   http.MethodPost,
		Function: controllers.CreatePost,
		Auth:     true,
	},
	{
		URI:      "/posts",
		Method:   http.MethodGet,
		Function: controllers.FindPostsFromSelfAndFollowedUsers,
		Auth:     true,
	},
	{
		URI:      "/posts/user/{userId}",
		Method:   http.MethodGet,
		Function: controllers.FindPostsFromUser,
		Auth:     true,
	},
	{
		URI:      "/posts/{postId}",
		Method:   http.MethodGet,
		Function: controllers.FindPostById,
		Auth:     true,
	},
	{
		URI:      "/posts/{postId}",
		Method:   http.MethodPut,
		Function: controllers.UpdatePost,
		Auth:     true,
	},
	{
		URI:      "/posts/{postId}",
		Method:   http.MethodDelete,
		Function: controllers.DeletePost,
		Auth:     true,
	},
}
