package tui

import (
	"fmt"
	// "os"

	utils "github.com/korbexmachina/matrix-tasks/utils"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
        UrgentImportant []utils.Task // tasks to display
        UrgentNotImportant []utils.Task // tasks to display
        NotUrgentImportant []utils.Task // tasks to display
        NotUrgentNotImportant []utils.Task // tasks to display
	cursor int // which task is pointed to
	selected map[int]struct{} // which tasks are selected
}

// Initializes the model
func InitialModel(path string) model {
    urgentImportant, err := utils.GetBucketTasks("UrgentImportant")
    if err != nil {
        panic(err)
    }
    urgentNotImportant, err := utils.GetBucketTasks("UrgentNotImportant")
    if err != nil {
        panic(err)
    }
    notUrgentImportant , err := utils.GetBucketTasks("NotUrgentImportant")
    if err != nil {
        panic(err)
    }
    notUrgentNotImportant, err := utils.GetBucketTasks("NotUrgentNotImportant")
    if err != nil {
        panic(err)
    }
    return model{
        UrgentImportant: urgentImportant,
        UrgentNotImportant: urgentNotImportant,
        NotUrgentImportant: notUrgentImportant,
        NotUrgentNotImportant: notUrgentNotImportant,
        cursor: 0,
        selected: make(map[int]struct{}),
        }
    }

        // Init initializes the program
        func (m model) Init() tea.Cmd {
        // Just return `nil`, which means "no I/O right now, please."
        return nil
    }

        func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    // Is it a key press?
    case tea.KeyMsg:

    // Cool, what was the actual key pressed?
    switch msg.String() {

    // These keys should exit the program.
    case "ctrl+c", "q":
    return m, tea.Quit

    // The "up" and "k" keys move the cursor up
    case "up", "k":
    if m.cursor > 0 {
    m.cursor--
}

// The "down" and "j" keys move the cursor down
case "down", "j":
if m.cursor < len(m.tasks)-1 {
    m.cursor++
}

// The "enter" key and the spacebar (a literal space) toggle
// the selected state for the item that the cursor is pointing at.
case "enter", " ":
_, ok := m.selected[m.cursor]
if ok {
    delete(m.selected, m.cursor)
} else {
    m.selected[m.cursor] = struct{}{}
}
        }	
    }

// Return the updated model to the Bubble Tea runtime for processing.
// Note that we're not returning a command.
return m, nil
}

func (m model) View() string {
    // The header
    s := "What should we buy at the market?\n\n"

    // Iterate over our choices
    for i, choice := range m.tasks {

        // Is the cursor pointing at this choice?
        cursor := " " // no cursor
        if m.cursor == i {
            cursor = ">" // cursor!
        }

        // Is this choice selected?
        checked := " " // not selected
        if _, ok := m.selected[i]; ok {
            checked = "x" // selected!
        }

        // Render the row
        s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
    }

    // The footer
    s += "\nPress q to quit.\n"

    // Send the UI for rendering
    return s
}
