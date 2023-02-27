package models

import (
	"strings"
	"unicode"
)

const specialChars string = "!@#$%^&*()-+\\/{}[]"

type ObjectRequest struct {
	Password string       `json:"Password"`
	Rules    []nestedRule `json:"Rules"`
}

type nestedRule struct {
	Rule  string `json:"rule"`
	Value int    `json:"value"`
}

type Result struct {
	Verify  bool     `json:"verify"`
	NoMatch []string `json:"noMatch"`
}

// função para tratar variáveis e retornar a struct
func CheckPassword(obj ObjectRequest) Result {
	var noRepeated bool
	var minSize, minUppercase, minLowercase, minDigit, minSpecialChars int
	var result Result

	password := obj.Password
	for _, nestedRule := range obj.Rules {
		switch nestedRule.Rule {
		case "noRepeted":
			noRepeated = nestedRule.Value != 0
		case "minSize":
			minSize = nestedRule.Value
		case "minUpperCase":
			minUppercase = nestedRule.Value
		case "minLowercase":
			minLowercase = nestedRule.Value
		case "minDigit":
			minDigit = nestedRule.Value
		case "minSpecialChars":
			minSpecialChars = nestedRule.Value
		}
	}

	result.NoMatch = checkPassword(password, noRepeated, minSize, minUppercase, minLowercase, minDigit, minSpecialChars)
	result.Verify = len(result.NoMatch) == 0

	return result
}

// password string, noRepeated bool, minSize, minUppercase, minLowercase, minDigit, minSpecialChars int
// função para checar se o password passa, retornando um slice com um comentário para cada check falho, caso o slice retorne vazio, o password passou por todas as checagens passadas pelos parâmetros
func checkPassword(password string, noRepeated bool, minSize, minUppercase, minLowercase, minDigit, minSpecialChars int) []string {
	//Tentei fazer essa função usando regex mas retorna esse erro
	//error parsing regexp: invalid or unsupported Perl syntax: `(?=`
	//é possível alterar o regex para que ele funcione, mas eu ia perder muito tempo, então resolvi focar em outras partes mais importantes do desafio
	checkList := []string{}

	if !checkMinSize(password, minSize) {
		checkList = append(checkList, "minSize")
	}
	if !checkMinUpperCase(password, minUppercase) {
		checkList = append(checkList, "minUppercaseo")
	}
	if !checkMinLowerCase(password, minLowercase) {
		checkList = append(checkList, "minLowercase")
	}
	if !checkMinDigit(password, minDigit) {
		checkList = append(checkList, "minDigit")
	}
	if !checkMinSpecialChar(password, minSpecialChars) {
		checkList = append(checkList, "minSpecialChars")
	}
	if checkRepeatedChar(password, noRepeated) {
		checkList = append(checkList, "noRepeted")
	}

	return checkList

}

func checkMinSize(password string, minSize int) bool {
	return len(password) >= minSize
}

func checkMinUpperCase(password string, minUppercase int) bool {
	var countMinUpper int = 0
	for _, r := range password {
		if unicode.IsUpper(r) {
			countMinUpper++
		}
	}
	return countMinUpper >= minUppercase
}

func checkMinLowerCase(password string, minLowercase int) bool {
	var countMinLower int = 0
	for _, r := range password {
		if unicode.IsLower(r) {
			countMinLower++
		}
	}
	return countMinLower >= minLowercase
}

func checkMinDigit(password string, minDigit int) bool {
	var countMinDigit int = 0
	for _, r := range password {
		if unicode.IsDigit(r) {
			countMinDigit++
		}
	}
	return countMinDigit >= minDigit
}

func checkMinSpecialChar(password string, minSpecialChar int) bool {
	var countMinSpecial int = 0
	for _, r := range password {
		if strings.ContainsRune(specialChars, r) {
			countMinSpecial++
		}
	}
	return countMinSpecial >= minSpecialChar
}

func checkRepeatedChar(password string, noRepeated bool) bool {
	if noRepeated {
		for x := range password {
			if x > 0 && password[x] == password[x-1] {
				return true
			}
		}
	}
	return false
}
