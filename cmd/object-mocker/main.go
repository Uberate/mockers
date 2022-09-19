package main

import (
	"fmt"
	"mockers/internal/web/server"
	"mockers/pkg/objects/o_v2/web"
)

func main() {
	if err := server.Start(server.GinConfig{
		Host:                 "",
		Port:                 "3000",
		EnableReadinessProbe: true,
		EnableLivenessProbe:  true,
	}, web.Handlers()...); err != nil {
		fmt.Println(err)
	}
}
