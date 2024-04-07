package handler

type Task struct {
	ID    int
	Title string
	Done  bool
}

var tasks = []Task{
	{1, "bla", false},
	{2, "ble", true},
}
