package main

import (
	"fmt"

	"github.com/uyouii/rule_engine"
)

func main() {

	params := []*rule_engine.Param{
		rule_engine.GetParam("i", 100),
		rule_engine.GetParam("f", 3.5),
		rule_engine.GetParam("s", "hello world"),
		rule_engine.GetParam("b", false),
		rule_engine.GetParam("d", 3.3),
		// get decimal from string
		rule_engine.GetParamWithType("d2", rule_engine.ValueTypeDecimal, "3.3"),
	}

	// use decimal: true
	praser, err := rule_engine.GetNewPraser(params, true)
	if err != nil {
		panic(err)
	}

	exampleList := []string{
		// integrate
		`4 * (2 + 3) - 5 * 3`,
		// float
		`3.0 * (2.5 - 3)`,
		// logic
		`1 < 2 and 2 < 3 and 4.0 != 4.0001`,
		`1 > 2 && 1 < 2`,
		// func exapmle
		`max(min(10.0, 20, 30), len("vstr"))`,
		`min(len("test"), abs(-4.5), min(5,6))`,
		`upper("abc") == "ABC"`,
		`startWith("hello world", "hel")`,
		`int(97) + 3 == max(100, -1)`,
		`regexMatch("^0[xX][a-fA-F0-9]+", "0xasf4")`,
		`string({{d}}) == "3.3"`,
		// use param
		`min(len({{s}}), {{i}}, {{f}}, {{d}})`,
		`int({{i}} / {{f}})`,
		`{{d}} * 10 - 3 - int({{i}} / {{f}})`,
		`{{d2}} - {{d}}`,
		// if else
		`{{d}} * 10 if len({{s}}) > 10 else {{f}} / 10`,
	}

	for _, example := range exampleList {
		res, err := praser.Parse(example)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v = %v\n", example, res.Value)
	}
}
