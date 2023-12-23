package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type problem struct {
	question string
	answer   string
}

func readFileRecords(fileName string) []problem {
	file, err := os.Open(fileName + ".csv")
	if err != nil {
		log.Fatal("Error opening csv - ", err)
		return nil
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading file contents - ", err)
		return nil
	}

	problems := make([]problem, len(records))
	for i, record := range records {
		problems[i] = problem{
			question: record[0],
			answer:   record[1],
		}
	}

	return problems
}

func startQuiz(problems []problem, timerValue int) {
	correctAnswers := 0
	timer := time.NewTimer(time.Duration(timerValue) * time.Second)

	for _, r := range problems {
		answerCh := make(chan string)
		go func() {
			var ans string
			fmt.Println("Question - " + r.question + " ? ")

			_, err := fmt.Scanln(&ans)
			if err != nil {
				log.Fatal("Error in user input - ", err)
				return
			}

			if ans == r.answer {
				fmt.Println("Correct Answer")
				correctAnswers++
			} else {
				fmt.Println("Wrong Answer")
			}
			answerCh <- ans
		}()

		select {
		case <-timer.C:
			fmt.Println("Time out")
			return
		case <-answerCh:
			continue
		}
	}
	fmt.Printf("You answered %d correct answers out of %d questions\n", correctAnswers, len(problems))
}

func main() {
	fmt.Println("Welcome to Quiz App")

	fileNamePtr := flag.String("FileName", "Default", "Name of the flag")
	timerPtr := flag.Int("Timer", 20, "Timer value")

	flag.Parse()

	problems := readFileRecords(*fileNamePtr)
	if problems == nil {
		return
	}

	startQuiz(problems, *timerPtr)
}
