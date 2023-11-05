package cmd

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	pb "github.dfsotop.quiz/m/proto"

	"google.golang.org/grpc"

	"github.com/spf13/cobra"
)

var quizCmd = &cobra.Command{
	Use:   "quiz",
	Short: "Start the quiz",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()

		c := pb.NewQuizServiceClient(conn)

		// Now you can call the service methods:
		response, err := c.GetQuestions(context.Background(), &pb.GetQuestionsRequest{})
		if err != nil {
			log.Fatalf("could not get questions: %v", err)
		}

		answers := make([]*pb.Answer, 0)
		for i, question := range response.Questions {
			var option string
			for {
				validOptions := make(map[string]bool)
				fmt.Printf("Question %d: %s\n", i+1, question.Text)
				for _, option := range question.Options {
					fmt.Printf("%s) %s\n", option.Id, option.Text)
					validOptions[option.Id] = true
				}

				reader := bufio.NewReader(os.Stdin)
				fmt.Print("Enter the option of your choice: ")
				text, _ := reader.ReadString('\n')
				option = strings.TrimSpace(text)

				if _, ok := validOptions[option]; ok {
					break
				} else {
					fmt.Println("Invalid option. Please try again.")
				}
			}

			fmt.Printf("You chose option %s\n", option)
			answer := &pb.Answer{
				QuestionId:     question.Id,
				ChosenOptionId: option,
			}
			answers = append(answers, answer)
		}

		registerAnswersReq := &pb.RegisterAnswersRequest{
			Answers: answers,
		}

		registerAnswersRes, err := c.RegisterAnswers(context.Background(), registerAnswersReq)
		if err != nil {
			log.Fatalf("could not register answers: %v", err)
		}
		totalCorrect := 0
		totalWrong := 0
		for _, questionResult := range registerAnswersRes.QuesitonsResults {
			if questionResult.ChosenOptionId == questionResult.CorrectOptionId {
				totalCorrect++
			} else {
				totalWrong++
			}
		}

		fmt.Printf("Your score is %0.2f%%\n", registerAnswersRes.Score*100)
		fmt.Printf("Correct answers %d. Wrong answers: %d\n", totalCorrect, totalWrong)
		fmt.Printf("You were better than %0.2f%% of all quizzers\n", registerAnswersRes.Statistics.BetterThan)

	},
}

func init() {
	rootCmd.AddCommand(quizCmd)
}
