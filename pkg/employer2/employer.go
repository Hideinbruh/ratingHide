package employer2

import (
	"github.com/Hideinbruh/initializing_project/internal/handlers/employer"
)

func GetNameOfVladick() string {
	e := employer.New("Влаdick")

	return e.GetName()
}
