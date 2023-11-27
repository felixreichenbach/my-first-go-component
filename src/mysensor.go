// Package mybase implements a base that only supports SetPower (basic forward/back/turn controls), IsMoving (check if in motion), and Stop (stop all motion).
// It extends the built-in resource subtype Base and implements methods to handle resource construction, attribute configuration, and reconfiguration.

package main

import (
	"context"

	"github.com/pkg/errors"

	"go.viam.com/rdk/components/sensor"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/module"
	"go.viam.com/rdk/resource"
	"go.viam.com/utils"
)

// Here is where we define your new model's colon-delimited-triplet (acme:demo:mybase)
// acme = namespace, demo = repo-name, mybase = model name.
var (
	Model            = resource.NewModel("viam-soleng", "sensor", "mysensor")
	errUnimplemented = errors.New("unimplemented")
)

type mySensor struct {
	resource.Named
	logger logging.Logger
}

func init() {
	resource.RegisterComponent(sensor.API, Model, resource.Registration[sensor.Sensor, *Config]{
		Constructor: newSensor,
	})
}

func newSensor(ctx context.Context, deps resource.Dependencies, conf resource.Config, logger logging.Logger) (sensor.Sensor, error) {
	b := &mySensor{
		Named:  conf.ResourceName().AsNamed(),
		logger: logger,
	}
	if err := b.Reconfigure(ctx, deps, conf); err != nil {
		return nil, err
	}
	return b, nil
}

// Config contains two component (motor) names.
type Config struct {
	//LeftMotor  string `json:"motorL"`
	//RightMotor string `json:"motorR"`
}

// Validate validates the config and returns implicit dependencies,
func (cfg *Config) Validate(path string) ([]string, error) {
	return []string{}, nil
}

// Reconfigure reconfigures with new settings.
func (b *mySensor) Reconfigure(ctx context.Context, deps resource.Dependencies, conf resource.Config) error {

	// This takes the generic resource.Config passed down from the parent and converts it to the
	// model-specific (aka "native") Config structure defined, above making it easier to directly access attributes.

	/*
		sensorConfig, err := resource.NativeConfig[*Config](conf)
		if err != nil {
			return err
		}
	*/
	return nil
}

// Get sensor reading
func (s *mySensor) Readings(ctx context.Context, _ map[string]interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{"hello": "world"}, nil
}

// DoCommand simply echos whatever was sent.
func (b *mySensor) DoCommand(ctx context.Context, cmd map[string]interface{}) (map[string]interface{}, error) {
	return nil, errUnimplemented
}

// Close function which is exectuted when the component is shut down
func (b *mySensor) Close(ctx context.Context) error {
	return nil
}

func main() {
	// NewLoggerFromArgs will create a logging.Logger at "DebugLevel" if
	// "--log-level=debug" is an argument in os.Args and at "InfoLevel" otherwise.
	utils.ContextualMain(mainWithArgs, module.NewLoggerFromArgs("yourmodule"))
}

func mainWithArgs(ctx context.Context, args []string, logger logging.Logger) (err error) {
	myMod, err := module.NewModuleFromArgs(ctx, logger)
	if err != nil {
		return err
	}

	// Models and APIs add helpers to the registry during their init().
	// They can then be added to the module here.
	err = myMod.AddModelFromRegistry(ctx, sensor.API, Model)
	if err != nil {
		return err
	}

	err = myMod.Start(ctx)
	defer myMod.Close(ctx)
	if err != nil {
		return err
	}
	<-ctx.Done()
	return nil
}
