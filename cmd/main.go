package main

import (
	"github.com/BerdanAkbulut/task-app-backend/pkg"
	"github.com/BerdanAkbulut/task-app-backend/utils"
)

func main() {
	utils.RunApp(&pkg.App{Port: "8081"})
}
