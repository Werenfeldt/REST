package Server

type Teacher struct {
	Id int64 `json:"id,omitempty"`

	//Course *Course `json:"Course,omitempty"`

	Firstname string `json:"firstname,omitempty"`

	Lastname string `json:"lastname,omitempty"`
}

var teacher1 = Teacher{Id: 11, Firstname: "Hanne", Lastname: "Hansen"}
var teacher2 = Teacher{Id: 33, Firstname: "Børge", Lastname: "Bentsen"}
var teacher3 = Teacher{Id: 55, Firstname: "Roberto", Lastname: "Robertsen"}
