// main.go

package main

import (
	"bufio"
	"estiam/dictionary"
	"fmt"
	"os"
)

func main() {

	myDictionary := dictionary.New()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("1. Ajouter un élément")
		fmt.Println("2. Définir une valeur")
		fmt.Println("3. Supprimer un élément")
		fmt.Println("4. Afficher la liste")
		fmt.Println("5. Quitter")

		fmt.Print("Choisissez une option (1-5): ")
		option, _, _ := reader.ReadRune()

		switch option {
		case '1':
			actionAdd(myDictionary, reader)
		case '2':
			actionDefine(myDictionary, reader)
		case '3':
			actionRemove(myDictionary, reader)
		case '4':
			actionList(myDictionary)
		case '5':
			fmt.Println("Au revoir!")
			return
		default:
			fmt.Println("Option invalide. Veuillez choisir une option valide.")
		}
	}
}

func actionAdd(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Print("Entrez le mot: ")
	word, _, _ := reader.ReadLine()

	fmt.Print("Entrez la définition: ")
	definition, _, _ := reader.ReadLine()

	d.Add(string(word), string(definition))
}

func actionDefine(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Print("Entrez le mot à définir: ")
	word, _, _ := reader.ReadLine()

	fmt.Print("Entrez la nouvelle définition: ")
	definition, _, _ := reader.ReadLine()

	d.Define(string(word), string(definition))
}

func actionRemove(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Print("Entrez le mot à supprimer: ")
	word, _, _ := reader.ReadLine()

	d.Remove(string(word))
}

func actionList(d *dictionary.Dictionary) {
	wordList, entryMap := d.List()

	fmt.Println("Liste des mots :")
	for _, word := range wordList {
		entry := entryMap[word]
		fmt.Printf("%s: %s\n", word, entry.String())
	}
}
