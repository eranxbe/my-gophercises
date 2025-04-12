package chooseYourOwnAdventure

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

// create the same functionality but for a cli app
// reading the json is the same JsonStory
// print chapter and options
// get user input
// print next chapter
// when no options (aka end of game) - exit

func PrintChapter(c Chapter) {
	fmt.Println(c.Title + "\n\n")
	for _, par := range c.Paragraphs {
		fmt.Println(par + "\n")
	}
	if len(c.Options) == 0 {
		fmt.Println("FIN.")
		os.Exit(0)
	}
	fmt.Println("Choose Your Answer:")
	for i, option := range c.Options{
		fmt.Printf("%d) %s\n", i + 1, option.Text)
	}
	fmt.Print("#")
}

func ProcessAnswer(c Chapter) (string, error) {
	var answer string
	fmt.Scan(&answer)
	num, err := strconv.Atoi(answer) // Atoi = ASCII to int
    if err != nil {
		return "", err
    }
	return c.Options[num - 1].Chapter, nil
}

func CmdCYOA() {
	filename := flag.String("file", "gopher.json", "the JSON file with the story")
	flag.Parse()
	file, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	story, err := JsonStory(file)
	if err != nil {
		panic(err)
	}
	currChapter := story["intro"]
	PrintChapter(currChapter)
	for {
		nextChapter, err := ProcessAnswer(currChapter)
		if err != nil {
			panic(err)
		}
		PrintChapter(story[nextChapter])
		currChapter = story[nextChapter]
	}
}