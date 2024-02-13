package main

const HOME_TEAM_WON = 1

// round robin competition
// no ties
// win get 3 point

// 1. initiate value
// 1.1 teamMap map[string]int
// 1.2 winner string
// 2. loop competitions
// 2.1 get result[I]
// 2.2 if result = 0, teamMap[competitions[i][1] += 1
// 2.3 else teamMap[competitions[i][0] += 1
// 2.4 get teamMap[winner]
// 2.4.1 if teamMap[winner] < teamMap[competitions[i][0/1]
//
//	then winner = competitions[i][0/1]
//
// 3. return winner
func TournamentWinner(competitions [][]string, results []int) string {
	teamMap := make(map[string]int)
	winner := ""

	for idx, v := range competitions {
		winRsult := results[idx]
		currentTeamWin := ""
		if winRsult == HOME_TEAM_WON {
			teamMap[v[0]] += 3
			currentTeamWin = v[0]
		} else {
			teamMap[v[1]] += 3
			currentTeamWin = v[1]
		}

		if winner == "" || teamMap[winner] < teamMap[currentTeamWin] {
			winner = currentTeamWin
		}
	}

	return winner
}
