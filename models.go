package main

import "fmt"

var emptyBadge = Badge{}

type RoleSalle int

const (
	Bureau RoleSalle = iota
	SalleTP
	SalleExperimentation
	SalleTD
	SalleMultimedia
	SalleLangue
)

type RolePersonne []RoleSalle

var (
	EnseigantChercheur     RolePersonne = []RoleSalle{SalleTP, SalleTD, Bureau, SalleExperimentation, SalleMultimedia, SalleLangue}
	PersonnelAdministratif              = []RoleSalle{SalleTP, SalleTD, Bureau, SalleExperimentation, SalleMultimedia, SalleLangue}
	Etudiant                            = []RoleSalle{SalleTP, SalleTD, SalleLangue, SalleMultimedia}
)

type Badge struct {
	identifiant int
	actif       bool
}

type Personne struct {
	nom         string
	prenom      string
	identifiant int
	role        RolePersonne
	badge       Badge
}

type Batiment struct {
	nom       string
	porte     bool
	role      RoleSalle
	alarme    bool
	laser     bool
	personnes TupleSpacePersonnes
}

type JournalDeBord struct {
	identite Personne
	heure    int
	batiment Batiment
}

type Badgeur struct {
	voyantVert  bool
	voyantRouge bool
}

type SalleDeCommande struct {
	responsable Personne
}

func (s *SalleDeCommande) CreerBadge(personne *Personne) error {
	if personne.badge == emptyBadge {
		return fmt.Errorf("Cette personne a déjà un badge")
	}
	personne.badge = Badge{
		identifiant: personne.identifiant,
		actif:       true,
	}
	return nil
}

func (s *SalleDeCommande) DesactiverBadge(badge *Badge) error {
	if !badge.actif {
		return fmt.Errorf("Ce badge est déja désactivé")
	}
	badge.actif = true
	return nil
}

func (s *SalleDeCommande) ListePersonneBatiment(batiment *Batiment) {
	return
}
