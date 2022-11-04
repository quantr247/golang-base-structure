package dto

type RegisterUserRequestDTO struct {
	PhoneNumber string
	Password    string
	UserName    string
}

type RegisterUserResponseDTO struct {
	Token        string
	RefreshToken string
	ExpireIn     int64
}

type UserResponseDTO struct {
	ID          string
	UserName    string
	PhoneNumber string
}
