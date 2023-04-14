package handler

// <-----------------

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/{id}", h.getFull)
	router.Post("/create", h.create)
	router.Delete("/{id}", h.delete)
	router.Put("/population/{id}", h.setPopulation)

	router.Get("/region/{region}", h.getFromRegion)
	router.Get("/district/{district}", h.getFromDistrict)
	router.Get("/population/range", h.getFromPopulation)
	router.Get("/foundation/range", h.getFromFoundation)

	return router
}
