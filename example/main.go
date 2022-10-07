package main

import (
	"context"
	"fmt"
	"os"

	"github.com/anvari1313/splitwise.go"
)

func main() {
	auth := splitwise.NewAPIKeyAuth(os.Getenv("API_KEY"))
	client := splitwise.NewClient(auth)

	userExamples(client)
	groupExamples(client)
	friendsExamples(client)
	expensesExamples(client)
	otherExamples(client)
}

func userExamples(client splitwise.Client) {
	currentUser, err := client.CurrentUser(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(currentUser)

	userByID, err := client.UserByID(context.Background(), currentUser.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println(userByID)

	updatedUser, err := client.UpdateUser(context.Background(), currentUser.ID, splitwise.UserFirstNameField("Ahmad"), splitwise.UserLastNameField("Anvari"))
	if err != nil {
		panic(err)
	}
	fmt.Println(updatedUser)
}

func groupExamples(client splitwise.Client) {
	groups, err := client.Groups(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(groups)

	group, err := client.GroupByID(context.Background(), 110)
	if err != nil {
		panic(err)
	}
	fmt.Println(group)
}

func friendsExamples(client splitwise.Client) {
	friends, err := client.Friends(context.Background())
	if err != nil {
		panic(err)
	}

	for _, friend := range friends {
		fmt.Println(friend.ID, friend.FirstName, friend.LastName, friend.Groups)
	}

	success, err := client.DeleteFriend(context.Background(), 123)
	if err != nil {
		panic(err)
	}

	fmt.Println("Delete fiend:", success)
}

func expensesExamples(client splitwise.Client) {
	expensesRes, err := client.Expenses(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", expensesRes)

	expenses, err := client.CreateExpenseByShare(context.Background(), splitwise.ExpenseByShare{
		Expense: splitwise.Expense{
			Cost:         "15000.00",
			Description:  "کافه امروز عصر",
			CurrencyCode: "IRR",
			GroupId:      0,
		},
		PaidUserID: 27163610,
		OwedUserID: 58839462,
		PaidShare:  "15000.00",
		OwedShare:  "15000.00",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", expenses)
}

func otherExamples(client splitwise.Client) {
	currencies, err := client.Currencies(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", currencies)
}
