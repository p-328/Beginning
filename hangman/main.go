package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
	"io/ioutil"
	"strings"
	"unicode"
	"sort"	
)
func binarySearch(needle string, haystack []string) bool {
	sort.Slice(haystack, func (i, j int) bool {
		return haystack[i] < haystack[j]
	})
	low := 0
	high := len(haystack) - 1

	for low <= high{
		median := (low + high) / 2

		if haystack[median] < needle {
			low = median + 1
		}else{
			high = median - 1
		}
	}

	if low == len(haystack) || haystack[low] != needle {
		return false
	}

	return true
}
func isEnglishWord(s string) bool {
	dict, err := ioutil.ReadFile("words.txt");
	if err != nil {
		fmt.Println(err);
		return false;
	}
	words := string(dict);
	wordList := strings.Fields(words);
	if unicode.IsUpper(rune(s[0])) && unicode.IsLetter(rune(s[0])) {
		for _, word := range wordList {
			if unicode.IsLetter(rune(word[0])) {
				unicode.ToUpper(rune(word[0]))
			}
		}
	}
	if binarySearch(s, wordList) {
		return true;
	}
	
	return false;
}
func findLetterInString(letter string, s string) []int {
	letter_bytes := []byte(letter)
	string_bytes := []byte(s)
	var instances []int
	letter_guess := string(letter_bytes[0])
	for i := 0; i < len(string_bytes); i++ {
		if letter_guess == string(string_bytes[i]) {
			instances = append(instances, i+1)
		}
	}
	return instances;
}
func checkDuplicate(arr []string) bool {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				return true;
			}
		} 
	}
	return false
}


func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Hangman");
	for i := 0; i < 50; i++ {
		fmt.Print("-");
	}
	var words []string
	var wordAmt int
	attempts := 0;
	attemptMax := 7;
	fmt.Println();
	fmt.Println("First: we need to enter the amount of words needed. So, enter the amount of words you want to have in your game.");
	fmt.Scan(&wordAmt);
	if wordAmt == 0 {
		fmt.Println("That ain't a valid amount of words chief, so imma use the default.");
		wordAmt = 20;
	}
	words = make([]string, wordAmt)
	fmt.Println("Second we need to enter in words. So enter", wordAmt, "words, please")
	for i := 0; i < wordAmt; i++ {
		fmt.Scan(&words[i])
		phrase := strings.Fields(words[i])
		for _, word := range phrase {
			if !isEnglishWord(word) {
				fmt.Println("Imma stop you right there chief, one of your words isn't english.")
				os.Exit(-1)
			}
		}
	}
	if checkDuplicate(words) {
		fmt.Println("You have a duplicate in your word set.")
		os.Exit(-1);
	}
	fmt.Println("Third, let's choose a mode. Enter easy for 10 attempts, hard for 5 attempts(default is medium for 7 attempts) ")
	var mode string
	fmt.Scan(&mode)
	switch mode {
		case "easy":
			attemptMax = 10;
		case "hard":
			attemptMax = 5;
		default:
			break;
	}
	fmt.Println("You are ready to go!");
	fmt.Println("BEGIN!");
	for i := 0; i < 50; i++ {
		fmt.Print("-");
	}
	fmt.Println()
	var running bool = true;
	wordToGuess := words[rand.Intn(len(words)-1)];
	for attempts <= attemptMax && running {
		fmt.Println("You have a few decisions: type word to guess the word, letter to guess a letter");
		fmt.Println("Enter your decision: ");
		var decision string;
		fmt.Scan(&decision)
		switch decision {
			case "letter":
				fmt.Println("Enter your letter: ");
				var letter string
				fmt.Scan(&letter)
				instances := findLetterInString(letter, wordToGuess)
				if len(instances) == 0 {
					fmt.Println("No instances of", letter)
					attempts++;
				} else {
					fmt.Println("Found instances of", letter)
					for i, val := range instances {
						fmt.Println("Instance", i, "at position no.", val)
					}
				}
			case "word":
				fmt.Println("Enter your word: ");
				var word string;
				fmt.Scan(&word)
				if word == wordToGuess {
					fmt.Println("Congratulations! You have guessed the word!")
					fmt.Println("Mistakes:", attempts);
					running = false;
				} else {
					fmt.Println("That ain't it chief");
					attempts++;
				}
			default:
				attempts = 10000;
		}
	}
	fmt.Println()
	for i := 0; i < 50; i++ {
		fmt.Print("-");
	}
	if attempts > attemptMax {
		fmt.Println("GAME OVER")
		fmt.Println("The word was: ", wordToGuess)
	}
	fmt.Println()
	if attempts == 10000 {
		fmt.Println("Nice try.")
	}
}
