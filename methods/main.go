package main

import "fmt"

func main() {
	fmt.Println("welcome to method code")
	vipin :=User{"vipin",21,"vipin@gmail.com"};
	fmt.Printf("%+v\n",vipin);
	fmt.Println(vipin.GetEmail())
	vipin.SetEmail("kumawatvipin@gmail.com")
	fmt.Println(vipin.GetEmail())

}

type User struct{
	Name string
	Age int
	Email string
}

func (u User) GetEmail() string{
	return u.Email;
}

func (u User) SetEmail(email string){
	u.Email = email;
	fmt.Printf("user %v email is %v\n",u.Name,u.Email);
}