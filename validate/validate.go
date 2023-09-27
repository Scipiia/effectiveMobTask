package validate

import (
	"fmt"
	"strconv"
)

func ValidateName(name string) error {
	if name == "" {
		return fmt.Errorf("некоректное значение %s", name)
	}

	if _, err := strconv.Atoi(name); err != nil {
		return fmt.Errorf("некоректное значение %s", name)
	}

	return nil
}
