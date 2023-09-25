// créer un programme go qui permet d'afficher a partir d'entrées utilisateur sur terminal un Mapping de villes avec leurs informations principales tels que le nombre d'habitants, le code postal, la superficie. Ces informations doivent etre dans une structure nommée "Ville".



package main

import "fmt"

type Ville struct {
	nom string
	nbHabitants int
	codePostal int
	superficie int
}

func main(){

	



	villes := make(map[string]Ville) // ou var mapping = map[string]Ville{}
	var nbrofcities int
	fmt.Printf("Entre le nombre de la villes")
	fmt.Scanln(&nbrofcities)
	fmt.Println("Numero de villes : ", nbrofcities)

	for i := 0; i < nbrofcities; i++ {
		var nomVille string
		var nbHabitantsn, codePostaln, superficien int
		fmt.Println("Entre le numero de la ville : ", i+1)
		fmt.Printf("Entre le nom de la ville : ")
		fmt.Scanln(&nomVille)
		fmt.Printf("Entre le nombre d'habitants : ")
		fmt.Scanln(&nbHabitantsn)
		fmt.Printf("Entre le code postal : ")
		fmt.Scanln(&codePostaln)
		fmt.Printf("Entre la superficie : ")
		fmt.Scanln(&superficien)
		villes[nomVille] = Ville{
			nom: nomVille,
			nbHabitants: nbHabitantsn,
			codePostal: codePostaln,
			superficie: superficien,
		}

	}
}
