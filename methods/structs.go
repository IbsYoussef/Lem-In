package methods

type LemIn struct {
	Ants  int
	Rooms []Room
	Start Room
	End   Room
	Links []Link
	Path  []Room
}
type Link struct {
	Room1 Room
	Room2 Room
}
type Room struct {
	Name      string
	X         int
	Y         int
	AntNumber int
}
