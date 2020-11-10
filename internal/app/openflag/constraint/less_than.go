package constraint

import (
	"strconv"

	"github.com/OpenFlag/OpenFlag/internal/app/openflag/model"
	"github.com/sirupsen/logrus"
)

// LessThanConstraint represents Openflag less than constraint.
type LessThanConstraint struct {
	Value    float64 `json:"value"`
	Property string  `json:"property,omitempty"`
}

// Name is an implementation for the Constraint interface.
func (l LessThanConstraint) Name() string {
	return LessThanConstraintName
}

// Validate is an implementation for the Constraint interface.
func (l LessThanConstraint) Validate() error {
	return nil
}

// Initialize is an implementation for the Constraint interface.
func (l *LessThanConstraint) Initialize() error {
	return nil
}

// Evaluate is an implementation for the Constraint interface.
func (l LessThanConstraint) Evaluate(e model.Entity) bool {
	if l.Property == "" {
		return float64(e.ID) < l.Value
	}

	var property string

	if l.Property == EntityTypeProperty {
		property = e.Type
	} else {
		var ok bool

		property, ok = e.Context[l.Property]
		if !ok {
			return false
		}
	}

	value, err := strconv.ParseFloat(property, 64)
	if err != nil {
		logrus.Errorf(
			"invalid property for less than constraint => property: %s, value: %s, err: %s",
			l.Property, property, err.Error(),
		)

		return false
	}

	return value < l.Value
}
