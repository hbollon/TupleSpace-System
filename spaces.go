package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/hbollon/go-tuplespace"
)

type TupleSpaceJournal struct {
	tuplespace.TupleSpace
}
type TupleSpacePersonnes struct {
	tuplespace.TupleSpace
}
type TupleSpaceService struct {
	tuplespace.TupleSpace
}
type TupleSpaceBatiment struct {
	tuplespace.TupleSpace
}

var (
	spaceJournal   TupleSpaceJournal
	spacePersonnes TupleSpacePersonnes
	spaceServices  TupleSpaceService
	spaceBatiment  TupleSpaceBatiment
)

func initSpaces() {
	spaceJournal.TupleSpace = tuplespace.NewSpace()
	spacePersonnes.TupleSpace = tuplespace.NewSpace()
	spaceServices.TupleSpace = tuplespace.NewSpace()
	spaceBatiment.TupleSpace = tuplespace.NewSpace()
}

func (space *TupleSpacePersonnes) addPerson() bool {
	recv1 := space.Read(tuplespace.New(0))
	tuple := <-recv1
	var studentList []Personne = tuple.Values()[0].([]Personne)
	for i, personne := range studentList {
		fmt.Printf("%d - %s %s\n", i+1, personne.nom, personne.prenom)
	}
	fmt.Println("\nQuelle personne voulez faire entrer dans un batiment ?")
	var validInput bool
	var inputPersonne int
	reader := bufio.NewReader(os.Stdin)
	for !validInput {
		inputText, _ := reader.ReadString('\n')
		inputText = strings.Replace(inputText, "\n", "", -1)
		inputPersonne, _ = strconv.Atoi(inputText)
		if inputPersonne <= 0 || inputPersonne > len(studentList) {
			fmt.Println("Personne invalide! Merci de refaire votre choix.")
		} else {
			validInput = true
		}
	}

	if res := spaceBatiment.isAlreadyInBatiment(studentList[inputPersonne-1]); !res {
		var inputBatiment int
		var batimentList = spaceBatiment.getAllBatiment()
		for i, batiment := range batimentList {
			fmt.Printf("%d - %s\n", i+1, batiment.nom)
		}
		fmt.Println("\nDans quel batiment ?")
		validInput = false
		for !validInput {
			inputText, _ := reader.ReadString('\n')
			inputText = strings.Replace(inputText, "\n", "", -1)
			inputBatiment, _ = strconv.Atoi(inputText)
			if inputBatiment <= 0 || inputBatiment > len(batimentList) {
				fmt.Println("Batiment invalide! Merci de refaire votre choix.")
			} else {
				validInput = true
			}
		}

		fmt.Printf("%s %s est entré dans le batiment %s.\n",
			studentList[inputPersonne-1].prenom,
			studentList[inputPersonne-1].nom,
			batimentList[inputBatiment-1].nom)
	} else {
		fmt.Printf("%s %s est déjà dans un batiment !\n",
			studentList[inputPersonne-1].prenom,
			studentList[inputPersonne-1].nom)
		return false
	}

	return true

}

func (space *TupleSpaceBatiment) getAllBatiment() []Batiment {
	recv1 := space.Read(tuplespace.New(0))
	tuple1 := <-recv1
	var batimentList []Batiment = tuple1.Values()[0].([]Batiment)
	return batimentList
}

func (space *TupleSpacePersonnes) removePerson() {
	recv1 := space.Read(tuplespace.New(0))
	tuple := <-recv1
	var studentList []Personne = tuple.Values()[0].([]Personne)
	for i, personne := range studentList {
		fmt.Printf("%d - %s %s\n", i+1, personne.nom, personne.prenom)
	}
}

func (space *TupleSpaceBatiment) isAlreadyInBatiment(p Personne) bool {
	recv1 := space.Read(tuplespace.New(0))
	tuple1 := <-recv1
	var batimentList []Batiment = tuple1.Values()[0].([]Batiment)
	for _, batiment := range batimentList {
		recv2 := batiment.personnes.TupleSpace.Read(tuplespace.New(0))
		for personne := range recv2 {
			if personne != nil {
				if p == personne.Values()[0].(Personne) {
					return true
				}
			} else {
				fmt.Printf("%v is nil\n", batiment.personnes)
				break
			}
		}
	}
	return false
}
