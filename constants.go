package main

var listePersonne = []Personne{
	{
		nom:         "Bollon",
		prenom:      "Hugo",
		identifiant: 1,
		role:        Enseigant,
		badge: Badge{
			identifiant: 1,
			actif:       true,
		},
	},

	{
		nom:         "Rodriguez-Lozano",
		prenom:      "Samuel",
		identifiant: 2,
		role:        Enseigant,
		badge: Badge{
			identifiant: 2,
			actif:       true,
		},
	},

	{
		nom:         "Kubasik",
		prenom:      "Tom",
		identifiant: 3,
		role:        Etudiant,
		badge: Badge{
			identifiant: 3,
			actif:       true,
		},
	},

	{
		nom:         "Cutting",
		prenom:      "Laurent",
		identifiant: 4,
		role:        Etudiant,
		badge: Badge{
			identifiant: 4,
			actif:       true,
		},
	},

	{
		nom:         "Hersemeule",
		prenom:      "Hugo",
		identifiant: 5,
		role:        Etudiant,
		badge: Badge{
			identifiant: 5,
			actif:       true,
		},
	},
}

var listeBatiments = []Batiment{
	{
		nom:    "salle1",
		porte:  true,
		role:   SalleTD,
		alarme: false,
		laser:  true,
	},
	{
		nom:    "salle2",
		porte:  false,
		role:   SalleTD,
		alarme: false,
		laser:  true,
	},
	{
		nom:    "Bureau3",
		porte:  false,
		role:   Bureau,
		alarme: false,
		laser:  true,
	},
	{
		nom:    "salle4",
		porte:  false,
		role:   SalleLangue,
		alarme: false,
		laser:  true,
	},
	{
		nom:    "multimedia5",
		porte:  true,
		role:   SalleMultimedia,
		alarme: false,
		laser:  true,
	},
	{
		nom:    "salle6",
		porte:  true,
		role:   SalleTP,
		alarme: false,
		laser:  true,
	},
	{
		nom:    "salle7",
		porte:  true,
		role:   SalleTD,
		alarme: false,
		laser:  true,
	},
}
