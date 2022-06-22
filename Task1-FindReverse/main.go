package main

import "fmt"

func main() {

	fmt.Println(findSameInReverse("FourscoreandsevenyearsagoourfaathersbroughtforthonthiscontainentanewnationconceivedinzLibertyanddedicatedtothepropositionthatallmenarecreatedequalNowweareengagedinagreahtcivilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"))
}

func findSameInReverse(value string) string { // in this function https://github.com/sangupta/greplin/blob/master/src/com/sangupta/greplin/Level3.java I referenced the this source
	bigVal := ""
	lenBigVal := 0

	lenAll := len(value)
	for i := 0; i < lenAll; i++ {
		for j := i + 1; j < lenAll; j++ {
			subString := value[i : j+1]

			if isPalindrome(subString) {
				if lenBigVal < len(subString) {
					lenBigVal = len(subString)
					bigVal = subString
				}

			}
		}
	}
	return bigVal
}

func reverseString(str string) string { //function to find the inverse of the entered value
	rns := []rune(str)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func isPalindrome(str string) bool { //function that tests the polyndrome of the entered value

	control := false
	revVal := reverseString(str)

	if revVal == str {
		control = true
	}
	return control

}
