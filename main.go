package main

import (
	"crypto/sha256"
	"math/rand"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pastc/vero/v2"
)

type Dice struct {
	BetAmount  float64 `json:"betAmount" form:"betAmount" query:"betAmount"`
	RollValue  float64 `json:"rollValue" form:"rollValue" query:"rollValue"`
	Multiplier float64 `json:"multiplier" form:"multiplier" query:"multiplier"`
	WinChance  float64 `json:"winChance" form:"winChance" query:"winChance"`
	Roll       bool    `json:"roll" form:"roll" query:"roll"`
}

type DiceDTO struct {
	BetAmount  float64
	RollValue  float64
	Multiplier float64
	WinChance  float64
	Roll       bool
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.POST("/dice", func(c echo.Context) error {
		u := new(Dice)
		if err := c.Bind(u); err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		// Load into separate struct for security
		dice := DiceDTO{
			BetAmount:  u.BetAmount,
			RollValue:  u.RollValue,
			Multiplier: u.Multiplier,
			WinChance:  u.WinChance,
			Roll:       u.Roll,
		}
		serverSeed := Hash256("SERVER_SEED")
		clientSeed := Hash256("CLIENT_SEED")
		nonce := rand.Int()
		value, err := vero.Dice(serverSeed, clientSeed, nonce, 0)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Internal Server Error")
		}
		if dice.Roll {
			if float64(value)/float64(100) > dice.RollValue {
				return c.JSON(http.StatusOK, map[string]interface{}{
					"result":     "win",
					"value":      value,
					"serverSeed": serverSeed,
					"clientSeed": clientSeed,
					"nonce":      nonce,
				})
			}
			return c.JSON(http.StatusOK, map[string]interface{}{
				"result":     "lose",
				"value":      value,
				"serverSeed": serverSeed,
				"clientSeed": clientSeed,
				"nonce":      nonce,
			})
		} else {
			if float64(value)/float64(100) < dice.RollValue {
				return c.JSON(http.StatusOK, map[string]interface{}{
					"result":     "win",
					"value":      value,
					"serverSeed": serverSeed,
					"clientSeed": clientSeed,
					"nonce":      nonce,
				})
			} else {
				return c.JSON(http.StatusOK, map[string]interface{}{
					"result":     "lose",
					"value":      value,
					"serverSeed": serverSeed,
					"clientSeed": clientSeed,
					"nonce":      nonce,
				})
			}
		}
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func Hash256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return string(h.Sum(nil))
}
