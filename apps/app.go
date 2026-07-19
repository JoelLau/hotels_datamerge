package apps

import "context"

type App interface {
	Run(context.Context) error
}
