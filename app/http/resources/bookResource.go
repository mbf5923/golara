package resources

type BookResource struct {
	ID          uint   `json:"id"`
	UserId      uint   `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       uint64 `json:"price"`
}
