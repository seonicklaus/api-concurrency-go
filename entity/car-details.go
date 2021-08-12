package entity

type CarDetails struct {
	ID        int    `json:"id"`
	Brand     string `json:"brand"`
	Model     string `json:"model"`
	Year      int    `json:"model_year"`
	FirstName string `json:"owner_firstame"`
	LastName  string `json:"owner_lastname"`
	Email     string `json:"owner_email"`
}
