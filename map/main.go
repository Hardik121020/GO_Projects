package main

func main() {
	//var colours map[string]string

	colours := make(map[string]string)

	//colours := map[string]string{
	//	"red":   "C1",
	//	"green": "C2",
	//}
	colours["Red"] = "C1"
	colours["Blue"] = "C2"
	colours["Green"] = "C3"
	mapPrint(colours)
	//delete(colours, "Red")
}
