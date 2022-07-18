package internal

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
)

func TimeInput(label string) string {
	prompt := promptui.Prompt{
		Label:     label,
		AllowEdit: true,
		Validate: func(input string) error {
			r, _ := regexp.Compile("[0-9]{2}:[0-9]{2}")

			if !r.MatchString(input) {
				return errors.New("required format: xx:xx")
			}

			time := strings.Split(input, ":")
			h, _ := strconv.Atoi(time[0])
			m, _ := strconv.Atoi(time[1])

			if h > 23 || m > 59 {
				return errors.New("must be valid 24hr time")
			}

			return nil
		},
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}

func DescriptionInput() string {
	prompt := promptui.Prompt{
		Label:     "description",
		AllowEdit: true,
		Validate: func(input string) error {
			if len(input) == 0 {
				return errors.New("description required")
			}

			return nil
		},
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}
