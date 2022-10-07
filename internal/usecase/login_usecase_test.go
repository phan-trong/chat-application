package usecase

import (
	"testing"
)

func Test_Login_UseCase(t *testing.T) {
	// ctl := gomock.NewController(t)
	// defer ctl.Finish()

	// mockUserRepo := domain.NewMockUserRepository(ctl)
	// uc := NewLoginUseCase(mockUserRepo)

	// email := "trongpq@beat.vn"

	// mockUserRepo.EXPECT().FindOneByEmail(gomock.Any(), email).Return(&domain.User{
	// 	ID:    uuid.New(),
	// 	Email: "trongpq@beat.vn",
	// }, nil)

	// response, err := uc.Handle(context.TODO(), &LoginRequest{
	// 	Email:    email,
	// 	Password: "123456789",
	// })

	// if err != nil {
	// 	t.Fatal()
	// }

	// fmt.Println(response)

	// if reflect.TypeOf(response).String() == "string" && len(response) == 0 {
	// 	t.Fatal()
	// }

	// fmt.Println("Success")
}
