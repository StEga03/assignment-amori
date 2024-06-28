package testfiles

import (
	"context"
	"math/rand"

	"github.com/assignment-amori/pkg/consistency"
	"github.com/google/uuid"
)

func DoRunAsUnit(ctx context.Context, action func(*consistency.ConsistencyElement) error) error {
	var cel *consistency.ConsistencyElement
	err := action(cel)
	return err
}

// IMPORTANT: Only call this function inside test files!
// Can call this multiple times to reset the generator.
func MockUUID() []uuid.UUID {
	rnd := rand.New(rand.NewSource(1))
	uuid.SetRand(rnd)

	// In every universe, the order of return of this seed will always be the same.
	return []uuid.UUID{
		uuid.MustParse("52fdfc07-2182-454f-963f-5f0f9a621d72"),
		uuid.MustParse("9566c74d-1003-4c4d-bbbb-0407d1e2c649"),
		uuid.MustParse("81855ad8-681d-4d86-91e9-1e00167939cb"),
		uuid.MustParse("6694d2c4-22ac-4208-a007-2939487f6999"),
		uuid.MustParse("eb9d18a4-4784-445d-87f3-c67cf22746e9"),
	}
}
