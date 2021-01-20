package route

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rizkysaputra4/moviwiki/server/env"
	h "github.com/rizkysaputra4/moviwiki/server/route/handler"
	"github.com/rizkysaputra4/moviwiki/server/route/middleware"
)

// InitRoute initialize function
func InitRoute() {
	c := env.GetConfiguration()

	r := chi.NewRouter()

	r.Route("/auth", func(r chi.Router) {
		r.Get("/register-username", h.CheckIfUserNameExist) // Check if username already exist
		r.Get("/register-email", h.CheckIfEmailExist)       // Check if email already exist
		r.Post("/register", h.RegisteringNewUser)           // Register new user
		r.Get("/login-username", h.CheckIfUserExist)        // Check if user exist using username or email when login
		r.Get("/login-password", h.CheckIfPasswordMatch)    // Match the username or email with password
		r.Get("/logout", h.LogOut)
	})

	r.Mount("/member", MemberRoute())
	r.Mount("/admin", adminRouter())

	// r.Route("/public", func(r chi.Router) {
	// 	r.Route("/user", func(r chi.Router) {
	// 		r.Get("/id?{id}", GetUserProfile)
	// 		r.Get("/search?{keyword}", SearchUserProfile)
	// 	})

	// 	r.Route("/movie", func(r chi.Router) {
	// 		r.Get("/", GetRecentAddedMovie)
	// 		r.Get("/search?{search-params}", SearchMovie)

	// 		r.Route("/movie?{movie-id}", func(r chi.Router) {
	// 			r.Get("/", GetMovieById)
	// 			r.Get("/reviews", GetMovieReviews)
	// 			r.Get("/review-comment", GetReviewComment)
	// 			r.Get("/comment", GetMovieComment)

	// 			r.Route("/eps?{eps}", func(r chi.Router) {
	// 				r.Get("/", GetSeriesEpsInfo)
	// 				r.Get("/person", GetPersonInEps)
	// 				r.Get("/reviews", GetEpsReviews)
	// 				r.Get("review-comment", GetReviewComment)
	// 				r.Get("comment", GetEpsComment)
	// 			})
	// 		})
	// 	})

	// 	r.Route("/act", func(r chi.Router) {
	// 		r.Get("/", GetRecentAddedMoviePeson)
	// 		r.Get("/search?{act-keyword}", SearchMoviePerson)

	// 		r.Route("/act?{act-id}", func(r chi.Router) {
	// 			r.Get("/", GetMoviePersonById)
	// 			r.Get("/movie", PersonMovieList)
	// 		})
	// 	})
	// })

	fmt.Println("server running on port", c.ServerAPIPort)
	http.ListenAndServe(c.ServerAPIPort, r)
}

func adminRouter() http.Handler {

	r := chi.NewRouter()
	r.Use(middleware.RoleEnforcer)
	r.Use(middleware.UpdateJWTExp)
	r.Use(middleware.UpdateSessionExp)

	r.Post("/register-new-admin", h.RegisterNewAdmin) // register new admin
	r.Put("/admin-level", h.ChangeAdminLevel)         // Promote regular user to admin
	// r.Put("/update-admin", UpdateAdminLevel) // Update the admin level

	// r.Route("/movie", func(r chi.Router) {
	// 	r.Post("/new-movie", AddNewMovie)
	// 	r.Delete("/movie?{movie-id}", DeleteMovie)
	// 	r.Put("/movie?{movie-id}", AdminApproveUpdateMovie)
	// 	r.Post("/tag", AdminAddNewMovieTag)
	// 	r.Put("/tag?{tag-id}", AdminApproveEditTag)
	// 	r.Post("/genre", AdminAddNewGenre)
	// 	r.Put("/genre?{genre-id}", AdminApproveEditGenre)
	// })

	// r.Route("/act-crew", func(r chi.Router) {
	// 	r.Put("/act-crew?{act-crew-id}", AdminApproveEditActInfo)
	// 	r.Delete("/act-crew?{act-id}", AdminDeleteAct)
	// 	r.Post("/add-role", AdminApproveRole)
	// })

	// r.Route("/review", func(r chi.Router) {
	// 	r.Put("/?{review-id}", AdminEditReview)
	// 	r.Delete("/?{review-id", AdminDeleteReview)

	// 	r.Route("/review?{review-id}", func(r chi.Router) {
	// 		r.Put("/comment?{comment-id}", AdminEditCommentReview)
	// 		r.Delete("/comment?{comment-id}", AdminDeleteCommentReview)
	// 	})
	// })

	// r.Route("/comment", func(r chi.Router) {
	// 	r.Put("/comment?{comment-id}", AdminEditComment)
	// 	r.Delete("comment?{comment-id}", AdminDeleteComment)
	// })

	// r.Route("/user", func(r chi.Router) {
	// 	r.Get("/", GetListOfUsers)
	// 	r.Post("/new-user", RegisteringNewUser)
	// 	r.Delete("/user?{user-id}", AdminDeleteUserAccount)
	// 	r.Put("/punish?{user-id}", AdminPunishUser)

	// })

	return r
}

// MemberRoute ...
func MemberRoute() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RoleEnforcer)
	r.Use(middleware.UpdateJWTExp)

	// 	r.Put("/id?{user-id}", UpdateFullUserInfo) // Update user profile
	// 	r.Get("/id?{user-id}", GetUserCredential) // Get user credential
	// 	r.Delete("/id?{user-id}", DeleteUser)
	r.Put("/uid", h.UpdateUserShortInfo) // Update user login credentials
	// 	r.Post("/report?{user-id}", ReportUser)

	// 	r.Route("/movie", func(r chi.Router) {
	// 		r.Post("/add-movie", SuggestAddNewMovie)
	// 		r.Put("/edit?{movie-id}", SuggestUpdateMovie)
	// 		r.Put("/delete?{movie-id}", UserSuggestDeleteMovie)

	// 		r.Route("/id?{movie-id}", func(r chi.Router) {
	// 			r.Post("/review", UserPostReview)
	// 			r.Put("/review?{review-id}", UpdateReview)
	// 			r.Post("/report?{review-id}", ReportReview)

	// 			r.Post("/comment", UserPostComment)
	// 			r.Put("/like", UserLikeComment)
	// 			r.Delete("comment?{comment-id}", UserDeleteComment)
	// 			r.Post("/comment?{comment-id}", UserReportComment)

	// 			r.Route("/eps?{eps-id}", func(r chi.Router) {
	// 				r.Post("/add-eps", SuggestAddNewEps)
	// 				r.Put("/edit", SuggestUpdateEps)
	// 				r.Put("/delete", UserSuggestDeleteEps)

	// 				r.Post("/review", UserPostReviewEps)
	// 				r.Put("/review?{eps-review-id}", UpdateReviewEps)
	// 				r.Post("/report?{eps-review-id}", ReportReviewEps)

	// 				r.Post("/comment", UserPostCommentEps)
	// 				r.Put("/like", UserLikeCommentEps)
	// 				r.Delete("comment?{eps-comment-id}", UserDeleteCommentEps)
	// 				r.Post("/comment?{eps-comment-id}", UserReportCommentEps)
	// 			})
	// 		})
	// 	})

	return r
}
