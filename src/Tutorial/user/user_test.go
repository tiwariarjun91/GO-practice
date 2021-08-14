package user_test

import(
	"testing"
	gomock "github.com/golang/mock/gomock"
	"Tutorial/mocks"
	"Tutorial/user"
	
)
// function name should be starting with Test !!!! 
func TestUse(t *testing.T){
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDoer := Mocks.NewMockDoer(mockCtrl)
    testUser := &user.User{Doer:mockDoer}

	// Expect Do to be called once with 123 and "Hello GoMock" as parameters, and return nil from the mocked call.
    mockDoer.EXPECT().DoSomething(123, 25).Return(147).Times(1)

    testUser.Use()
}