package accounts

import (
	"errors"
	"fmt"
)

// 에러를 변수로 선언해 처리함
// 에러는 변수명이 err로 시작해야함
var errNoMoney = errors.New("Not enough money")

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

// Deposit receiver 'Account'
/*
	func와 func명 사이에 (변수명 타입) 을 선언하여 receiver 선언
	receiver 시 변수명은 타입의 첫글자 소문자로 사용
	Account에서 이 func에 접근가능(=method)
	*Account 포인터 없을 시 새롭게 복사해서 씀
	현재 Account에 변경을 하고 싶을 땐(복사를 원치 않을땐 포인터 선언)
*/
func (a *Account) Deposit(amount int) {
	fmt.Println("a", a)
	fmt.Println("&a", &a)
	a.balance += amount
}

// Balance return current balance
/*
	값만 return 시엔 특별히 복사되도 상관없음으로
	포인터 사용 안해도 됨
*/
func (a Account) Balance() int {
	return a.balance
}

// Withdraw receiver 'Account' withdraw your account
/*
	go는 따로 error 없음
	개발자가 직접 에러 처리 해줘야함
	error 던저도 따로 에러로 떨어지거나
	찍히는거 없음.

	에러 시 찾으려는 금액과 error return
	성공 시 최종 balance와 error가 없으므로 nil return
*/
func (a *Account) Withdraw(amount int) (int, error) {
	if a.balance < amount {
		//return amount, errors.New("Not enough money")
		return amount, errNoMoney
	}
	a.balance -= amount

	return a.balance, nil
}
