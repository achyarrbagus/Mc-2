package main

import "fmt"

type Profile struct {
	Name   string
	Health int
	Power  int
	Exp    int
}

func makeProfile(yourmc string) Profile {
	if yourmc == "Sasuke" {
		return Profile{
			Name:   "Sasuke",
			Health: 200,
			Power:  100,
			Exp:    0,
		}
	} else if yourmc == "Goku" {
		return Profile{
			Name:   "Goku",
			Health: 400,
			Power:  300,
			Exp:    100,
		}
	} else if yourmc == "Naruto" {
		return Profile{
			Name:   "Naruto",
			Health: 150,
			Power:  200,
			Exp:    50,
		}
	} else {
		return Profile{}
	}
}

func powerUp(yourmc string, power int) Profile {
	if yourmc == "Sasuke" {
		return Profile{
			Name:   "Sasuke",
			Health: 200 + (200 * power),
			Power:  100 + (100 * power),
			Exp:    0 + (0 * power),
		}
	} else if yourmc == "Goku" {
		return Profile{
			Name:   "Goku",
			Health: 400 + (400 * power),
			Power:  300 + (300 * power),
			Exp:    100 + (100 * power),
		}
	} else if yourmc == "Naruto" {
		return Profile{
			Name:   "Naruto",
			Health: 150 + (150 * power),
			Power:  200 + (200 * power),
			Exp:    50 + (50 * power),
		}
	} else {
		return Profile{}
	}
}

func main() {
	myProfile := makeProfile("Sasuke")
	power := powerUp("Sasuke", 3)
	fmt.Println(myProfile.Name)
	fmt.Println(power.Health)

}
