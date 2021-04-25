package update

import (
	"fmt"
	"log"
	"strings"

	"github.com/tcnksm/go-latest"
)

func CheckUpdate(version string) {

	defer (func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	})()

	if version == "development" {
		panic(fmt.Errorf("use production build for check update"))
	}

	githubTag := &latest.GithubTag{
		Owner:             "iamganeshagrawal",
		Repository:        "todo-cli",
		FixVersionStrFunc: latest.DeleteFrontV(),
	}

	res, err := latest.Check(githubTag, strings.Replace(version, "v", "", 1))
	if err != nil {
		fmt.Println("update check failed.")
	}
	if res.Outdated {
		fmt.Printf("%s is not latest\nUpgrade to v%s\nChange Log:\n\t%s", version, res.Current, res.Meta.Message)
	} else {
		fmt.Println("you have alredy latest version.")
	}
}
