import (
	"fmt"
	"strings"

	playgroundvalidator "github.com/go-playground/validator/v10"
)

var validate = playgroundvalidator.New()

// Validate checks struct fields based on their `validate` tags.
func Validate(s any) error {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	var msgs []string
	for _, e := range err.(playgroundvalidator.ValidationErrors) {
		msgs = append(msgs, fmt.Sprintf("%s: failed on '%s'", e.Field(), e.Tag()))
	}
	return fmt.Errorf("validation failed: %s", strings.Join(msgs, "; "))
}
