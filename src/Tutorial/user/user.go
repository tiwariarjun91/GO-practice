package user

type User Struct{
	Doer doer.Doer
}

func (u *User) use() error{

	return u.Doer.DoSomething(123, "Arjun")
}