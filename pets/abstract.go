package pets

import (
	"fmt"

	"go-breeders/models"
)

type AnimalInterface interface {
	Show() string
}

type DogFromFactory struct {
	Pet *models.Dog
}

func (dogFromFactory *DogFromFactory) Show() string {
	return fmt.Sprintf("This animal is a %s", dogFromFactory.Pet.Breed.Breed)
}

type CatFromFactory struct {
	Pet *models.Cat
}

func (catFromFactory *CatFromFactory) Show() string {
	return fmt.Sprintf("This animal is a %s", catFromFactory.Pet.Breed.Breed)
}

type PetFactoryInterface interface {
	newPet() AnimalInterface
}

type DogAbstractFactory struct{}

func (dogAbstractFactory *DogAbstractFactory) newPet() AnimalInterface {
	return &DogFromFactory{
		Pet: &models.Dog{},
	}
}

type CatAbstractFactory struct{}

func (catAbstractFactory *CatAbstractFactory) newPet() AnimalInterface {
	return &CatFromFactory{
		Pet: &models.Cat{},
	}
}

func NewPetFromAbstractFactory(species string) (AnimalInterface, error) {
	switch species {

	case "dog":
		var dogFactory DogAbstractFactory
		dog := dogFactory.newPet()
		return dog, nil

	case "cat":
		var catFactory CatAbstractFactory
		dog := catFactory.newPet()
		return dog, nil

	default:
		return nil, fmt.Errorf("unknown species: %s", species)
	}
}
