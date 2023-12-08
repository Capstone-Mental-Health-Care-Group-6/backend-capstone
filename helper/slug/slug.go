package slug

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

type SlugInterface interface {
	GenerateSlug(name string) string
}

type slug struct{}

func New() SlugInterface {
	return &slug{}
}

func (sl *slug) GenerateSlug(name string) string {
	numberRand := func() string {
		return strconv.Itoa(rand.Intn(1000))
	}

	name = strings.ReplaceAll(name, " ", "-")
	reg := regexp.MustCompile("[^a-zA-Z0-9-]+")
	name = reg.ReplaceAllString(name, "")
	randomNumber := numberRand()

	return fmt.Sprintf("%s-%s", randomNumber, strings.ToLower(name))
}
