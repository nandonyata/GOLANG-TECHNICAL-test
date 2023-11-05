package main

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"
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
		Managers float64 `json:"managers"`
		Team float64 `json:"team"`
		Others float64 `json:"others"`
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

		var uniqueUser []int

		for _, v := range payload.Scores.Managers {
			managerScores = managerScores + v.Score

			if slices.Contains(uniqueUser, v.UserID) {
				result.Errors = append(result.Errors, "Duplicate userId on manager")
			}

			uniqueUser = append(uniqueUser, v.UserID)
		}
		
		if len(payload.Scores.Team) > 3 {
			for _, v := range payload.Scores.Team {
				teamScores = teamScores + v.Score

				if slices.Contains(uniqueUser, v.UserID) {
					result.Errors = append(result.Errors, "Duplicate userId on team")
				}

				uniqueUser = append(uniqueUser, v.UserID)
			}
		}

		if len(payload.Scores.Others) > 3 {
			for _, v := range payload.Scores.Others {
				otherScores = otherScores + v.Score

				if slices.Contains(uniqueUser, v.UserID) {
					result.Errors = append(result.Errors, "Duplicate userId on other")
				}

				uniqueUser = append(uniqueUser, v.UserID)
			}
		}

		// Return if there are errors
		if len(result.Errors) > 0 {
			return c.Status(fiber.StatusBadRequest).JSON(result)
		}else {
			newData := Data{
				Scores: struct {
					Managers float64 `json:"managers"`
				Team     float64 `json:"team"`
				Others   float64 `json:"others"`
			}{
				Managers: float64(managerScores) / float64(len(payload.Scores.Managers)),
				Team: float64(teamScores) / float64(len(payload.Scores.Team)),
				Others: float64(otherScores) / float64(len(payload.Scores.Others)),
			},
		}

		result.Success = true
		result.Data = newData
	}
// not done yettttt
		return c.Status(fiber.StatusCreated).JSON(result)
	})

    app.Listen(":3000")
}