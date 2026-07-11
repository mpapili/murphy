package bus

func ServerState(gameID string) string      { return "budowac.server." + gameID + ".state" }
func MurphyOccurrence(gameID string) string { return "budowac.murphy." + gameID + ".occurrence" }
