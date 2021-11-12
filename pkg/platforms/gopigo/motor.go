package gopigo

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dexter/gopigo3"
)

type Motor struct {
	gobot.Eventer
	*gopigo3.Driver
}

type MotorState struct {
	Power uint16
	Speed int
}

// NewMotor creates a new gopigo Motor
func NewMotor(driver *gopigo3.Driver) *Motor {
	m := &Motor{
		Eventer: gobot.NewEventer(),
		Driver:  driver,
	}

	m.AddEvent(MOTOR_STOP)
	m.AddEvent(MOTOR_BACKWARD)
	m.AddEvent(MOTOR_FORWARD)
	m.AddEvent(MOTOR_LEFT)
	m.AddEvent(MOTOR_RIGHT)

	return m
}

func (m *Motor) State() map[string]MotorState {
	_, lpower, _, ldps, _ := m.GetMotorStatus(gopigo3.MOTOR_LEFT)
	leftmotor := MotorState{
		Power: lpower,
		Speed: ldps,
	}

	_, rpower, _, rdps, _ := m.GetMotorStatus(gopigo3.MOTOR_LEFT)
	rightmotor := MotorState{
		Power: rpower,
		Speed: rdps,
	}

	return map[string]MotorState{
		"left":  leftmotor,
		"right": rightmotor,
	}
}

// Forward moves the robot forwards
func (m *Motor) Forward(speed int) error {
	m.reset()

	err := m.SetMotorDps(gopigo3.MOTOR_LEFT, speed)
	if err != nil {
		return err
	}

	err = m.SetMotorDps(gopigo3.MOTOR_RIGHT, speed)
	if err != nil {
		m.Publish(m.Event(MOTOR_ERR), speed)
		return err
	}

	m.Publish(m.Event(MOTOR_FORWARD), speed)

	return nil
}

// Left moves the robot to the left
func (m *Motor) Left(speed int) error {
	m.reset()

	err := m.SetMotorDps(gopigo3.MOTOR_RIGHT, speed)
	if err != nil {
		return err
	}
	m.Publish(m.Event(MOTOR_LEFT), speed)

	return nil
}

// Right moves the robot to the right
func (m *Motor) Right(speed int) error {
	m.reset()

	err := m.SetMotorDps(gopigo3.MOTOR_LEFT, speed)
	if err != nil {
		return err
	}
	m.Publish(m.Event(MOTOR_RIGHT), speed)

	return nil
}

// Backwards moves backwards at a default speed
func (m *Motor) Backward(speed uint) error {
	m.reset()

	s := int(speed) * -1
	err := m.SetMotorDps(gopigo3.MOTOR_LEFT, s)
	if err != nil {
		return err
	}

	err = m.SetMotorDps(gopigo3.MOTOR_RIGHT, int(speed)*-1)
	if err != nil {
		return err
	}

	m.Publish(m.Event(MOTOR_BACKWARD), speed)

	return nil
}

// Stop sets the motor dps to zero
func (m *Motor) Stop() error {
	m.reset()

	speed := 0
	err := m.SetMotorDps(gopigo3.MOTOR_LEFT, speed)
	if err != nil {
		return err
	}

	err = m.SetMotorDps(gopigo3.MOTOR_RIGHT, speed)
	if err != nil {
		return err
	}

	m.Publish(m.Event(MOTOR_STOP), speed)

	return nil
}

func (m *Motor) reset() error {
	speed := 0
	err := m.SetMotorDps(gopigo3.MOTOR_LEFT, speed)
	if err != nil {
		return err
	}
	err = m.SetMotorDps(gopigo3.MOTOR_RIGHT, speed)
	if err != nil {
		return err
	}

	m.Publish(m.Event(MOTOR_STOP), speed)

	return nil
}
