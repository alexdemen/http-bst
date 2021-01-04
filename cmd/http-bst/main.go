package main

import (
	"encoding/json"
	"github.com/alexdemen/http-bst/internal/core/log"
	"github.com/alexdemen/http-bst/internal/core/middleware"
	"github.com/alexdemen/http-bst/internal/router"
	"github.com/alexdemen/http-bst/internal/service"
	"github.com/go-chi/chi"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	initFile := os.Getenv("TREE_INIT_FILE")

	logger := log.NewLogger(os.Stdout)

	treeService := service.NewTreeService(logger)
	if initFile != "" {
		if initData, err := readInitData(initFile); err != nil {
			return
		} else {
			for _, val := range initData {
				if err = treeService.Add(val); err != nil {
					return
				}

			}
		}
	}

	r := chi.NewRouter()
	r.Use(middleware.NewStructuredLogger(logger))
	r.Mount("/", router.NewBSTRouter(*treeService))

	if err := http.ListenAndServe(":"+port, r); err != nil {
		logger.Error(map[string]interface{}{
			"error_text": err.Error(),
		})
	}
}

func readInitData(initFile string) ([]int, error) {
	f, err := os.Open(initFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	fData, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	initData := make([]int, 0, 30)
	err = json.Unmarshal(fData, &initData)
	return initData, err
}
