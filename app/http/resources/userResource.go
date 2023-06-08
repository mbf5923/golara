package resources

type UserResource struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserResourceCollection struct {
	Data UserResource `json:"data"`
}
