package main

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (n directMessage) importance() int {
	var importanceScore int
	if n.isUrgent {
		importanceScore = 50
	} else {
		importanceScore = n.priorityLevel
	}
	return importanceScore
}
func (n groupMessage) importance() int {
	return n.priorityLevel
}
func (n systemAlert) importance() int {
	return 100
}

func processNotification(n notification) (string, int) {
	switch n := n.(type) {
	case directMessage:
		return n.senderUsername, n.importance()
	case groupMessage:
		return n.groupName, n.importance()
	case systemAlert:
		return n.alertCode, n.importance()
	default:
		return "", 0
	}
}
