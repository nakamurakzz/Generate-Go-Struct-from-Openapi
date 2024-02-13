// Code generated from Pkl module `nakamurakzz.presentation.response.person`. DO NOT EDIT.
package person

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type Person struct {
	Id int `pkl:"id" json:"id"`

	Name string `pkl:"name" json:"name"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Person
func LoadFromPath(ctx context.Context, path string) (ret *Person, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.FileSource(path))
	return ret, err
}

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Person
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*Person, error) {
	var ret Person
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
