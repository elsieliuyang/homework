package main

import (
	"fmt"
	"strings"
)

func main() {
	animals := []string{"smallkitty", "midllekitty", "bigkitty", "hugekitty"}
	fmt.Println(circulation(animals))
}

func circulation(animal []string) string {
	str := "There was an old lady who swallowed a fly.\n" +
		"I don't know why she swallowed a fly - perhaps she'll die!!!\n" +
		"There was an old lady who swallowed a spider;\n" +
		"That wriggled and wiggled and tickled inside her.\n" +
		"She swallowed the spider to catch the fly;\n" +
		"I don't know why she swallowed a fly - perhaps she'll die!!!\n" +
		"There was an old lady who swallowed a bird;\n" +
		"How absurd to swallow a bird.\n" +
		"She swallowed the bird to catch the spider,\n" +
		"She swallowed the spider to catch the fly;\n" +
		"I don't know why she swallowed a fly - perhaps she'll die!!!\n" +
		"There was an old lady who swallowed a frog;\n" +
		"Fancy that to swallow a frog!\n" +
		"She swallowed the frog to catch the bird,\n" +
		"She swallowed the bird to catch the spider,\n" +
		"She swallowed the spider to catch the fly;\n" +
		"I don't know why she swallowed a fly - perhaps she'll die!!!\n" +
		"There was an old lady who swallowed a dog;\n" +
		"What a hog, to swallow a dog!\n" +
		"She swallowed the dog to catch the frog,\n" +
		"She swallowed the frog to catch the bird,\n" +
		"She swallowed the bird to catch the spider,\n" +
		"She swallowed the spider to catch the fly;\n" +
		"I don't know why she swallowed a fly - perhaps she'll die!!!\n" +
		"There was an old lady who swallowed a cow;\n" +
		"I don't know how she swallowed a cow!\n" +
		"She swallowed the cow to catch the dog,\n" +
		"She swallowed the dog to catch the frog,\n" +
		"She swallowed the frog to catch the bird,\n" +
		"She swallowed the bird to catch the spider,\n" +
		"She swallowed the spider to catch the fly;\n" +
		"I don't know why she swallowed a fly - perhaps she'll die!!!\n" +
		"There was an old lady who swallowed a horse...\n" +
		"...She's dead, of course!"
	a := []string{"fly", "spider", "bird", "frog", "dog", "cow", "horse"}
	strArray := strings.Split(str, "!!")
	var result string
	for i := 0; i < len(animal); i++ {
		result += strArray[i]
		for j := i; j >= 0; j-- {
			result = strings.Replace(result, a[j], animal[j], -1)
		}
	}
	return result
}
