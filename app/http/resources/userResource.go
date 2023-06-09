package resources

type UserResource struct {
	ID       uint   `json:"id"`
	FullName string `json:"name"`
	Email    string `json:"email"`
}
