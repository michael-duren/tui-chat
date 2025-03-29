package controllers

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/michael-duren/tui-chat/ui/models"
)

func Login(m *models.AppModel, msg tea.Msg) (*models.AppModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "s", "enter":
			m.Logger.Debug("In case")
			return m, nil
		default:
			m.Logger.Infof("In default : %v\n", msg)
		}
	}

	var cmds []tea.Cmd

	form, cmd := m.Login.Form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.Login.Form = f
		cmds = append(cmds, cmd)
	}

	if m.Login.Form.State == huh.StateCompleted {
		// TODO: Update with actual logic
		m.Logger.Info("in state completed")
		m.Logger.Infof("Form values - Address: %s, Username: %s, Secret: %s",
			m.Login.Address,
			m.Login.Username,
			m.Login.Secret)
		m.CurrentView = models.Loading
		// Initialize the loading view
		return m, m.Loading.Init()
	}

	return m, tea.Batch(cmds...)
}
