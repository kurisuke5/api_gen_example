// Package props is a scaffold file for props of controllers
package props

import (
	"github.com/kurisuke5/api_gen_example/backend/domain/health"
)

// ControllerProps is passed from Bootstrap() to all controllers
type ControllerProps struct {
	HealthService *health.Service
}
