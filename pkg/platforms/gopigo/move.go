package gopigo

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gobot.io/x/gobot/platforms/dexter/gopigo3"
)

type Motor struct {
	driver *gopigo3.Driver
	logger zerolog.Logger
}

type MotorOption func(*Motor)

// NewMotor creates a new gopigo Motor
func NewMotor(driver *gopigo3.Driver, options ...MotorOption) *Motor {
	var (
		defaultLogger = zerolog.New(os.Stderr)
	)

	m := &Motor{
		driver: driver,
		logger: defaultLogger,
	}
	for _, option := range options {
		option(m)
	}
	return m
}

func (m *Motor) WithDebug() MotorOption {
	return func(m *Motor) {
		m.logger = m.logger.Level(zerolog.DebugLevel)
	}
}

func (m *Motor) debug(motor gopigo3.Motor) map[string]interface{} {
	_, power, _, dps, _ := m.driver.GetMotorStatus(motor)
	fields := map[string]interface{}{
		"motor": motor,
		"power": power,
		"dps":   dps,
	}
	log.Debug().Fields(fields).Send()
	return fields
}

// Forward moves the robot forwards
func (m *Motor) Forward(speed int) error {
	m.debug(gopigo3.MOTOR_LEFT)
	err := m.driver.SetMotorDps(gopigo3.MOTOR_LEFT, speed)
	if err != nil {
		return err
	}

	m.debug(gopigo3.MOTOR_RIGHT)
	err = m.driver.SetMotorDps(gopigo3.MOTOR_RIGHT, speed)
	if err != nil {
		return err
	}

	return nil
}

// Left moves the robot to the left
func (m *Motor) Left(speed int) error {
	m.debug(gopigo3.MOTOR_LEFT)
	err := m.driver.SetMotorDps(gopigo3.MOTOR_RIGHT, speed)
	if err != nil {
		return err
	}

	return nil
}

// Right moves the robot to the right
func (m *Motor) Right(speed int) error {
	m.debug(gopigo3.MOTOR_RIGHT)
	err := m.driver.SetMotorDps(gopigo3.MOTOR_LEFT, speed)
	if err != nil {
		return err
	}

	return nil
}

// Backwards moves backwards at a default speed
func (m *Motor) Backward(speed uint) error {
	s := int(speed) * -1
	m.debug(gopigo3.MOTOR_RIGHT)
	err := m.driver.SetMotorDps(gopigo3.MOTOR_LEFT, s)
	if err != nil {
		return err
	}

	err = m.driver.SetMotorDps(gopigo3.MOTOR_RIGHT, int(speed)*-1)
	if err != nil {
		return err
	}

	return nil
}

// Stop sets the motor dps to zero
func (m *Motor) Stop() error {
	m.debug(gopigo3.MOTOR_LEFT)
	err := m.driver.SetMotorDps(gopigo3.MOTOR_LEFT, 0)
	if err != nil {
		return err
	}

	m.debug(gopigo3.MOTOR_RIGHT)
	err = m.driver.SetMotorDps(gopigo3.MOTOR_RIGHT, 0)
	if err != nil {
		return err
	}

	return nil
}
