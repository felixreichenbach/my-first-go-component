// This module implements a simple sensor returning the value of the setting attribute in the sensor's configuration.
// It extends the built-in resource subtype Sensor and implements its methods to handle resource construction, attribute configuration, and reconfiguration.

package main

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"go.viam.com/rdk/components/sensor"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/module"
	"go.viam.com/rdk/resource"
	"go.viam.com/utils"
)

// Module initializer which registers this component with the Viam RDK register.
func init() {
	resource.RegisterComponent(sensor.API, Model, resource.Registration[sensor.Sensor, *Config]{
		Constructor: newSensor,
	})
}

// Here you define your new model's colon-delimited-triplet (acme:demo:mybase). If you plan to upload this module to the Viam registry, "acme" must match your registry namespace.
// acme = namespace, demo = repo-name, mybase = model name.
var (
	Model            = resource.NewModel("viam-soleng", "sensor", "mysensor")
	errUnimplemented = errors.New("unimplemented")
)

// The sensor struct with the value attribute which will be returned by the sensor "readings" method.
type mySensor struct {
	resource.Named
	logger logging.Logger
	value  int
}

// Sensor constructor
func newSensor(ctx context.Context, deps resource.Dependencies, conf resource.Config, logger logging.Logger) (sensor.Sensor, error) {
	s := &mySensor{
		Named:  conf.ResourceName().AsNamed(),
		logger: logger,
	}
	if err := s.Reconfigure(ctx, deps, conf); err != nil {
		return nil, err
	}
	return s, nil
}

// Maps JSON component configuration attributes.
type Config struct {
	Setting int `json:"setting"`
}

// Implement component configuration validation and and return implicit dependencies.
func (cfg *Config) Validate(path string) ([]string, error) {

	if cfg.Setting == 0 {
		return nil, fmt.Errorf(`expected "setting" attribute bigger than 0 for mysensor %q`, path)
	}
	return []string{}, nil
}

// Reconfigure reconfigures with new settings.
func (s *mySensor) Reconfigure(ctx context.Context, deps resource.Dependencies, conf resource.Config) error {

	// This takes the generic resource.Config passed down from the parent and converts it to the
	// model-specific (aka "native") Config structure defined, above making it easier to directly access attributes.
	sensorConfig, err := resource.NativeConfig[*Config](conf)
	if err != nil {
		return err
	}
	s.logger.Infof(`Reconfiguring sensor value from %v to %v`, s.value, sensorConfig.Setting)
	s.value = sensorConfig.Setting
	return err
}

// Get sensor reading
func (s *mySensor) Readings(ctx context.Context, _ map[string]interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{"setting": s.value}, nil
}

// DoCommand can be implemented to extend sensor functionality but returns unimplemented in this example.
func (s *mySensor) DoCommand(ctx context.Context, cmd map[string]interface{}) (map[string]interface{}, error) {
	return nil, errUnimplemented
}

// The close method is executed when the component is shut down
func (s *mySensor) Close(ctx context.Context) error {
	return nil
}

func main() {
	// NewLoggerFromArgs will create a logging.Logger at "DebugLevel" if
	// "--log-level=debug" is an argument in os.Args and at "InfoLevel" otherwise.
	utils.ContextualMain(mainWithArgs, module.NewLoggerFromArgs("My Go Sensor Module"))
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
