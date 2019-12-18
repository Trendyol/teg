package git

type options struct {
	UserName string
	Password string
}

type option func(*options)

func WithUserNameAndPassword(userName, password string) option {
	return func(args *options) {
		args.UserName = userName
		args.Password = password
	}
}

func WithAccessToken(val string) option {
	return func(args *options) {
		args.Password = val
	}
}
