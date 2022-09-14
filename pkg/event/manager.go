package event

import (
	"mockers/pkg/errors"
	"mockers/pkg/utils"
	"sync"
	"time"
)

const AllAction = "_all"

type Hock func(action, origin, info string, label map[string]string, createTime int64, arrivalTime int64, eventHash string)

type Manager struct {
	Channel chan *Model

	ActionHock map[string][]Hock
}

func (m *Manager) RegisterHock(h Hock, action string) {
	if _, ok := m.ActionHock[action]; !ok {
		m.ActionHock[action] = []Hock{}
	}
	m.ActionHock[action] = append(m.ActionHock[action], h)
}

func (m *Manager) Send(model *Model) error {
	if model == nil {
		return errors.ValueUnExpectValue.Param("PrtValue of {Model}", "nil(ptr)")
	}
	m.Channel <- model
	return nil
}

func (m *Manager) spread(event *Model, arrivalTime int64) {
	var allListener []Hock

	if event == nil {
		utils.Logger.Error(errors.ValueUnExpectValue.Param("PrtValue of {Model}", "nil(ptr)").Error())
		return
	}

	if actions, ok := m.ActionHock[AllAction]; ok {
		allListener = append(allListener, actions...)
	}

	if actions, ok := m.ActionHock[event.Action]; ok {
		allListener = append(allListener, actions...)
	}

	hockWaitGroup := sync.WaitGroup{}
	hockWaitGroup.Add(len(allListener))
	for _, hock := range allListener {
		// to missing some wrong-things.
		hock := hock
		go func() {
			hock(event.Action, event.Origin, event.Info, event.Labels.Map(), event.CreateTime, arrivalTime, event.Hash())
			hockWaitGroup.Done()
		}()
	}

	hockWaitGroup.Wait()
}

func (m *Manager) Listen() {
	for {
		select {
		case v := <-m.Channel:
			if v == nil {
				continue
			}
			t := time.Now().Unix()

			utils.Logger.Infof("Receive an event: %s", v.Log())

			go m.spread(v, t)
		}
	}
}
