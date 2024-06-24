package api

func (a *API) InitRoutes() (*API, error) {
	a.router.Get("/posts", a.PostGetAllHandler)
	a.router.Get("/posts/{id}", a.PostGetByIdHandler)
	a.router.Post("/posts", a.PostCreateHandler)
	a.router.Put("/posts", a.PostUpdateHandler)
	a.router.Delete("/posts/{id}", a.PostDeleteHandler)

	return a, nil
}
