package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/hbollon/go-tuplespace"
	"github.com/sirupsen/logrus"
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

func (bat *Batiment) checkDoorTimer() bool {
	recv := bat.accessControl.Read(tuplespace.New(0, "door timer"))
	tuple := <-recv
	if tuple == nil {
		bat.accessControl.Write(tuplespace.New(doorTimerDur, "door timer"))
		logrus.Debug("creation")
		logrus.Info("Entrée autorisé !")
		return true
	} else if tuple.IsExpired() {
		tuple.Renew()
		logrus.Debug("renew")
		logrus.Info("Entrée autorisé !")
		return true
	}
	logrus.Error("Attention ! Entrée non autorisé !")
	return false
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
			logrus.Warnf("Personne invalide! Merci de refaire votre choix.")
		} else {
			validInput = true
		}
	}
	if res := spaceBatiment.isAlreadyInBatiment(studentList[inputPersonne-1]); !res {
		var inputBatiment int
		var batimentList = spaceBatiment.getAllBatiment()
		validInput = false
		for !validInput {
			for i, batiment := range batimentList {
				var str string
				recv := batiment.accessControl.Read(tuplespace.New(0, "door timer"))
				tuple := <-recv
				if tuple != nil && !tuple.IsExpired() {
					str = " - Porte encore ouverte, accès interdit!"
				}
				fmt.Printf("%d - %s%s\n", i+1, batiment.nom, str)
			}
			fmt.Println("\nDans quel batiment ?")
			inputText, _ := reader.ReadString('\n')
			inputText = strings.Replace(inputText, "\n", "", -1)
			inputBatiment, _ = strconv.Atoi(inputText)
			if inputBatiment <= 0 || inputBatiment > len(batimentList) {
				fmt.Println("Batiment invalide! Merci de refaire votre choix.")
			} else {
				validInput = true
			}
		}
		batimentChoice := batimentList[inputBatiment-1]
		if batimentChoice.personHaveAccess(studentList[inputPersonne-1]) {
			spaceBatiment.addPersonneInBatiment(studentList[inputPersonne-1], batimentChoice)
			speech.Speak(
				fmt.Sprintf("%s %s est entré dans le batiment %s.\n",
					studentList[inputPersonne-1].prenom,
					studentList[inputPersonne-1].nom,
					batimentChoice.nom),
			)
			if !batimentChoice.checkDoorTimer() {
				go func() {
					for i := 0; i < 3; i++ {
						speech.Speak("Attention alarme !")
					}
				}()
			} else {
				recv1 := space.Read(tuplespace.New(expireTimer, "Door Timer"))
				doorOpenService := <-recv1
				if doorOpenService != nil {
					doorOpenService.Renew()
				}
			}
		} else {
			fmt.Printf("%s %s n'a pas le droit d'entré dans le batiment %s.\n",
				studentList[inputPersonne-1].prenom,
				studentList[inputPersonne-1].nom,
				batimentChoice.nom)
		}
	} else {
		fmt.Printf("%s %s est déjà dans un batiment !\n",
			studentList[inputPersonne-1].prenom,
			studentList[inputPersonne-1].nom)
		return false
	}
	return true
}

func (space *TupleSpaceBatiment) addPersonneInBatiment(p Personne, b Batiment) bool {
	recv1 := space.Read(tuplespace.New(0))
	tuple1 := <-recv1
	var batimentList []Batiment = tuple1.Values()[0].([]Batiment)
	for _, batiment := range batimentList {
		if reflect.DeepEqual(batiment, b) {
			batiment.personnes.TupleSpace.Write(tuplespace.New(expireTimer, p))
			return true
		}
	}
	return false
}

func (space *TupleSpaceBatiment) getAllBatiment() []Batiment {
	recv1 := space.Read(tuplespace.New(0))
	tuple1 := <-recv1
	var batimentList []Batiment = tuple1.Values()[0].([]Batiment)
	return batimentList
}

func (space *TupleSpacePersonnes) removePerson() bool {
	recv1 := space.Read(tuplespace.New(0))
	tuple := <-recv1
	var studentList []Personne = tuple.Values()[0].([]Personne)
	for i, personne := range studentList {
		fmt.Printf("%d - %s %s\n", i+1, personne.nom, personne.prenom)
	}
	fmt.Println("\nQuelle personne voulez faire sortir des batiments ?")
	var validInput bool
	var inputPersonne int
	reader := bufio.NewReader(os.Stdin)
	for !validInput {
		inputText, _ := reader.ReadString('\n')
		inputText = strings.Replace(inputText, "\n", "", -1)
		inputPersonne, _ = strconv.Atoi(inputText)
		if inputPersonne <= 0 || inputPersonne > len(studentList) {
			logrus.Warnf("Personne invalide! Merci de refaire votre choix.")
		} else {
			validInput = true
		}
	}
	personne := studentList[inputPersonne-1]
	if res := spaceBatiment.isAlreadyInBatiment(personne); res {
		emptyBatiment := Batiment{}
		if batiment := spaceBatiment.findInBatiments(personne); batiment != emptyBatiment {
			recv2 := batiment.personnes.TupleSpace.Take(tuplespace.New(0, personne))
			if extractedPersonne := <-recv2; extractedPersonne != nil {
				fmt.Printf("%s %s est sorti du batiment avec succès !\n",
					studentList[inputPersonne-1].prenom,
					studentList[inputPersonne-1].nom)
				return true
			} else {
				fmt.Printf("%s %s erreur lors de sa sortie !\n",
					studentList[inputPersonne-1].prenom,
					studentList[inputPersonne-1].nom)
				return true
			}
		} else {
			fmt.Printf("%s %s n'a pas été trouvé !\n",
				studentList[inputPersonne-1].prenom,
				studentList[inputPersonne-1].nom)
			return false
		}

	} else {
		fmt.Printf("%s %s n'est pas dans le batiment !\n",
			studentList[inputPersonne-1].prenom,
			studentList[inputPersonne-1].nom)
		return false
	}
}

func (space *TupleSpaceBatiment) findInBatiments(p Personne) Batiment {
	findedBatiment := Batiment{}
	recv1 := space.Read(tuplespace.New(0))
	tuple1 := <-recv1
	var batimentList []Batiment = tuple1.Values()[0].([]Batiment)
	for _, batiment := range batimentList {
		recv2 := batiment.personnes.TupleSpace.Read(tuplespace.New(0))
		for personne := range recv2 {
			if personne != nil {
				if reflect.DeepEqual(p, personne.Values()[0].(Personne)) {
					findedBatiment = batiment
				}
			} else {
				logrus.Warnf("%v is nil\n", batiment.personnes)
			}
		}
	}
	return findedBatiment
}

func (space *TupleSpaceBatiment) isAlreadyInBatiment(p Personne) bool {
	res := false
	recv1 := space.Read(tuplespace.New(0))
	tuple1 := <-recv1
	var batimentList []Batiment = tuple1.Values()[0].([]Batiment)
	for _, batiment := range batimentList {
		recv2 := batiment.personnes.TupleSpace.Read(tuplespace.New(0))
		for personne := range recv2 {
			if personne != nil {
				if reflect.DeepEqual(p, personne.Values()[0].(Personne)) {
					res = true
				}
			} else {
				logrus.Warnf("%v is nil\n", batiment.personnes)
				break
			}
		}
	}
	return res
}

func (space *TupleSpaceBatiment) getAllPersonneInBatiment() []Personne {
	var inputBatiment int
	var batimentList = spaceBatiment.getAllBatiment()
	validInput := false
	reader := bufio.NewReader(os.Stdin)
	for !validInput {
		for i, batiment := range batimentList {
			fmt.Printf("%d - %s\n", i+1, batiment.nom)
		}
		fmt.Println("\nVoir les personnes dans quel batiment ?")
		inputText, _ := reader.ReadString('\n')
		inputText = strings.Replace(inputText, "\n", "", -1)
		inputBatiment, _ = strconv.Atoi(inputText)
		if inputBatiment <= 0 || inputBatiment > len(batimentList) {
			logrus.Warnf("Batiment invalide! Merci de refaire votre choix.")
		} else {
			validInput = true
		}
	}
	var personneList []Personne
	recv2 := batimentList[inputBatiment-1].personnes.TupleSpace.Read(tuplespace.New(0))
	for personne := range recv2 {
		if personne != nil {
			personneList = append(personneList, personne.Values()[0].(Personne))
		}
	}
	return personneList
}

func (badge *Badge) desactiverBadge() {
	badge.actif = false
}

func (badge *Badge) creerBadge(p Personne) {
	badge.actif = true
	badge.identifiant = p.identifiant
}

func (space *TupleSpacePersonnes) desactiverBadge() {
	recv1 := space.Read(tuplespace.New(0))
	tuple := <-recv1
	var studentList []Personne = tuple.Values()[0].([]Personne)
	for i, personne := range studentList {
		fmt.Printf("%d - %s %s\n", i+1, personne.nom, personne.prenom)
	}
	var validInput bool
	var inputPersonne int
	reader := bufio.NewReader(os.Stdin)
	for !validInput {
		fmt.Println("\nA qui desactiver le badge ?")
		inputText, _ := reader.ReadString('\n')
		inputText = strings.Replace(inputText, "\n", "", -1)
		inputPersonne, _ = strconv.Atoi(inputText)
		if inputPersonne <= 0 || inputPersonne > len(studentList) {
			logrus.Warnf("Personne invalide! Merci de refaire votre choix.")
		} else {
			validInput = true
		}
	}
	studentList[inputPersonne].badge.desactiverBadge()
}

func (space *TupleSpacePersonnes) addbadge() {
	recv1 := space.Read(tuplespace.New(0))
	tuple := <-recv1
	var studentList []Personne = tuple.Values()[0].([]Personne)
	for i, personne := range studentList {
		fmt.Printf("%d - %s %s\n", i+1, personne.nom, personne.prenom)
	}
	fmt.Println("\nA qui faire un nouveau badge ?")
	var validInput bool
	var inputPersonne int
	reader := bufio.NewReader(os.Stdin)
	for !validInput {
		inputText, _ := reader.ReadString('\n')
		inputText = strings.Replace(inputText, "\n", "", -1)
		inputPersonne, _ = strconv.Atoi(inputText)
		if inputPersonne <= 0 || inputPersonne > len(studentList) {
			logrus.Warnf("Personne invalide! Merci de refaire votre choix.")
		} else {
			validInput = true
		}
	}
	studentList[inputPersonne-1].badge.creerBadge(studentList[inputPersonne-1])
}

func (batiment *Batiment) personHaveAccess(personne Personne) bool {
	for _, personneRole := range personne.role {
		if personneRole == batiment.role {
			return true
		}
	}
	return false
}
