## Rule Engine Example
example of [rule_engine](https://github.com/Uyouii/rule_engine)

```go
// integrate
4 * (2 + 3) - 5 * 3 = 5

// float
3.0 * (2.5 - 3) = -1.5

// logic
1 < 2 and 2 < 3 and 4.0 != 4.0001 = true
1 > 2 && 1 < 2 = false

// func exapmle
max(min(10.0, 20, 30), len("vstr")) = 10
min(len("test"), abs(-4.5), min(5,6)) = 4
upper("abc") == "ABC" = true
startWith("hello world", "hel") = true
int(97) + 3 == max(100, -1) = true
regexMatch("^0[xX][a-fA-F0-9]+", "0xasf4") = true
string({{d}}) == "3.3" = true

// use param
min(len({{s}}), {{i}}, {{f}}, {{d}}) = 3.3
int({{i}} / {{f}}) = 28
{{d}} * 10 - 3 - int({{i}} / {{f}}) = 2
{{d2}} - {{d}} = 0

// if else
{{d}} * 10 if len({{s}}) > 10 else {{f}} / 10 = 33
```