package story

import (
	"fmt"
)

// each action detail
type StoryAction struct {
	Text           string
	MoveTo string
	Condition      StoryCondition

	Valid          *StoryAction
	Invalid        *StoryAction
}

// Run the action: set/test condition, move to another position or whatever
func (storyAction StoryAction) Run() string {
	if storyAction.Text != "" {
		fmt.Println(storyAction.Text)
	}

	// do we have a condition ?
	if storyAction.hasCondition() {
		storyAction.Condition.Init()

		if storyAction.Condition.IsTestable() {
			if storyAction.Condition.Verify() {
				fmt.Println()
				return storyAction.Valid.Run()
			} else {
				return storyAction.Invalid.Run()
			}
		}
	}

	// do we need to move to another position ?
	if storyAction.hasMoveTo() {
		return storyAction.MoveTo
	}

	return ""
}

// does the action have a condition
func (storyAction StoryAction) hasCondition() bool {
	if (storyAction.Condition != StoryCondition{}) {
		return true
	}

	return false
}

// does the action have a condition
func (storyAction StoryAction) hasMoveTo() bool {
	if (storyAction.MoveTo != "") {
		return true
	}

	return false
}
