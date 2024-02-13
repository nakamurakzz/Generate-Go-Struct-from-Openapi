// Code generated from Pkl module `nakamurakzz.presentation.response.group`. DO NOT EDIT.
package group

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type Group struct {
	GroupName string `pkl:"groupName" json:"groupName"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Group
func LoadFromPath(ctx context.Context, path string) (ret *Group, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Group
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*Group, error) {
	var ret Group
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
