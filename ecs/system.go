package ecs

// System specifies a system.
type System interface {
	Init(mngr *SystemManager)
	Update(dt float64)
	HandleMessage(msg Message, data interface{}) interface{}
}

// Priority implies that some things are more important than others.
type Priority interface {
	Priority() uint
}

// byPriority helps to sort the systems by priority
type byPriority []System

// Len does something (sort function).
func (s byPriority) Len() int {
	return len(s)
}

// Less does something (sort function).
func (s byPriority) Less(i, j int) bool {
	var prio1, prio2 uint = 50, 50

	if prior1, ok := s[i].(Priority); ok {
		prio1 = prior1.Priority()
	}
	if prior2, ok := s[j].(Priority); ok {
		prio2 = prior2.Priority()
	}

	return prio1 > prio2
}

// Swap does something (sort function).
func (s byPriority) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
