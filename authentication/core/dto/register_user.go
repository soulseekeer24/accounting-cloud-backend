package dto

// RegisterUser object to be passed when wnat to register a user it contiains
// basic to the register process
type RegisterUser struct {
	Password  string `json:"password"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}
