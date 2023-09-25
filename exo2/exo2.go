// créer un programme go qui permet de calculer le volume de transaction d'un Exchange par unité de temps (heure, minute, seconde). L'utilisateur devra renseigner le nombre de transactions, et un laps de temps aléatoire pour le temps d'execution de celles-ci.


package exo2

import "fmt"



func main() {
var nbroftransactions int
var time string
var hour int
var minute int
var second int
fmt.Printf("Entre le nombre de transactions : ")
fmt.Scanln(&nbroftransactions)
fmt.Printf("Entrez l'unité de temps (heure, minute, seconde) : ")
fmt.Scanln(&time)
switch time {
case "heure":
	fmt.Println("Nombre de transactions par heure : ", nbroftransactions/hour)
case "minute":
	fmt.Println("Nombre de transactions par minute : ", nbroftransactions/minute)
case "seconde":
	fmt.Println("Nombre de transactions par seconde : ", nbroftransactions/second)
default:
	fmt.Println("Erreur")



	
}
}