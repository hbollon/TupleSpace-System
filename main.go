package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hbollon/go-tuplespace"
)

const expireTimer = 100000

func printChoice() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nQue voulez-vous faire ?")
	fmt.Println("1 - Faire entrez une personne dans une salle")
	fmt.Println("2 - Faire sortir une personne d'une salle")
	fmt.Println("3 - Creer un badge")
	fmt.Println("4 - Désactiver un badge")
	fmt.Println("5 - Avoir la liste des personnes d'une salle")

	var validInput bool
	for !validInput {
		inputText, _ := reader.ReadString('\n')
		inputText = strings.Replace(inputText, "\n", "", -1)
		switch inputText {
		case "1":
			spacePersonnes.addPerson()
			validInput = true
		case "2":
			spacePersonnes.removePerson()
			validInput = true
		case "3":
			spacePersonnes.addbadge()
			validInput = true
		case "4":
			spacePersonnes.desactiverBadge()
			validInput = true
		case "5":
			listPersonne := spaceBatiment.getAllPersonneInBatiment()
			for _, personne := range listPersonne {
				fmt.Printf("%s %s\n", personne.nom, personne.prenom)
			}
			validInput = true
		default:
			fmt.Println("Input invalide! Refaites votre choix")
		}
	}

}

func main() {
	initSpaces()
	spacePersonnes.Write(tuplespace.New(expireTimer, listePersonne))
	spaceBatiment.Write(tuplespace.New(expireTimer, listeBatiments))
	listeBatiments[0].personnes.TupleSpace.Write(tuplespace.New(expireTimer, listePersonne[0]))
	listeBatiments[0].personnes.TupleSpace.Write(tuplespace.New(expireTimer, listePersonne[1]))

	for {
		printChoice()
	}
}
