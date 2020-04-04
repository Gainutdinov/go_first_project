package main

import (
	"errors"
	"strings"
	"log"
)

var ErrInvalidSSNLength = errors.New("invalid SSN length")
var ErrInvalidSSNNumbers = errors.New("SSN has non-numeric digits")
var ErrInvalidSSNPrefix  = errors.New("SSN that has three zeros as the prefix")
var ErrInvalidDigitPlace = errors.New("SSNs that starts with a 9 it requires 7 or 9 in the fourth place")

//func retInvSSN() int {
//	return 12345678
//}

func checkSSNLng (ssn string) (error) {
	//s := strconv.Itoa(ssn)
	ssn = strings.Replace(ssn, "-", "", -1)
	if len(ssn) < 9 {
		return ErrInvalidSSNNumbers
	}
	return nil
}

func checkSSNPrx(ssn string) (error) {
	//s := strconv.Itoa(ssn)
	ssn = strings.Replace(ssn, "-", "", -1)
	if strings.Contains(ssn, "000") {
		return ErrInvalidSSNPrefix
	}
	return nil
}

func checkSSNLogic(ssn string) (error) {
	//s := strconv.Itoa(ssn)
	ssn = strings.Replace(ssn, "-", "", -1)
	if ssn[0] == '9' && (ssn[3] != '7' || ssn[3] != '9') {
		return ErrInvalidDigitPlace
	}
	return nil
}

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	validateSSN := []string{"123-45-6789", "012-8-678", "000-12-0962", "999-33-3333", "087-65-4321","123-45-zzzz"}
	log.Println("%#v", validateSSN)
	for idx, ssn := range validateSSN {
		log.Printf("Validate data %#v %d of %d ",ssn,idx+1,len(validateSSN))
		err := checkSSNLng(ssn)
        if err != nil {
            log.Print(err)
        }
		err = checkSSNPrx(ssn)
        if err != nil {
            log.Print(err)
        }
		err = checkSSNLogic(ssn)
        if err != nil {
            log.Print(err)
        }
	}
}
