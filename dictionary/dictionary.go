package dictionary

import "fmt"

type Entry struct {
}

func (e Entry) String() string {

	return ""
}

type Dictionary struct {
	entries map[string]Entry
}

func New() *Dictionary {
	return &Dictionary{entries: make(map[string]Entry)}
}

func (d *Dictionary) Add(word string, definition string) {
	d.entries[word] = Entry{}
}

func (d *Dictionary) Define(word string, definition string) {
	if _, exists := d.entries[word]; exists {
		// La clé existe, mise à jour de la définition
		d.entries[word] = Entry{}
	} else {
		// La clé n'existe pas, ajout d'une nouvelle entrée
		d.Add(word, definition)
	}
}

func (d *Dictionary) Get(word string) (Entry, error) {
	entry, exists := d.entries[word]
	if !exists {
		return Entry{}, fmt.Errorf("la clé '%s' n'existe pas dans le dictionnaire", word)
	}
	return entry, nil
}

func (d *Dictionary) Remove(word string) {
	delete(d.entries, word)
}

func (d *Dictionary) List() ([]string, map[string]Entry) {
	wordList := make([]string, 0, len(d.entries))
	for word := range d.entries {
		wordList = append(wordList, word)
	}
	return wordList, d.entries
}
