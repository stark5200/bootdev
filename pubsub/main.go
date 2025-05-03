package main

func (u user) march(p piece, publishCh chan<- move) {
	newMove := move {
		userName: u.name, 
		piece: p,
	}
	publishCh <- newMove
}

// don't touch below this line

type user struct {
	name   string
	pieces []piece
}

type move struct {
	userName string
	piece    piece
}

type piece struct {
	location string
	name     string
}

func doBattles(publishCh <-chan move, users []user) []piece {
	fights := []piece{}
	for mv := range publishCh {
		for _, u := range users {
			if u.name == mv.userName {
				continue
			}
			for _, piece := range u.pieces {
				if piece.location == mv.piece.location {
					fights = append(fights, piece)
				}
			}
		}
	}
	return fights
}

// battles on user
func (u user) doBattles(subCh <-chan move) []piece {
	pieces_in_battle := []piece{}
	for mv := range subCh {
		for _, p := range u.pieces {
			if p.location == mv.piece.location {
				pieces_in_battle = append(pieces_in_battle, p)
			}
		}
	}
	return pieces_in_battle
}

func distributeBattles(publishCh <-chan move, subChans []chan move) {
	for mv := range publishCh {
		for _, subCh := range subChans {
			subCh <- mv
		}
	}
}
