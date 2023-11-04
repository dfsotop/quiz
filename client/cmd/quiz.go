package cmd

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
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
		// Use the response
		for i, question := range response.Questions {
			fmt.Printf("Question %d: %s\n", i+1, question.Text)
			for j, option := range question.Options {
				fmt.Printf("Option %d: %s\n", j+1, option.Text)
			}

			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter the option number: ")
			text, _ := reader.ReadString('\n')
			text = strings.TrimSpace(text)
			optionNumber, err := strconv.ParseInt(text, 10, 32)
			if err != nil {
				log.Fatalf("invalid input: %v", err)
			}

			fmt.Printf("You chose option %d\n", optionNumber)
			answer := &pb.Answer{
				QuestionId:     question.Id,
				ChosenOptionId: (int32)(optionNumber),
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
		fmt.Printf("Your score is %d\n", registerAnswersRes.QuizRating)
	},
}

func init() {
	rootCmd.AddCommand(quizCmd)
}
