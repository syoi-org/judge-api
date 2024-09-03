package main

import "github.com/syoi-org/judge-api/codeforces"

func main() {
	service := codeforces.NewService(codeforces.ServiceParams{
		Config: &codeforces.Config{
			BaseURL: "http://localhost:8080",
		},
	})
	service.GetProblem(2008, "A")
}
