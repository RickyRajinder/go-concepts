package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {

	//Tests a long string (Excerpt rom `The Great Gatsby`) 948 letters
	key := generateKey()
	longMsg := "He didn’t say any more, but we’ve always been unusually communicative in a reserved way, and I understood that he meant a great deal more than that. In consequence, I’m inclined to reserve all judgments, a habit that has opened up many curious natures to me and also made me the victim of not a few veteran bores. The abnormal mind is quick to detect and attach itself to this quality when it appears in a normal person, and so it came about that in college I was unjustly accused of being a politician, because I was privy to the secret griefs of wild, unknown men. Most of the confidences were unsought — frequently I have feigned sleep, preoccupation, or a hostile levity when I realized by some unmistakable sign that an intimate revelation was quivering on the horizon; for the intimate revelations of young men, or at least the terms in which they express them, are usually plagiaristic and marred by obvious suppressions. Reserving judgments is a matter of infinite hope. I am still a little afraid of missing something if I forget that, as my father snobbishly suggested, and I snobbishly repeat, a sense of the fundamental decencies is parcelled out unequally at birth."
	encryptedLong := encrypt(key, longMsg)
	decryptedLong := decrypt(key, encryptedLong)
	println(encryptedLong)
	println(decryptedLong)
	println(findAndReplace(encryptedLong))

	println("______________________________________________________________________________________________\n")

	//Tests a ten letter string
	key = generateKey()
	tenLetterString := "This is a ten"
	encryptedShort := encrypt(key, tenLetterString)
	decryptedShort := decrypt(key, encryptedShort)
	println(encryptedShort)
	println(decryptedShort)
	println(findAndReplace(encryptedShort))

	println("______________________________________________________________________________________________\n")

	//Attempts to decrypt an encrypted novel, 551714 letters
	key = generateKey()
	file, err := os.Open("pride-and-prejudice.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	book := string(bytes)
	encryptedBook := encrypt(key, book)
	println(findAndReplace(encryptedBook))
}
