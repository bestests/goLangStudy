package accounts

// Account is type of account
/*
	첫글자를 대문자로 해서 export 했지만
	내부에 owner, balance에는 외무에서 접근을 못함
	Owner, Balance로 첫글자를 대문자로 하면 접근은 할수있지만(=public)
	함부로 접근 못하게 하고 싶을때(=private)
	go의 생산자를 사용
*/
type Account struct {
	owner   string
	balance int
}

// NewAccount create new account
/*
	생산자이므로 외부에서 접근할수있도록 첫글자 대문자(=public)
	생성해서 주소를 return 해줌
	이 function이 go의 기본적인 생성자 방식임
*/
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}
