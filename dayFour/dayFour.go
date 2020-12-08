package dayFour

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	BirthYear  string
	IssueYear  string
	ExpYear    string
	Height     string
	HairClr    string
	EyeClr     string
	PassportId string
	CountryId  string
}

func newPassport() Passport {
	return Passport{}
}

func (p *Passport) ValidateLoosely() bool {
	return p.BirthYear != "" && p.IssueYear != "" && p.ExpYear != "" && p.Height != "" &&
		p.HairClr != "" && p.EyeClr != "" && p.PassportId != ""
}

func (p *Passport) ValidateStringently() bool {
	if p.ValidateLoosely() {
		birthYear, err := strconv.Atoi(p.BirthYear)
		if !(err == nil &&
			len(p.BirthYear) == 4 &&
			1920 <= birthYear &&
			2002 >= birthYear) {
			return false
		}
		issueYear, err := strconv.Atoi(p.IssueYear)
		if !(err == nil &&
			len(p.IssueYear) == 4 &&
			2010 <= issueYear &&
			2020 >= issueYear) {
			return false
		}
		expYear, err := strconv.Atoi(p.ExpYear)
		if !(err == nil &&
			len(p.ExpYear) == 4 &&
			2020 <= expYear &&
			2030 >= expYear) {
			return false
		}
		if strings.Contains(p.Height, "cm") {
			height, err := strconv.Atoi(strings.TrimSuffix(p.Height, "cm"))
			if !(err == nil &&
				150 <= height &&
				193 >= height) {
				return false
			}
		} else if strings.Contains(p.Height, "in") {
			height, err := strconv.Atoi(strings.TrimSuffix(p.Height, "in"))
			if !(err == nil &&
				59 <= height &&
				76 >= height) {
				return false
			}
		} else {
			return false
		}
		var isAlphanumeric = regexp.MustCompile(`^[a-zA-Z0-9]{6}$`).MatchString
		if !(p.HairClr[0] == '#' &&
			isAlphanumeric(strings.TrimPrefix(p.HairClr, "#"))) {
			return false
		}
		if !(p.EyeClr == "amb" ||
			p.EyeClr == "blu" ||
			p.EyeClr == "brn" ||
			p.EyeClr == "gry" ||
			p.EyeClr == "grn" ||
			p.EyeClr == "hzl" ||
			p.EyeClr == "oth") {
			return false
		}
		var isNumeric = regexp.MustCompile(`^[0-9]{9}$`).MatchString
		return isNumeric(p.PassportId)
	}
	return false
}

func ScanPasswords() {
	passportFilePath := "passport.txt"

	file, err := os.Open(passportFilePath)
	if err != nil {
		fmt.Println("Failed to open passport file")
		os.Exit(3)
	}
	passports := parsePassportFile(file)
	looselyValidCount := 0
	stringentlyValidCount := 0
	for i := 0; i < len(passports); i++ {
		if passports[i].ValidateLoosely() {
			looselyValidCount++
		}
		if passports[i].ValidateStringently() {
			stringentlyValidCount++
		}
	}

	fmt.Println("Total Passports: " + strconv.Itoa(len(passports)))
	fmt.Println("Loosely Valid Passports: " + strconv.Itoa(looselyValidCount))
	fmt.Println("Stringently Valid Passports: " + strconv.Itoa(stringentlyValidCount))
}

func parsePassportFile(reader io.Reader) []Passport {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	var passports []Passport
	scannedPassport := newPassport()
	for scanner.Scan() {
		line := scanner.Text()
		//Separator
		if line == "" {
			passports = append(passports, scannedPassport)
			scannedPassport = newPassport()
		} else {
			params := strings.Split(line, " ")
			for i := 0; i < len(params); i++ {
				key := strings.Split(params[i], ":")[0]
				value := strings.Split(params[i], ":")[1]
				switch key {
				case "byr":
					scannedPassport.BirthYear = value
				case "iyr":
					scannedPassport.IssueYear = value
				case "eyr":
					scannedPassport.ExpYear = value
				case "hgt":
					scannedPassport.Height = value
				case "hcl":
					scannedPassport.HairClr = value
				case "ecl":
					scannedPassport.EyeClr = value
				case "pid":
					scannedPassport.PassportId = value
				case "cid":
					scannedPassport.CountryId = value
				default:
					os.Exit(4)
				}
			}
		}
	}
	passports = append(passports, scannedPassport)
	return passports
}
