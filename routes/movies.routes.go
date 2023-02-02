package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Movie struct {
	Title string `json:"title"`
	Id    int    `json:"id"`
}

func UseMoviesRoutes(router fiber.Router) {

	movies := []*Movie{
		{
			Title: "Scary Movie",
			Id:    1,
		},
		{
			Title: "Cars",
			Id:    2,
		},
		{
			Title: "Ley de Calle",
			Id:    3,
		},
	}

	router.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"movies": movies,
		})
	})

	router.Get("/:id", func(c *fiber.Ctx) error {

		id, _ := c.ParamsInt("id")

		var movieFound Movie

		for _, movie := range movies {
			if movie.Id == id {
				movieFound = *movie
			}
		}

		return c.JSON(fiber.Map{
			"movie": movieFound,
		})
	})

	router.Post("/", func(c *fiber.Ctx) error {
		type Request struct {
			Title string
			Id    int
		}

		var body Request

		c.BodyParser(&body)

		fmt.Println(body)

		newMovie := Movie{
			Title: body.Title,
			Id:    len(movies) + 1,
		}

		fmt.Println(newMovie)

		movies = append(movies, &newMovie)

		return c.JSON(fiber.Map{
			"movies": movies,
		})

	})

	router.Put("/:id", func(c *fiber.Ctx) error {

		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "there's a error",
			})
		}

		type Request struct {
			Title string
		}

		var body Request

		c.BodyParser(&body)

		for _, movie := range movies {
			if movie.Id == id {
				movie.Title = body.Title
			}
		}

		return c.JSON(fiber.Map{
			"movies": movies,
		})
	})

	router.Delete("/:id", func(c *fiber.Ctx) error {

		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "there's a error",
			})
		}

		for idx, movie := range movies {
			if movie.Id == id {
				movies = append(movies[:idx], movies[idx+1:]...)
			}
		}

		return c.JSON(fiber.Map{
			"movies": movies,
		})
	})

}
