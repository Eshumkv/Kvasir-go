package ecs

// System specifies a system.
type System interface {
	Init(mngr *SystemManager)
	Add(e *Entity)
	Update(dt float64)
	Delete(e Entity)
}

// Priority implies that some things are more important than others.
type Priority interface {
	Priority() uint
}

type byPriority []System

func (s byPriority) Len() int {
	return len(s)
}

func (s byPriority) Less(i, j int) bool {
	var prio1, prio2 uint

	if prior1, ok := s[i].(Priority); ok {
		prio1 = prior1.Priority()
	}
	if prior2, ok := s[j].(Priority); ok {
		prio2 = prior2.Priority()
	}

	return prio1 > prio2
}

func (s byPriority) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
