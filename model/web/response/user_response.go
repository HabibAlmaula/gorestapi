package response

type UserResponse struct {
	Id       string `json:"id"`
	Fullname string `json:"full_name"`
	Email    string `json:"email"`
}
