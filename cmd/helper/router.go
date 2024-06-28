package helper

import (
	"fmt"

	"github.com/assignment-amori/internal/constant"
)

func GenModulePattern(m constant.Module) string {
	return fmt.Sprintf("/%s", m)
}
