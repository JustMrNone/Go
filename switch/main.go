package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	whatOs()
	isItWeekday()
	isItNoon()
}

func whatOs() {
	fmt.Println("You are running:")
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		// freebsd, openbsd,
		// plan9, windows..
		fmt.Printf("%s.\n", os)
	}
}

// isItWeekday checks if today is a weekday or the weekend
func isItWeekday() {
	today := time.Now().Weekday()
	switch today {
	case time.Saturday, time.Sunday:
		// If today is Saturday or Sunday, it's the weekend
		fmt.Println("It's the Weekend")
	default:
		// Otherwise, it's a weekday
		fmt.Println("It's a weekday")
	}
}

// isItNoon checks if the current time is before or after noon
func isItNoon() {
	mytime := time.Now()
	switch {
	case mytime.Hour() < 12:
		// If the current hour is less than 12, it's before noon
		fmt.Println("It is before noon")
	default:
		// Otherwise, it's after noon
		fmt.Println("It is after noon")
	}
}

/*
In Go, the `switch` statement can be used in two primary ways:

1. **Expression-based `switch`:** This is where you specify a value or expression after the `switch` keyword, and the `case` statements compare against that value. For example:

    ```go
    switch myTime {
    case "morning":
        fmt.Println("Good morning!")
    case "evening":
        fmt.Println("Good evening!")
    default:
        fmt.Println("Hello!")
    }
    ```

    Here, `myTime` is evaluated, and the `case` blocks are matched against its value.

2. **Condition-based `switch`:** This is where you omit the value after the `switch` keyword. Instead, each `case` contains a condition that evaluates to `true` or `false`. The first `case` that evaluates to `true` is executed. For example:

    ```go
    switch {
    case myTime == "morning":
        fmt.Println("Good morning!")
    case myTime == "evening":
        fmt.Println("Good evening!")
    default:
        fmt.Println("Hello!")
    }
    ```

### Why omit the value in the second function?

When you omit the value after `switch`, it allows you to write more complex conditions in each `case`. This is useful when:

- You need to evaluate multiple unrelated conditions.
- The conditions are not directly tied to a single variable or expression.
- You want to use logical operators (e.g., `&&`, `||`) or more complex comparisons.

For example, if you want to check multiple conditions that aren't just about the value of `myTime`, the condition-based `switch` is more flexible:

```go
switch {
case myTime == "morning" && isWeekend:
    fmt.Println("Relax, it's a weekend morning!")
case myTime == "evening" && !isWeekend:
    fmt.Println("Prepare for work tomorrow!")
default:
    fmt.Println("Enjoy your day!")
}
```

In contrast, the expression-based `switch` is simpler and more concise when all cases depend on the same variable or expression.

### Summary

- Use `switch <expression>` when all cases depend on the value of a single variable or expression.
- Use `switch` (without an expression) when you need to evaluate multiple independent conditions.
*/
