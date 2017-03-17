package main

import "fmt"

var currentId int

var beers Beers

// Give us some seed data
func init() {
	RepoCreateBeer(Beer{Name: "Guiness" , Type: "Stout"})
	RepoCreateBeer(Beer{Name: "Sam Adams", Type: "Lager"})
}

func RepoFindBeer(id int) Beer {
	for _, t := range beers {
		if t.Id == id {
			return t
		}
	}
	// return empty Beer if not found
	return Beer{}
}

//this is bad, I don't think it passes race condtions
func RepoCreateBeer(t Beer) Beer {
	currentId += 1
	t.Id = currentId
	beers = append(beers, t)
	return t
}

func RepoDestroyBeer(id int) error {
	for i, t := range beers {
		if t.Id == id {
			beers = append(beers[:i], beers[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Beer with id of %d to delete", id)
}
