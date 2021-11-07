package gopigo

import "gobot.io/x/gobot/platforms/dexter/gopigo3"

type Motor struct {
	driver *gopigo3.Driver
	motor  gopigo3.Motor
}

// NewMotor creates a new gopigo Motor
func NewMotor(driver *gopigo3.Driver, motor gopigo3.Motor) *Motor {
	return &Motor{
		driver: driver,
		motor:  motor,
	}
}

// Forward moves forwards at a default speed
func (m *Motor) Forward() error {
	defaultSpeed := 360
	err := m.driver.SetMotorDps(gopigo3.MOTOR_LEFT, defaultSpeed)
	if err != nil {
		return err
	}

	err = m.driver.SetMotorDps(gopigo3.MOTOR_RIGHT, defaultSpeed)
	if err != nil {
		return err
	}

	return nil
}

// Stop sets the motor dps to zero
func (m *Motor) Stop() error {
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
