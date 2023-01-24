package main

import (
	"github.com/siasalar/alien-invasion/invasion"
)

const (
	mapFilePath         = "map.txt"
	alienMovementThresh = 10000
)

func main() {
	numAliens, err := invasion.GetNumberOfAliens()
	if err != nil {
		panic(err)
	}

	cityMap, err := invasion.ReadWorldMapFile(mapFilePath)
	if err != nil {
		panic(err)
	}

	invasion.Run(cityMap, numAliens, alienMovementThresh)

	// file, err := os.Open("file.txt")
	//fmt.print(file)

	// Read the contents of the file into a byte slice
	// data, err := ioutil.ReadFile("example.txt")
	//if err != nil {
	//	fmt.Println("Error reading file:", err)
	//	return
	//}
	//
	//// Convert the byte slice to a string
	//contents := string(data)
	//
	//// Print the contents of the file
	//fmt.Println(contents)
}

// Foo north=Bar west=Baz south=Qu-ux
//Bar south=Foo west=Bee
//Baz east=Foo north=Qux
//Qu-ux north=Foo east=Quux
//Bee west=Bar south=Qux
//Qux south=Bee east=Quux
//Quux west=Qux
//
//City1 north=City2 south=City3 east=City4
//City2 south=City1 west=City5 east=City6
//City3 north=City1 west=City7 east=City8
//City4 west=City1 north=City9
//City5 east=City2 south=City10
//City6 west=City2 north=City11
//City7 east=City3 south=City12
//City8 west=City3 north=City13
//City9 south=City4 east=City14
//City10 north=City5 east=City15
//City11 south=City6 west=City16
//City12 north=City7 east=City17
//City13 south=City8 west=City18
//City14 west=City9 north=City19
//City15 south=City10 west=City20
//City16 east=City11 north=City21
//City17 west=City12 south=City22
//City18 east=City13 north=City23
//City19 south=City14 east=City24
//City20 north=City15 west=City25
//City21 south=City16 east=City26
//City22 north=City17 west=City27
//City23 south=City18 east=City28
//City24 west=City19 north=City29
//City25 east=City20 south=City30
//City26 west=City21 north=City31
//City27 east=City22 south=City32
//City28 west=City23 north=City33
//City29 south=City24 east=City34
//City30 north=City25 west=City35
//City31 south=City26 east=City36
//City32 north=City27 west=City37
//City33 south=City28 east=City38
//City34 west=City29 north=City39
//City35 east=City30 south=City40
//City36 west=City31 north=City41
//City37 east=City32 south=City42
//City38 west=City33 north=City43
//City39 south=City34 east=City44
//City40 north=City35 west=City45
//City41 south=City36 east=City46
//City42 north=City37 west=City47
//City43 south=City38 east=City48
//City44 west=City39 north=City49
//City45 east=City40 south=City50
//City46 west=City41 north=City51
//City47 east=City42 south=City52
//City48 west=City43 north=City53
//City49 south=City44 east=City54
//City50 north=City45 west=City55
//City51 south=City46 east=City56
//City52 north=City47 west=City57
//City53 south=City48 east=City58
//City54 west=City49 north=City59
//City55 east=City50 south=City60
//City56 west=City51 north=City61
//City57 east=City52 south=City62
//City58 west=City53 north=City63
//City59 south=City54 east=City64
//City60 north=City55 west=City65
//
