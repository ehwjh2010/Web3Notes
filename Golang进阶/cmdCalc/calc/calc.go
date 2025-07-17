package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/urfave/cli/v3"
	"log"
	"math"
	"os"
)

func printErr(err error) {
	fmt.Printf("calc error: %s\n", err)
}

func printResult(r float64) {
	fmt.Printf("Result: %.2f\n", r)
}

func Echo(r float64, err error) error {
	if err != nil {
		printErr(err)
	} else {
		printResult(r)
	}
	return nil
}

func calculate(elements []float64, operation string) (float64, error) {
	r := elements[0]
	switch operation {
	case "+":
		for _, v := range elements[1:] {
			r += v
		}
		return r, nil
	case "-":
		for _, v := range elements[1:] {
			r -= v
		}
		return r, nil

	case "*":
		for _, v := range elements[1:] {
			r *= v
		}
		return r, nil

	case "/":
		for _, v := range elements[1:] {
			if v == 0 {
				return 0, errors.New("division by zero")
			}
			r /= v
		}
		return r, nil

	case "%":
		for _, v := range elements[1:] {
			if v == 0 {
				return 0, errors.New("modulo by zero")
			}

			r = math.Mod(r, v)
		}
		return r, nil
	default:
		return 0, errors.New("invalid operation")
	}
}

func main() {
	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:                  "add",
				Usage:                 "add two numbers.",
				Aliases:               []string{"+"},
				UsageText:             "calc add 1 2",
				EnableShellCompletion: true,
				Arguments: []cli.Argument{
					&cli.Float64Args{
						Name: "addInput",
						Min:  2,
						Max:  2,
					},
				},
				Action: func(ctx context.Context, command *cli.Command) error {
					r, err := calculate(command.Float64Args("addInput"), "+")
					return Echo(r, err)
				},
			},
			{
				Name:                  "sub",
				Usage:                 "sub two numbers.",
				Aliases:               []string{"-"},
				UsageText:             "calc sub 1 2",
				EnableShellCompletion: true,
				Arguments: []cli.Argument{
					&cli.Float64Args{
						Name: "subInput",
						Min:  2,
						Max:  2,
					},
				},
				Action: func(ctx context.Context, command *cli.Command) error {
					r, err := calculate(command.Float64Args("subInput"), "-")
					return Echo(r, err)
				},
			},

			{
				Name:                  "mul",
				Usage:                 "mul two numbers.",
				Aliases:               []string{"*"},
				UsageText:             "calc mul 1 2",
				EnableShellCompletion: true,
				Arguments: []cli.Argument{
					&cli.Float64Args{
						Name: "mulInput",
						Min:  2,
						Max:  2,
					},
				},
				Action: func(ctx context.Context, command *cli.Command) error {
					r, err := calculate(command.Float64Args("mulInput"), "*")
					return Echo(r, err)
				},
			},

			{
				Name:                  "div",
				Usage:                 "div two numbers.",
				UsageText:             "calc div 8 4",
				Aliases:               []string{"/"},
				EnableShellCompletion: true,
				Arguments: []cli.Argument{
					&cli.Float64Args{
						Name: "divInput",
						Min:  2,
						Max:  2,
					},
				},
				Action: func(ctx context.Context, command *cli.Command) error {
					r, err := calculate(command.Float64Args("divInput"), "/")
					return Echo(r, err)
				},
			},

			{
				Name:                  "mod",
				Usage:                 "mod two numbers.",
				UsageText:             "calc mod 1 2",
				Aliases:               []string{"%"},
				EnableShellCompletion: true,
				Arguments: []cli.Argument{
					&cli.Float64Args{
						Name: "modInput",
						Min:  2,
						Max:  2,
					},
				},
				Action: func(ctx context.Context, command *cli.Command) error {
					r, err := calculate(command.Float64Args("modInput"), "%")
					return Echo(r, err)
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
