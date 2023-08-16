package main

var LeaderboardKey = "leaderboard"

type Leaderboard struct {
	Count int `json:"count"`
	Users []*User
}

func (db *Database) GetLeaderboard() (*Leaderboard, error) {
	// choose specified range
	scores := db.Client.ZRangeWithScores(Ctx, LeaderboardKey, 0, -1) // 0~-1 shows all members
	if scores == nil {
		return nil, ErrNil
	}

	count := len(scores.Val())
	users := make([]*User, count)

	for idx, member := range scores.Val() {
		users[idx] = &User{
			Username: member.Member.(string),
			Points:   int(member.Score),
			Rank:     idx,
		}
	}

	leaderboard := &Leaderboard{
		Count: count,
		Users: users,
	}

	return leaderboard, nil
}
