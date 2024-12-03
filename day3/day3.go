package day3

import (
	"aoc2024/utils"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Token struct {
	Type    string
	Literal string
}

type MulStmt struct {
	Type string
	Arg1 int
	Arg2 int
}

const (
	M          = "m"
	U          = "u"
	L          = "l"
	LPAREN     = "("
	RPAREN     = ")"
	COMMA      = ","
	DIGIT      = "DIGIT"
	EOL        = "EOL"
	MUL        = "MUL"
	EXPRESSION = "EXPRESSION"
	WHITESPACE = "WHITESPACE"
)

func parseLine(line string) []MulStmt {
	curr := 0
	var tokens []Token
	var stmts []MulStmt

	for curr != len(line) {
		currChar := string(line[curr])
		switch currChar {
		case M:
			tokens = append(tokens, Token{Type: M, Literal: "m"})
			curr++
		case U:
			tokens = append(tokens, Token{Type: U, Literal: "u"})
			curr++
		case L:
			tokens = append(tokens, Token{Type: L, Literal: "l"})
			curr++
		case LPAREN:
			tokens = append(tokens, Token{Type: LPAREN, Literal: "("})
			curr++
		case RPAREN:
			tokens = append(tokens, Token{Type: RPAREN, Literal: ")"})
			curr++
		case COMMA:
			tokens = append(tokens, Token{Type: COMMA, Literal: ","})
			curr++
		default:
			if unicode.IsSpace(rune(line[curr])) {
				tokens = append(tokens, Token{Type: WHITESPACE, Literal: "WHITESPACE"})
				curr++
			} else {
				num := ""
				for utils.IsDigit(line[curr]) {
					num += string(line[curr])
					curr++
				}
				if num != "" {
					tokens = append(tokens, Token{Type: DIGIT, Literal: num})
					continue
				}
				curr++
			}
		}
	}
	tokens = append(tokens, Token{Type: EOL, Literal: "EOL"})
	// fmt.Println(tokens)
	parserCurr := 0
	for tokens[parserCurr].Type != EOL {
		if match(M, tokens[parserCurr], &parserCurr) &&
			match(U, tokens[parserCurr], &parserCurr) &&
			match(L, tokens[parserCurr], &parserCurr) {

			if match(LPAREN, tokens[parserCurr], &parserCurr) {
				arg1, err := matchAndGet(DIGIT, tokens[parserCurr], &parserCurr)
				if err != nil {
					parserCurr += 1
					continue
				}
				if match(COMMA, tokens[parserCurr], &parserCurr) {
					arg2, err := matchAndGet(DIGIT, tokens[parserCurr], &parserCurr)

					// fmt.Println("i run")
					if err != nil {
						parserCurr += 1
						continue
					}

					if match(RPAREN, tokens[parserCurr], &parserCurr) {

						a1, err := strconv.Atoi(arg1)

						if err != nil {
							panic(err)
						}
						a2, err := strconv.Atoi(arg2)

						if err != nil {
							panic(err)
						}
						stmts = append(stmts, MulStmt{Type: "MUL", Arg1: a1, Arg2: a2})
					} else {
                        parserCurr+=1
                    }
				} else {
                    parserCurr+=1
                }
			} else {
                parserCurr+=1
            }
		} else {
            parserCurr += 1
        }

	}

	return stmts
}

func match(expectedType string, actual Token, curr *int) bool {
	if expectedType != actual.Type {
		return false
	}
	*curr += 1
	return true
}

func matchAndGet(expectedType string, actual Token, curr *int) (string, error) {
	if expectedType != actual.Type {
		return "", errors.New("parser error")
	}
	*curr += 1
	return actual.Literal, nil
}

func eval(stmts []MulStmt) int {
	answer := 0
	for _, stmt := range stmts {
		answer += stmt.Arg1 * stmt.Arg2
	}
	return answer
}

func First(lines []string) {
	// fmt.Println(lines)
	lines = utils.ReadLinesToSlice("day3/input.txt")
	line := strings.Join(lines, "")
	stmts := parseLine(line)
	answer := 0
	answer += eval(stmts)

	// 180187256 too low
	// 184340710 too high
    // 183511462 wrong

	// for _, line := range lines {
	// 	stmts := parseLine(line)
	// 	answer += eval(stmts)
	// }

	fmt.Println(answer)
}

func Second(lines []string) {

}
