package main

import (
	"fmt"
)

// find the sum of all digits that match the next digit in the list

//649713959682898259577777982349515784822684939966191359164369933435366431847754488661965363557985166219358714739318371382388296151195361571216131925158492441461844687324923315381358331571577613789649166486152237945917987977793891739865149734755993241361886336926538482271124755359572791451335842534893192693558659991171983849285489139421425933638614884415896938914992732492192458636484523228244532331587584779552788544667253577324649915274115924611758345676183443982992733966373498385685965768929241477983727921279826727976872556315428434799161759734932659829934562339385328119656823483954856427365892627728163524721467938449943358192632262354854593635831559352247443975945144163183563723562891357859367964126289445982135523535923113589316417623483631637569291941782992213889513714525342468563349385271884221685549996534333765731243895662624829924982971685443825366827923589435254514211489649482374876434549682785459698885521673258939413255158196525696236457911447599947449665542554251486847388823576937167237476556782133227279324526834946534444718161524129285919477959937684728882592779941734186144138883994322742484853925383518651687147246943421311287324867663698432546619583638976637733345251834869985746385371617743498627111441933546356934671639545342515392536574744795732243617113574641284231928489312683617154536648219244996491745718658151648246791826466973654765284263928884137863647623237345882469142933142637583644258427416972595241737254449718531724176538648369253796688931245191382956961544775856872281317743828552629843551844927913147518377362266554334386721313244223233396453291224932499277961525785755863852487141946626663835195286762947172384186667439516367219391823774338692151926472717373235612911848773387771244144969149482477519437822863422662157461968444281972353149695515494992537927492111388193837553844671719291482442337761321272333982924289323437277224565149928416255435841327756139118119744528993269157174414264387573331116323982614862952264597611885999285995516357519648695594299657387614793341626318866519144574571816535351149394735916975448425618171572917195165594323552199346814729617189679698944337146
var input = []int{6, 4, 9, 7, 1, 3, 9, 5, 9, 6, 8, 2, 8, 9, 8, 2, 5, 9, 5, 7, 7, 7, 7, 7, 9, 8, 2, 3, 4, 9, 5, 1, 5, 7, 8, 4, 8, 2, 2, 6, 8, 4, 9, 3, 9, 9, 6, 6, 1, 9, 1, 3, 5, 9, 1, 6, 4, 3, 6, 9, 9, 3, 3, 4, 3, 5, 3, 6, 6, 4, 3, 1, 8, 4, 7, 7, 5, 4, 4, 8, 8, 6, 6, 1, 9, 6, 5, 3, 6, 3, 5, 5, 7, 9, 8, 5, 1, 6, 6, 2, 1, 9, 3, 5, 8, 7, 1, 4, 7, 3, 9, 3, 1, 8, 3, 7, 1, 3, 8, 2, 3, 8, 8, 2, 9, 6, 1, 5, 1, 1, 9, 5, 3, 6, 1, 5, 7, 1, 2, 1, 6, 1, 3, 1, 9, 2, 5, 1, 5, 8, 4, 9, 2, 4, 4, 1, 4, 6, 1, 8, 4, 4, 6, 8, 7, 3, 2, 4, 9, 2, 3, 3, 1, 5, 3, 8, 1, 3, 5, 8, 3, 3, 1, 5, 7, 1, 5, 7, 7, 6, 1, 3, 7, 8, 9, 6, 4, 9, 1, 6, 6, 4, 8, 6, 1, 5, 2, 2, 3, 7, 9, 4, 5, 9, 1, 7, 9, 8, 7, 9, 7, 7, 7, 9, 3, 8, 9, 1, 7, 3, 9, 8, 6, 5, 1, 4, 9, 7, 3, 4, 7, 5, 5, 9, 9, 3, 2, 4, 1, 3, 6, 1, 8, 8, 6, 3, 3, 6, 9, 2, 6, 5, 3, 8, 4, 8, 2, 2, 7, 1, 1, 2, 4, 7, 5, 5, 3, 5, 9, 5, 7, 2, 7, 9, 1, 4, 5, 1, 3, 3, 5, 8, 4, 2, 5, 3, 4, 8, 9, 3, 1, 9, 2, 6, 9, 3, 5, 5, 8, 6, 5, 9, 9, 9, 1, 1, 7, 1, 9, 8, 3, 8, 4, 9, 2, 8, 5, 4, 8, 9, 1, 3, 9, 4, 2, 1, 4, 2, 5, 9, 3, 3, 6, 3, 8, 6, 1, 4, 8, 8, 4, 4, 1, 5, 8, 9, 6, 9, 3, 8, 9, 1, 4, 9, 9, 2, 7, 3, 2, 4, 9, 2, 1, 9, 2, 4, 5, 8, 6, 3, 6, 4, 8, 4, 5, 2, 3, 2, 2, 8, 2, 4, 4, 5, 3, 2, 3, 3, 1, 5, 8, 7, 5, 8, 4, 7, 7, 9, 5, 5, 2, 7, 8, 8, 5, 4, 4, 6, 6, 7, 2, 5, 3, 5, 7, 7, 3, 2, 4, 6, 4, 9, 9, 1, 5, 2, 7, 4, 1, 1, 5, 9, 2, 4, 6, 1, 1, 7, 5, 8, 3, 4, 5, 6, 7, 6, 1, 8, 3, 4, 4, 3, 9, 8, 2, 9, 9, 2, 7, 3, 3, 9, 6, 6, 3, 7, 3, 4, 9, 8, 3, 8, 5, 6, 8, 5, 9, 6, 5, 7, 6, 8, 9, 2, 9, 2, 4, 1, 4, 7, 7, 9, 8, 3, 7, 2, 7, 9, 2, 1, 2, 7, 9, 8, 2, 6, 7, 2, 7, 9, 7, 6, 8, 7, 2, 5, 5, 6, 3, 1, 5, 4, 2, 8, 4, 3, 4, 7, 9, 9, 1, 6, 1, 7, 5, 9, 7, 3, 4, 9, 3, 2, 6, 5, 9, 8, 2, 9, 9, 3, 4, 5, 6, 2, 3, 3, 9, 3, 8, 5, 3, 2, 8, 1, 1, 9, 6, 5, 6, 8, 2, 3, 4, 8, 3, 9, 5, 4, 8, 5, 6, 4, 2, 7, 3, 6, 5, 8, 9, 2, 6, 2, 7, 7, 2, 8, 1, 6, 3, 5, 2, 4, 7, 2, 1, 4, 6, 7, 9, 3, 8, 4, 4, 9, 9, 4, 3, 3, 5, 8, 1, 9, 2, 6, 3, 2, 2, 6, 2, 3, 5, 4, 8, 5, 4, 5, 9, 3, 6, 3, 5, 8, 3, 1, 5, 5, 9, 3, 5, 2, 2, 4, 7, 4, 4, 3, 9, 7, 5, 9, 4, 5, 1, 4, 4, 1, 6, 3, 1, 8, 3, 5, 6, 3, 7, 2, 3, 5, 6, 2, 8, 9, 1, 3, 5, 7, 8, 5, 9, 3, 6, 7, 9, 6, 4, 1, 2, 6, 2, 8, 9, 4, 4, 5, 9, 8, 2, 1, 3, 5, 5, 2, 3, 5, 3, 5, 9, 2, 3, 1, 1, 3, 5, 8, 9, 3, 1, 6, 4, 1, 7, 6, 2, 3, 4, 8, 3, 6, 3, 1, 6, 3, 7, 5, 6, 9, 2, 9, 1, 9, 4, 1, 7, 8, 2, 9, 9, 2, 2, 1, 3, 8, 8, 9, 5, 1, 3, 7, 1, 4, 5, 2, 5, 3, 4, 2, 4, 6, 8, 5, 6, 3, 3, 4, 9, 3, 8, 5, 2, 7, 1, 8, 8, 4, 2, 2, 1, 6, 8, 5, 5, 4, 9, 9, 9, 6, 5, 3, 4, 3, 3, 3, 7, 6, 5, 7, 3, 1, 2, 4, 3, 8, 9, 5, 6, 6, 2, 6, 2, 4, 8, 2, 9, 9, 2, 4, 9, 8, 2, 9, 7, 1, 6, 8, 5, 4, 4, 3, 8, 2, 5, 3, 6, 6, 8, 2, 7, 9, 2, 3, 5, 8, 9, 4, 3, 5, 2, 5, 4, 5, 1, 4, 2, 1, 1, 4, 8, 9, 6, 4, 9, 4, 8, 2, 3, 7, 4, 8, 7, 6, 4, 3, 4, 5, 4, 9, 6, 8, 2, 7, 8, 5, 4, 5, 9, 6, 9, 8, 8, 8, 5, 5, 2, 1, 6, 7, 3, 2, 5, 8, 9, 3, 9, 4, 1, 3, 2, 5, 5, 1, 5, 8, 1, 9, 6, 5, 2, 5, 6, 9, 6, 2, 3, 6, 4, 5, 7, 9, 1, 1, 4, 4, 7, 5, 9, 9, 9, 4, 7, 4, 4, 9, 6, 6, 5, 5, 4, 2, 5, 5, 4, 2, 5, 1, 4, 8, 6, 8, 4, 7, 3, 8, 8, 8, 2, 3, 5, 7, 6, 9, 3, 7, 1, 6, 7, 2, 3, 7, 4, 7, 6, 5, 5, 6, 7, 8, 2, 1, 3, 3, 2, 2, 7, 2, 7, 9, 3, 2, 4, 5, 2, 6, 8, 3, 4, 9, 4, 6, 5, 3, 4, 4, 4, 4, 7, 1, 8, 1, 6, 1, 5, 2, 4, 1, 2, 9, 2, 8, 5, 9, 1, 9, 4, 7, 7, 9, 5, 9, 9, 3, 7, 6, 8, 4, 7, 2, 8, 8, 8, 2, 5, 9, 2, 7, 7, 9, 9, 4, 1, 7, 3, 4, 1, 8, 6, 1, 4, 4, 1, 3, 8, 8, 8, 3, 9, 9, 4, 3, 2, 2, 7, 4, 2, 4, 8, 4, 8, 5, 3, 9, 2, 5, 3, 8, 3, 5, 1, 8, 6, 5, 1, 6, 8, 7, 1, 4, 7, 2, 4, 6, 9, 4, 3, 4, 2, 1, 3, 1, 1, 2, 8, 7, 3, 2, 4, 8, 6, 7, 6, 6, 3, 6, 9, 8, 4, 3, 2, 5, 4, 6, 6, 1, 9, 5, 8, 3, 6, 3, 8, 9, 7, 6, 6, 3, 7, 7, 3, 3, 3, 4, 5, 2, 5, 1, 8, 3, 4, 8, 6, 9, 9, 8, 5, 7, 4, 6, 3, 8, 5, 3, 7, 1, 6, 1, 7, 7, 4, 3, 4, 9, 8, 6, 2, 7, 1, 1, 1, 4, 4, 1, 9, 3, 3, 5, 4, 6, 3, 5, 6, 9, 3, 4, 6, 7, 1, 6, 3, 9, 5, 4, 5, 3, 4, 2, 5, 1, 5, 3, 9, 2, 5, 3, 6, 5, 7, 4, 7, 4, 4, 7, 9, 5, 7, 3, 2, 2, 4, 3, 6, 1, 7, 1, 1, 3, 5, 7, 4, 6, 4, 1, 2, 8, 4, 2, 3, 1, 9, 2, 8, 4, 8, 9, 3, 1, 2, 6, 8, 3, 6, 1, 7, 1, 5, 4, 5, 3, 6, 6, 4, 8, 2, 1, 9, 2, 4, 4, 9, 9, 6, 4, 9, 1, 7, 4, 5, 7, 1, 8, 6, 5, 8, 1, 5, 1, 6, 4, 8, 2, 4, 6, 7, 9, 1, 8, 2, 6, 4, 6, 6, 9, 7, 3, 6, 5, 4, 7, 6, 5, 2, 8, 4, 2, 6, 3, 9, 2, 8, 8, 8, 4, 1, 3, 7, 8, 6, 3, 6, 4, 7, 6, 2, 3, 2, 3, 7, 3, 4, 5, 8, 8, 2, 4, 6, 9, 1, 4, 2, 9, 3, 3, 1, 4, 2, 6, 3, 7, 5, 8, 3, 6, 4, 4, 2, 5, 8, 4, 2, 7, 4, 1, 6, 9, 7, 2, 5, 9, 5, 2, 4, 1, 7, 3, 7, 2, 5, 4, 4, 4, 9, 7, 1, 8, 5, 3, 1, 7, 2, 4, 1, 7, 6, 5, 3, 8, 6, 4, 8, 3, 6, 9, 2, 5, 3, 7, 9, 6, 6, 8, 8, 9, 3, 1, 2, 4, 5, 1, 9, 1, 3, 8, 2, 9, 5, 6, 9, 6, 1, 5, 4, 4, 7, 7, 5, 8, 5, 6, 8, 7, 2, 2, 8, 1, 3, 1, 7, 7, 4, 3, 8, 2, 8, 5, 5, 2, 6, 2, 9, 8, 4, 3, 5, 5, 1, 8, 4, 4, 9, 2, 7, 9, 1, 3, 1, 4, 7, 5, 1, 8, 3, 7, 7, 3, 6, 2, 2, 6, 6, 5, 5, 4, 3, 3, 4, 3, 8, 6, 7, 2, 1, 3, 1, 3, 2, 4, 4, 2, 2, 3, 2, 3, 3, 3, 9, 6, 4, 5, 3, 2, 9, 1, 2, 2, 4, 9, 3, 2, 4, 9, 9, 2, 7, 7, 9, 6, 1, 5, 2, 5, 7, 8, 5, 7, 5, 5, 8, 6, 3, 8, 5, 2, 4, 8, 7, 1, 4, 1, 9, 4, 6, 6, 2, 6, 6, 6, 3, 8, 3, 5, 1, 9, 5, 2, 8, 6, 7, 6, 2, 9, 4, 7, 1, 7, 2, 3, 8, 4, 1, 8, 6, 6, 6, 7, 4, 3, 9, 5, 1, 6, 3, 6, 7, 2, 1, 9, 3, 9, 1, 8, 2, 3, 7, 7, 4, 3, 3, 8, 6, 9, 2, 1, 5, 1, 9, 2, 6, 4, 7, 2, 7, 1, 7, 3, 7, 3, 2, 3, 5, 6, 1, 2, 9, 1, 1, 8, 4, 8, 7, 7, 3, 3, 8, 7, 7, 7, 1, 2, 4, 4, 1, 4, 4, 9, 6, 9, 1, 4, 9, 4, 8, 2, 4, 7, 7, 5, 1, 9, 4, 3, 7, 8, 2, 2, 8, 6, 3, 4, 2, 2, 6, 6, 2, 1, 5, 7, 4, 6, 1, 9, 6, 8, 4, 4, 4, 2, 8, 1, 9, 7, 2, 3, 5, 3, 1, 4, 9, 6, 9, 5, 5, 1, 5, 4, 9, 4, 9, 9, 2, 5, 3, 7, 9, 2, 7, 4, 9, 2, 1, 1, 1, 3, 8, 8, 1, 9, 3, 8, 3, 7, 5, 5, 3, 8, 4, 4, 6, 7, 1, 7, 1, 9, 2, 9, 1, 4, 8, 2, 4, 4, 2, 3, 3, 7, 7, 6, 1, 3, 2, 1, 2, 7, 2, 3, 3, 3, 9, 8, 2, 9, 2, 4, 2, 8, 9, 3, 2, 3, 4, 3, 7, 2, 7, 7, 2, 2, 4, 5, 6, 5, 1, 4, 9, 9, 2, 8, 4, 1, 6, 2, 5, 5, 4, 3, 5, 8, 4, 1, 3, 2, 7, 7, 5, 6, 1, 3, 9, 1, 1, 8, 1, 1, 9, 7, 4, 4, 5, 2, 8, 9, 9, 3, 2, 6, 9, 1, 5, 7, 1, 7, 4, 4, 1, 4, 2, 6, 4, 3, 8, 7, 5, 7, 3, 3, 3, 1, 1, 1, 6, 3, 2, 3, 9, 8, 2, 6, 1, 4, 8, 6, 2, 9, 5, 2, 2, 6, 4, 5, 9, 7, 6, 1, 1, 8, 8, 5, 9, 9, 9, 2, 8, 5, 9, 9, 5, 5, 1, 6, 3, 5, 7, 5, 1, 9, 6, 4, 8, 6, 9, 5, 5, 9, 4, 2, 9, 9, 6, 5, 7, 3, 8, 7, 6, 1, 4, 7, 9, 3, 3, 4, 1, 6, 2, 6, 3, 1, 8, 8, 6, 6, 5, 1, 9, 1, 4, 4, 5, 7, 4, 5, 7, 1, 8, 1, 6, 5, 3, 5, 3, 5, 1, 1, 4, 9, 3, 9, 4, 7, 3, 5, 9, 1, 6, 9, 7, 5, 4, 4, 8, 4, 2, 5, 6, 1, 8, 1, 7, 1, 5, 7, 2, 9, 1, 7, 1, 9, 5, 1, 6, 5, 5, 9, 4, 3, 2, 3, 5, 5, 2, 1, 9, 9, 3, 4, 6, 8, 1, 4, 7, 2, 9, 6, 1, 7, 1, 8, 9, 6, 7, 9, 6, 9, 8, 9, 4, 4, 3, 3, 7, 1, 4, 6}

//var input = []int{1, 2, 3, 1, 2, 3}

func main() {
	sizeOfInput := len(input)
	step := sizeOfInput / 2
	fmt.Println("sizeofinput:", sizeOfInput)
	fmt.Println("step:", step)
	sum := 0
	//comparator := 0
	for index, number := range input {
		//if index == 0 {
		//	comparator = number
		//	continue
		//}
		//fmt.Println("comparator:", comparator)
		//if comparator == opposite(index) {
		if number == opposite(index) {
			sum = sum + number
		}
		//if index+1 == len(input) {
		//	if input[0] == input[len(input)-1] {
		//		sum = sum + number
		//	}
		//}
		//comparator = number
	}
	fmt.Println("sum:", sum)
}

func opposite(location int) int {
	fmt.Println("location:", location)
	sizeOfInput := len(input)
	fmt.Println("sizeofinput:", sizeOfInput)
	step := sizeOfInput / 2
	fmt.Println("step:", step)
	superInput := append(input, input...)
	fmt.Println("superInput:", superInput)
	locationLooped := (location) + step
	fmt.Println("locationLooped:", locationLooped)
	fmt.Println("superInput[locationLooped]:", superInput[locationLooped])
	return superInput[locationLooped]
}