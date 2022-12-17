package response

import "github.com/kelompok4-loyaltypointagent/backend/models"

type FavoritesResponse struct {
	UserID  string           `json:"user_id"`
	Product *ProductResponse `json:"product"`
}

func NewFavoritesResponse(favorites models.Favorites) *FavoritesResponse {

	favoritesResponse := &FavoritesResponse{
		UserID: favorites.UserID.String(),
	}

	if favorites.Product != nil {
		favoritesResponse.Product = NewProductResponse(*favorites.Product)
	}

	return favoritesResponse

}

func NewFavoritesListResponse(favorites []models.Favorites) *[]FavoritesResponse {

	var favoritesResponse []FavoritesResponse

	for _, f := range favorites {
		if f.Product != nil {
			favoritesResponse = append(favoritesResponse, *NewFavoritesResponse(f))
		}
	}

	return &favoritesResponse

}
