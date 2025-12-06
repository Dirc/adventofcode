package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func getBoundary(idRangeString string) (int, int) {
	idRange := strings.Split(idRangeString, "-")
	firstInt, _ := strconv.Atoi(idRange[0])
	lastInt, _ := strconv.Atoi(idRange[1])
	return firstInt, lastInt
}

func getInput(filename string) ([][2]int, []int) {
	input := FileToArray(filename)
	spacePassed := false
	// idRanges list of first, last
	var idRanges [][2]int // tuples of boundarier (first, last)
	var ids []int
	for _, item := range input {
		if !spacePassed && item != "" {
			first, last := getBoundary(item)
			idRanges = append(idRanges, [2]int{first, last})
		}
		if item == "" {
			spacePassed = true
		}
		if spacePassed && item != "" {
			// Zou ook hier al kunnen sorteren
			integer, _ := strconv.Atoi(item)
			ids = append(ids, integer)
		}
	}
	return idRanges, ids
}

func intInRange(integer int, idRange [2]int) bool {
	if idRange[0] <= integer && integer <= idRange[1] {
		return true
	} else {
		return false
	}
}

func day5Part1(filename string) {
	idRanges, ids := getInput(filename)
	// lijkt me prettig?
	sort.Ints(ids)

	countFreshIds := 0
outerloop:
	for _, integer := range ids {
		for _, idRange := range idRanges {
			if intInRange(integer, idRange) {
				countFreshIds += 1
				continue outerloop
			}
		}
	}
	fmt.Println(idRanges)

	fmt.Println(ids)

	fmt.Println(countFreshIds)
}

// part 2
// - merge all overlapping ranges
// - answer = sum of all the len()

func mergeWithAllRanges(idRange [2]int, idRanges [][2]int) (int, int, int, int) {
	firstMin := idRange[0]
	lastMax := idRange[1]
	indexOfOverlappingRange1 := -1
	indexOfOverlappingRange2 := -1
	for i, idRange2 := range idRanges {
		// Overlapping ranges: if boundary (first, last) is in range
		if intInRange(firstMin, idRange2) {
			firstMin = idRange2[0]
			indexOfOverlappingRange1 = i
		}
		if intInRange(lastMax, idRange2) {
			lastMax = idRange2[1]
			indexOfOverlappingRange2 = i
		}
	}
	return firstMin, lastMax, indexOfOverlappingRange1, indexOfOverlappingRange2
}

func day5part2(filename string) {
	//idRanges, _ := getInput(filename)
	idRanges := [][2]int{{22741982971367, 24889881640387}, {541758546236603, 542296823918748}, {133005383318310, 133330351323182}, {521083224190969, 521083224190969}, {323915646289082, 327541745812280}, {192376219702627, 196507946925726}, {389337854383785, 389337854383785}, {137574801086037, 138398453840747}, {33992130933150, 38971398547270}, {154287753352708, 158810731646579}, {54302094410256, 59362935885072}, {439728990433, 8695261932182}, {376339501158728, 381604611064044}, {465459606490180, 468531579328884}, {123541227517975, 128703662180045}, {282882398773183, 291215147883697}, {525120536832404, 527310270401867}, {144606359606871, 145053217103956}, {163325612946356, 168516617422163}, {32894348022501, 33992130933148}, {327541745812281, 329283712774839}, {203080109153217, 207119117015433}, {42398150238901, 49653844991042}, {252766907402332, 259310791876078}, {138819382319231, 139261836115272}, {213549602869941, 215610823082662}, {303166215835169, 307691785267781}, {335816903520438, 340945307276811}, {473220288492448, 478599771270052}, {24889881640388, 29580576597309}, {527310270401869, 531618665702854}, {352501370303502, 353947624514333}, {267713944749686, 267713944749686}, {446433024217039, 448729561601283}, {303166215835168, 303166215835168}, {465459606490179, 465459606490179}, {540740105822766, 541001458282830}, {145418266283652, 146042777572641}, {493439568005479, 498314283532894}, {265131673201056, 267713944749685}, {196507946925728, 199508664326497}, {241715954764948, 249687755509541}, {453408464322441, 461801405832566}, {96329036110605, 100405168856315}, {13466688954434, 16534508927096}, {544643837005016, 550567021791178}, {183631556863435, 188096341984346}, {514950117014985, 521083224190968}, {149371693809555, 149694281324230}, {393776597708105, 400013752976450}, {142876008903557, 144167382118218}, {478599771270053, 479694347153683}, {74438960104880, 78459017137408}, {557039788176752, 561115148619896}, {146706776157239, 147325888330153}, {356240875051004, 357776726872387}, {358177250882224, 359479697496862}, {134830963604127, 136776394271078}, {172313000665694, 180382365642341}, {294364601598411, 299113922153738}, {82553613272457, 86364789802515}, {362007676422951, 371693076236875}, {62686232891971, 66366693869777}, {203080109153216, 203080109153216}, {360027425937529, 361613033113892}, {231868879864303, 236911062469528}, {215610823082663, 220392041729537}, {224570348282517, 228974295150441}, {90675707173237, 100405168856315}, {354161446406288, 355774685480476}, {374951558953544, 376339501158726}, {342630215904251, 348313634096109}, {131639004748978, 132879501052702}, {16534508927097, 16534508927097}, {483061504193911, 491902108020924}, {312407473700249, 317326359817995}, {101143648200130, 108471899337871}, {74438960104879, 74438960104879}, {404876905072551, 410392259177559}, {422730901225282, 431073659156880}, {534338078581316, 535119714633304}, {271653799343150, 281416026149419}, {506797107522421, 511208081376071}, {412903774044090, 421829883531644}, {148104690033483, 149072469910221}, {536151366421783, 536898988602940}, {111830263709630, 118899817015048}, {385578459420493, 389337854383784}, {537932457023000, 540036087289810}, {432664653548112, 441955307865263}}
	fmt.Println(len(idRanges))

	idRangesMerged := [][2]int{}

	// Naive: for each range: check with all others if it can be merged
	for _, idRange := range idRanges {
		//firstMin := idRange[0]
		//lastMax := idRange[1]

		// check all ranges
		firstMin, lastMax, _, _ := mergeWithAllRanges(idRange, idRanges)

		// also check in all currently merged ranges
		firstMin, lastMax, i, j := mergeWithAllRanges([2]int{firstMin, lastMax}, idRangesMerged)

		// Remove ranges that are merged
		if i != -1 { // remove index i
			idRangesMerged = append(idRangesMerged[:i], idRangesMerged[i+1:]...)
		}
		if j != -1 && j != i {
			idRangesMerged = append(idRangesMerged[:j], idRangesMerged[j+1:]...)
		}

		idRangesMerged = append(idRangesMerged, [2]int{firstMin, lastMax})
	}

	fmt.Println(idRanges)
	fmt.Println(idRangesMerged)

	totalLengthRanges := 0
	for _, idRange := range idRangesMerged {
		totalLengthRanges += idRange[1] - idRange[0] + 1
	}

	fmt.Println(totalLengthRanges)
	fmt.Println(" < ")
	fmt.Println(337074415782480)
	fmt.Println(len(idRangesMerged))

}
