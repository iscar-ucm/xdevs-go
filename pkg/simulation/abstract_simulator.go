/*
 * Copyright (c) 2021, Román Cárdenas Rodríguez.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package simulation

import "github.com/iscar-ucm/xdevs-go/pkg/modeling"

type AbstractSimulator interface {
	Initialize()                  // performs all the required operations before starting a simulation.
	Exit()                        // performs all the required operations to exit after a simulation.
	TA() float64                  // time advance function.
	Collect()                     // collection function.
	Transition()                  // transition function.
	Clear()                       // clear function.
	GetModel() modeling.Component // returns the Component attached to the simulator.
	GetTL() float64               // returns simulator's last simulation time.
	SetTL(tl float64)             // sets simulator's last simulation time.
	GetTN() float64               // returns simulator's next timeout.
	SetTN(tn float64)             // sets simulator's next timeout.
	GetClock() Clock              // returns abstract simulator's simulation Clock.
}

func NewAbstractSimulator(clock Clock) AbstractSimulator {
	s := abstractSimulator{clock, 0, 0}
	return &s
}

type abstractSimulator struct {
	clock Clock   // Simulation Clock.
	tL    float64 // Time of last event
	tN    float64 // Time of next event
}

// Initialize performs all the required operations before starting a simulation.
func (a *abstractSimulator) Initialize() {
	panic("implement me")
}

// Exit performs all the required operations to exit after a simulation.
func (a *abstractSimulator) Exit() {
	panic("implement me")
}

// TA is the time advance function. It returns the elapsed time for the next timeout.
func (a *abstractSimulator) TA() float64 {
	panic("implement me")
}

// Collect is the collection function. It is invoked on imminent simulators before the transition function.
func (a *abstractSimulator) Collect() {
	panic("implement me")
}

// Transition is the transition function. It deals with internal and external transitions.
func (a *abstractSimulator) Transition() {
	panic("implement me")
}

// Clear performs all the required operation to clear the simulation state.
func (a *abstractSimulator) Clear() {
	panic("implement me")
}

// GetModel returns simulator's DEVS component.
func (a *abstractSimulator) GetModel() modeling.Component {
	panic("implement me")
}

// GetTL returns simulator's last simulation time.
func (a *abstractSimulator) GetTL() float64 {
	return a.tL
}

// SetTL sets simulator's last simulation time.
func (a *abstractSimulator) SetTL(tL float64) {
	a.tL = tL
}

// GetTN returns simulator's next timeout.
func (a *abstractSimulator) GetTN() float64 {
	return a.tN
}

// SetTN sets simulator's next timeout.
func (a *abstractSimulator) SetTN(tN float64) {
	a.tN = tN
}

// GetClock returns abstract simulator's simulation Clock.
func (a *abstractSimulator) GetClock() Clock {
	return a.clock
}
