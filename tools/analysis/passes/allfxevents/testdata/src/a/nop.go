package a

import "github.com/openumi/fx/fxevent"

type nopLogger struct{}

func (nopLogger) LogEvent(fxevent.Event) {
	// Don't do anything with the event. Should not cause a
	// diagnostic.
}
