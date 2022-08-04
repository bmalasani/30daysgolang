/*
Package bytes implements functions for the manipulation of byte slices. It is analogous to the facilities of the strings package.
*/
package main

import (
	"bytes"
	"fmt"

	"meb.com/utils"
)

func compare(a, b []byte) {
	i := bytes.Compare(a, b)
	if i > 0 {
		utils.WriteLog(string(a) + " > " + string(b))
	} else if i < 0 {
		utils.WriteLog(string(a) + " < " + string(b))
	} else {
		utils.WriteLog(string(a) + " = " + string(b))
	}

}

func main() {
	for {
		option := utils.ReadFromConsole(
			`Choose an option (using bytes):
			1. compare two strings and equal
			2. contains
			3. count byte
			4. cut bytes based on separate
			5. split based on space and func go Feilds
			6. startswith and ends with
			7. get index
			8. join 
			9. lastindex
			10. repeat
			11. replace
			12. runes
			13. split
			14. conversiont toLower
			15. tovalidutf
			16. trim`)

		switch option {
		case "1":
			a := bytes.NewBufferString(utils.ReadFromConsole("first string")).Bytes()
			b := bytes.NewBufferString(utils.ReadFromConsole("second string")).Bytes()
			compare(a, b)
			utils.WriteLog(fmt.Sprintf("Equal %v", bytes.Equal(a, b)))
			utils.WriteLog(fmt.Sprintf("Equal Fold %v", bytes.EqualFold(a, b)))
			break
		case "2":
			a := bytes.NewBufferString(utils.ReadFromConsole("first string")).Bytes()
			b := bytes.NewBufferString(utils.ReadFromConsole("second string")).Bytes()
			r := bytes.Runes(b)[0]
			s := utils.ReadFromConsole("contains any string")
			utils.WriteLog(fmt.Sprintf("contains %v", bytes.Contains(a, b)))
			utils.WriteLog(fmt.Sprintf("contains any %v", bytes.ContainsAny(a, s)))
			utils.WriteLog(fmt.Sprintf("contains rune %v", bytes.ContainsRune(a, r)))
		case "3":
			a := []byte(utils.ReadFromConsole("string"))
			b := []byte(utils.ReadFromConsole("separator"))
			fmt.Printf("bytes.Count(a, b): %v\n", bytes.Count(a, b))
		case "4":
			a := []byte(utils.ReadFromConsole("string"))
			b := []byte(utils.ReadFromConsole("separator"))
			before, after, found := bytes.Cut(a, b)
			fmt.Printf("before: %v\n", before)
			fmt.Printf("after: %v\n", after)
			fmt.Printf("found: %v\n", found)
		case "5":
			a := []byte(utils.ReadFromConsole("string"))
			aSplits := bytes.Fields(a)
			fmt.Printf("aSplits: %s\n len: %v", aSplits, len(aSplits))
			eSplits := bytes.FieldsFunc(a, func(r rune) bool { return r == rune('🥰') })
			fmt.Printf("eSplits: %s\n len: %v", eSplits, len(eSplits))
		case "6":
			a := []byte(utils.ReadFromConsole("string"))
			b := []byte(utils.ReadFromConsole("separator"))
			fmt.Printf("bytes.HasPrefix(a, b): %v\n", bytes.HasPrefix(a, b))
			fmt.Printf("bytes.HasSuffix(a, b): %v\n", bytes.HasSuffix(a, b))
		case "7":
			a := []byte(utils.ReadFromConsole("string"))
			b := []byte(utils.ReadFromConsole("separator"))
			fmt.Printf("bytes.Index(a, b): %d\n", bytes.Index(a, b))
			fmt.Printf("bytes.IndexByte(a, byte('9')): %v\n", bytes.IndexByte(a, byte('9')))
			fmt.Printf("bytes.IndexAny(a, \"తీగ\"): %v\n", bytes.IndexAny(a, "తీగ"))
			fmt.Printf("bytes.IndexFunc(a, func(r rune) bool { return r == rune('త') }): %v\n", bytes.IndexFunc(a, func(r rune) bool { return r == rune('త') }))
			fmt.Printf("bytes.IndexRune(a, rune('గ')): %v\n", bytes.IndexRune(a, rune('గ')))
		case "8":
			a := []byte(utils.ReadFromConsole("string"))
			b := []byte(utils.ReadFromConsole("separator"))
			f := bytes.Fields(a)
			fmt.Printf("bytes.Join(f, b): %v\n", bytes.Join(f, b))
		case "9":
			a := []byte(utils.ReadFromConsole("string"))
			b := []byte(utils.ReadFromConsole("separator"))
			fmt.Printf("bytes.LastIndex(a, b): %d\n", bytes.LastIndex(a, b))
			fmt.Printf("bytes.LastIndexByte(a, byte('9')): %v\n", bytes.LastIndexByte(a, byte('9')))
			fmt.Printf("bytes.LastIndexAny(a, \"తీగ\"): %v\n", bytes.LastIndexAny(a, "తీగ"))
			fmt.Printf("bytes.LastIndexFunc(a, func(r rune) bool { return r == rune('త') }): %v\n", bytes.LastIndexFunc(a, func(r rune) bool { return r == rune('త') }))
		case "10":
			a := []byte(utils.ReadFromConsole("string"))
			fmt.Printf("bytes.Repeat(a, 2): %v\n", bytes.Repeat(a, 2))
		case "11":
			a := []byte(utils.ReadFromConsole("string"))
			b := []byte(utils.ReadFromConsole("old"))
			c := []byte(utils.ReadFromConsole("new"))
			fmt.Printf("bytes.Replace(a, b, c, -1): %v\n", bytes.Replace(a, b, c, 1))
			fmt.Printf("bytes.ReplaceAll(a, b, c): %v\n", bytes.ReplaceAll(a, b, c))
		case "12":
			a := []byte(utils.ReadFromConsole("string"))
			fmt.Printf("bytes.Runes(a): %v\n", bytes.Runes(a))
		case "13":
			a := []byte(utils.ReadFromConsole("string"))
			b := []byte(utils.ReadFromConsole("separator"))
			res := bytes.Split(a, b)
			fmt.Printf("bytes.Split(a,b): %s\n", res)
			fmt.Printf("bytes.SplitAfter(a, b): %s\n", bytes.SplitAfter(a, b))
			fmt.Printf("bytes.SplitN(a, b, 1): %v\n", bytes.SplitN(a, b, 1))
			fmt.Printf("bytes.SplitAfterN(a, b, 1): %v\n", bytes.SplitAfterN(a, b, 1))
		case "14":
			a := []byte(utils.ReadFromConsole("string"))
			fmt.Printf("bytes.ToLower(a): %s\n", bytes.ToLower(a))
			fmt.Printf("bytes.ToTitle(a): %s\n", bytes.ToTitle(a))
			fmt.Printf("bytes.ToUpper(a): %s\n", bytes.ToUpper(a))
		case "15":
			a := []byte(utils.ReadFromConsole("string"))
			b := []byte(utils.ReadFromConsole("string"))
			fmt.Printf("bytes.ToValidUTF8(a, b): %s\n", bytes.ToValidUTF8(a, b))
		case "16":
			a := []byte(utils.ReadFromConsole("string"))
			b := []byte(utils.ReadFromConsole("prefix/suffix"))
			fmt.Printf("bytes.Trim(a, \"9876\"): %s\n", bytes.Trim(a, "9876"))
			fmt.Printf("bytes.TrimLeft(a, \"098\"): %s\n", bytes.TrimLeft(a, "098"))
			fmt.Printf("bytes.TrimRight(a, \"098\"): %s\n", bytes.TrimRight(a, "098"))
			fmt.Printf("bytes.TrimSpace(a): %s\n", bytes.TrimSpace(a))
			fmt.Printf("bytes.TrimPrefix(a, b): %s\n", bytes.TrimPrefix(a, b))
			fmt.Printf("bytes.TrimSuffix(a, b): %s\n", bytes.TrimSuffix(a, b))
		default:
			utils.WriteLog("Invalid choice..")

		}
	}
}
