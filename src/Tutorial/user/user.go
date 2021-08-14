package user
import(
	"Tutorial/doer"
)

type User struct{
	Doer doer.Doer
}

func (u *User) Use() int{

	return u.Doer.DoSomething(123, 25)
}