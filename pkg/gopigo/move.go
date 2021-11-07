package gopigo

import (
	"math"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gobot.io/x/gobot/platforms/dexter/gopigo3"
)

type Motor struct {
	driver *gopigo3.Driver
	debug  bool
}

// NewMotor creates a new gopigo Motor
func NewMotor(driver *gopigo3.Driver, debug bool) *Motor {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	return &Motor{
		driver: driver,
		debug:  debug,
	}
}

// Forward moves the robot forwards
func (m *Motor) Forward(speed int) error {
	log.Debug().Msgf("moving forwards at speed %s", speed)

	err := m.driver.SetMotorDps(gopigo3.MOTOR_LEFT, speed)
	if err != nil {
		return err
	}
	err = m.driver.SetMotorDps(gopigo3.MOTOR_RIGHT, speed)
	if err != nil {
		return err
	}

	return nil
}

// Left moves the robot to the left
func (m *Motor) Left(speed int) error {
	log.Debug().Str("func", "left").Msgf("turning left at speed %s", speed)
	err := m.driver.SetMotorDps(gopigo3.MOTOR_RIGHT, speed)
	if err != nil {
		return err
	}

	return nil
}

// Right moves the robot to the right
func (m *Motor) Right(speed int) error {
	s := int(math.Abs(float64(speed)))

	log.Debug().Str("func", "right").Msgf("turning right at speed %s", s)
	err := m.driver.SetMotorDps(gopigo3.MOTOR_RIGHT, s)
	if err != nil {
		return err
	}

	return nil
}

// Backwards moves backwards at a default speed
func (m *Motor) Backwards(speed uint) error {
	s := int(speed) * -1
	log.Debug().Str("func", "backwards").Msgf("moving backwards at speed %s", s)
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
	log.Debug().Str("func", "stop").Msgf("stopping motor by setting rotations to %s", 0)
	err := m.driver.SetMotorDps(gopigo3.MOTOR_LEFT, 0)
	if err != nil {
		return err
	}

	err = m.driver.SetMotorDps(gopigo3.MOTOR_RIGHT, 0)
	if err != nil {
		return err
	}

	return nil
}
