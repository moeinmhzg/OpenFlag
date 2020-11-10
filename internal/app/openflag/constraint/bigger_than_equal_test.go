package constraint_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/OpenFlag/OpenFlag/internal/app/openflag/constraint"

	"github.com/OpenFlag/OpenFlag/internal/app/openflag/model"
	"github.com/stretchr/testify/suite"
)

type BiggerThanEqualConstraintSuite struct {
	ConstraintSuite
}

func (suite *BiggerThanEqualConstraintSuite) TestBiggerThanEqualConstraint() {
	cases := []ConstraintTestCase{
		{
			Name: "successfully create constraint and evaluate using entity id",
			Constraint: model.Constraint{
				Name: constraint.BiggerThanEqualConstraintName,
				Parameters: json.RawMessage(
					`{"value": 10}`,
				),
			},
			ErrExpected: false,
			Entity: model.Entity{
				ID: 11,
			},
			EvaluateExpected: true,
		},
		{
			Name: "successfully create constraint and evaluate using entity type",
			Constraint: model.Constraint{
				Name: constraint.BiggerThanEqualConstraintName,
				Parameters: json.RawMessage(
					fmt.Sprintf(`{"value": 10, "property": "%s"}`, constraint.EntityTypeProperty),
				),
			},
			ErrExpected: false,
			Entity: model.Entity{
				ID:   8,
				Type: "11",
			},
			EvaluateExpected: true,
		},
		{
			Name: "successfully create constraint and evaluate using entity context",
			Constraint: model.Constraint{
				Name: constraint.BiggerThanEqualConstraintName,
				Parameters: json.RawMessage(
					`{"value": 10, "property": "test"}`,
				),
			},
			ErrExpected: false,
			Entity: model.Entity{
				ID:      8,
				Type:    "t",
				Context: map[string]string{"test": "11"},
			},
			EvaluateExpected: true,
		},
		{
			Name: "successfully create constraint and evaluate using zero value",
			Constraint: model.Constraint{
				Name: constraint.BiggerThanEqualConstraintName,
				Parameters: json.RawMessage(
					`{"value": 0, "property": "test"}`,
				),
			},
			ErrExpected: false,
			Entity: model.Entity{
				ID:      8,
				Type:    "t",
				Context: map[string]string{"test": "1"},
			},
			EvaluateExpected: true,
		},
		{
			Name: "successfully create constraint and evaluate using invalid property",
			Constraint: model.Constraint{
				Name: constraint.BiggerThanEqualConstraintName,
				Parameters: json.RawMessage(
					`{"value": 0, "property": "test"}`,
				),
			},
			ErrExpected: false,
			Entity: model.Entity{
				ID:      8,
				Type:    "t",
				Context: map[string]string{"test": "t"},
			},
			EvaluateExpected: false,
		},
		{
			Name: "successfully create constraint and evaluate using negative property",
			Constraint: model.Constraint{
				Name: constraint.BiggerThanEqualConstraintName,
				Parameters: json.RawMessage(
					`{"value": 0, "property": "test"}`,
				),
			},
			ErrExpected: false,
			Entity: model.Entity{
				ID:      8,
				Type:    "t",
				Context: map[string]string{"test": "-1"},
			},
			EvaluateExpected: false,
		},
		{
			Name: "successfully create constraint and evaluate using entity id equality",
			Constraint: model.Constraint{
				Name: constraint.BiggerThanEqualConstraintName,
				Parameters: json.RawMessage(
					`{"value": 1}`,
				),
			},
			ErrExpected: false,
			Entity: model.Entity{
				ID: 1,
			},
			EvaluateExpected: true,
		},
		{
			Name: "successfully create constraint and evaluate using entity type equality",
			Constraint: model.Constraint{
				Name: constraint.BiggerThanEqualConstraintName,
				Parameters: json.RawMessage(
					fmt.Sprintf(`{"value": 10, "property": "%s"}`, constraint.EntityTypeProperty),
				),
			},
			ErrExpected: false,
			Entity: model.Entity{
				ID:   1,
				Type: "10",
			},
			EvaluateExpected: true,
		},
		{
			Name: "successfully create constraint and evaluate using entity context equality",
			Constraint: model.Constraint{
				Name: constraint.BiggerThanEqualConstraintName,
				Parameters: json.RawMessage(
					`{"value": 10, "property": "test"}`,
				),
			},
			ErrExpected: false,
			Entity: model.Entity{
				ID:      1,
				Type:    "t",
				Context: map[string]string{"test": "10"},
			},
			EvaluateExpected: true,
		},
	}

	suite.RunCases(cases)
}

func TestBiggerThanEqualConstraintSuite(t *testing.T) {
	suite.Run(t, new(BiggerThanEqualConstraintSuite))
}