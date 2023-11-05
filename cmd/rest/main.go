package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

const PORT = ":3000"

type Payload struct {
	Scores struct {
		Managers []struct {
			UserID int `json:"userId"`
			Score  int `json:"score"`
		} `json:"managers"`
		Team []struct {
			UserID int `json:"userId"`
			Score  int `json:"score"`
		} `json:"team"`
		Others []struct {
			UserID int `json:"userId"`
			Score  int `json:"score"`
		} `json:"others"`
	} `json:"scores"`
}

type Result struct {
	Success bool `json:"success"`
	Data Data `json:"data"`
	Errors []string `json:"errors"`
}

type Data struct {
	Scores struct{
		Managers int `json:"managers"`
		Team int `json:"team"`
		Others int `json:"others"`
	}`json:"scores"`
}

func main() {
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		var payload Payload

		 managerScores := 0
		 teamScores := 0
		 otherScores := 0

		result := Result{
			Success: false,
			Data: Data{},
			Errors: []string{},
		}

		
	
		if err := c.BodyParser(&payload); err != nil {
			return err
		}


		for _, v := range payload.Scores.Managers {
			managerScores = managerScores + v.Score
		}
		for _, v := range payload.Scores.Team {
			teamScores = teamScores + v.Score
		}
		for _, v := range payload.Scores.Others {
			otherScores = otherScores + v.Score
		}

		newData := Data{
			Scores: struct {
				Managers int `json:"managers"`
				Team     int `json:"team"`
				Others   int `json:"others"`
			}{
				Managers: managerScores,
				Team: teamScores,
				Others: otherScores,
			},
		}

		result.Success = true
		result.Data = newData


		fmt.Printf("%+v", payload)
	
		return c.JSON(result)
	})

    app.Listen(":3000")
}