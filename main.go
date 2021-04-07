package main

import (
	"os"
)

func main() {
	_ = os.Setenv("DATABASE_URL", "postgresql://postgres:postgres@0.0.0.0:5432/postgres")
}
