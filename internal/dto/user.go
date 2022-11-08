package dto

// swagger:model RegisterUserRequestDTO
type RegisterUserRequestDTO struct {
	PhoneNumber string
	Password    string
	UserName    string
}

// swagger:model RegisterUserResponseDTO
type RegisterUserResponseDTO struct {
	Token        string
	RefreshToken string
	ExpireIn     int64
}

// swagger:parameters getUserByID
type UserByIDRequestDTO struct {
	// in: string
	ID string
}

// swagger:model UserResponseDTO
type UserResponseDTO struct {
	ID          string
	UserName    string
	PhoneNumber string
}
