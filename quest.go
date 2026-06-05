package main

import "fmt"

type Question struct {
	Text    string
	Options []string
	Answer  string
}
type Difficulty struct {
	Beginner     []Game
	Intermediate []Game
	Advanced     []Game
	Expert       []Game
}
type Game struct {
	Level    string
	Amount   string
	Category string
	Quiz     Question
}

func quests() {
	database := Difficulty{
		Beginner: []Game{
			{Level: "1", Amount: "$10", Category: "Basic Syntax", Quiz: Question{
				Text:    "What is the only loop keyword natively available in Go?",
				Options: []string{"while", "for", "loop", "foreach"}, Answer: "B",
			}},
			{Level: "2", Amount: "$20", Category: "Packages", Quiz: Question{
				Text:    "How do you declare the entry-point executable package layout in Go?",
				Options: []string{"import main", "package main", "pkg main", "define main"}, Answer: "B",
			}},
			{Level: "3", Amount: "$30", Category: "Variables", Quiz: Question{
				Text:    "Which symbol block represents a short variable declaration syntax shorthand?",
				Options: []string{"=", "==", ":=", "->"}, Answer: "C",
			}},
			{Level: "4", Amount: "$40", Category: "Basic Types", Quiz: Question{
				Text:    "What is the default zero-value assigned to an uninitialized integer variable?",
				Options: []string{"null", "0", "undefined", "-1"}, Answer: "B",
			}},
			{Level: "5", Amount: "$50", Category: "Control Flow", Quiz: Question{
				Text:    "Which keyword delays function execution sequence actions until the wrapping block ends?",
				Options: []string{"delay", "later", "defer", "wait"}, Answer: "C",
			}},
			{Level: "6", Amount: "$60", Category: "Strings", Quiz: Question{
				Text:    "What underlying primitive type elements are evaluated when iterating over a Go string with a standard range loop?",
				Options: []string{"byte", "rune", "string", "int"}, Answer: "B",
			}},
			{Level: "7", Amount: "$70", Category: "Arrays", Quiz: Question{
				Text:    "Which statement best describes a Go array compared to a Go slice?",
				Options: []string{"Arrays have dynamic length", "Arrays can grow with append()", "Arrays have a fixed size defined at compile time", "Arrays are references"}, Answer: "C",
			}},
			{Level: "8", Amount: "$80", Category: "Booleans", Quiz: Question{
				Text:    "What is the default initial state value of an uninitialized boolean variable?",
				Options: []string{"true", "false", "nil", "0"}, Answer: "B",
			}},
			{Level: "9", Amount: "$90", Category: "Constants", Quiz: Question{
				Text:    "Which keyword defines a fixed, unchangeable value evaluated during compilation?",
				Options: []string{"var", "let", "const", "immutable"}, Answer: "C",
			}},
			{Level: "10", Amount: "$100", Category: "Operators", Quiz: Question{
				Text:    "Which symbol combination is used to perform a logical 'AND' operation in conditional expressions?",
				Options: []string{"&", "&&", "and", "|"}, Answer: "B",
			}},
		},

		Intermediate: []Game{
			{Level: "1", Amount: "$10", Category: "Functions", Quiz: Question{
				Text:    "Can a standard Go function return multiple discrete values?",
				Options: []string{"No, never", "Yes, but only up to 2 items", "Yes, an arbitrary number of values can be returned", "Only if bundled inside an array"}, Answer: "C",
			}},
			{Level: "2", Amount: "$20", Category: "Structs", Quiz: Question{
				Text:    "How can you declare an anonymous struct instance block on the fly?",
				Options: []string{"struct{}", "var s object", "struct{ Name string }{}", "new struct"}, Answer: "C",
			}},
			{Level: "3", Amount: "$30", Category: "Methods", Quiz: Question{
				Text:    "What parameter defines the connection of a specific method block onto a target Struct type?",
				Options: []string{"The arguments list", "The interface payload", "The receiver parameter", "The struct header"}, Answer: "C",
			}},
			{Level: "4", Amount: "$40", Category: "Errors", Quiz: Question{
				Text:    "What is the standard built-in interface type signature used for error signaling returns?",
				Options: []string{"Exception", "ErrorClass", "error", "panic"}, Answer: "C",
			}},
			{Level: "5", Amount: "$50", Category: "Closures", Quiz: Question{
				Text:    "What unique functional state behavior is supported by Go's anonymous functions?",
				Options: []string{"They cannot accept arguments", "They can capture and reference variables outside their body", "They run on isolated CPUs", "They must have names"}, Answer: "B",
			}},
			{Level: "6", Amount: "$60", Category: "Slices", Quiz: Question{
				Text:    "What native built-in function dynamically expands slice capacity memory segments?",
				Options: []string{"push()", "add()", "extend()", "append()"}, Answer: "D",
			}},
			{Level: "7", Amount: "$70", Category: "Maps", Quiz: Question{
				Text:    "What happens when you look up a non-existent key element inside an initialized Go map?",
				Options: []string{"It triggers a panic", "It causes a compilation error", "It returns the value type's zero value", "It returns nil for all types"}, Answer: "C",
			}},
			{Level: "8", Amount: "$80", Category: "Variadic", Quiz: Question{
				Text:    "What operator notation prefix marks a function argument slice to accept variable argument counts?",
				Options: []string{"...", "*", "[]", "args"}, Answer: "A",
			}},
			{Level: "9", Amount: "$90", Category: "Imports", Quiz: Question{
				Text:    "How can you import a package to execute its init() function without using its exported code identifiers?",
				Options: []string{"import . \"pkg\"", "import _ \"pkg\"", "import hide \"pkg\"", "import pkg"}, Answer: "B",
			}},
			{Level: "10", Amount: "$100", Category: "Struct Tags", Quiz: Question{
				Text:    "Where are struct tags placed to provide metadata information instructions to encoders like JSON?",
				Options: []string{"Before the struct declaration", "Inside backticks after the field type definition", "Inside comment blocks", "Inside parentheses"}, Answer: "B",
			}},
		},

		Advanced: []Game{
			{Level: "1", Amount: "$10", Category: "Pointers", Quiz: Question{
				Text:    "Which operator is used to generate a direct memory address pointer reference pointing to an existing variable data source?",
				Options: []string{"*", "&", "->", "new"}, Answer: "B",
			}},
			{Level: "2", Amount: "$20", Category: "Dereferencing", Quiz: Question{
				Text:    "Which operator acts as the de-referencing modifier symbol to read or change memory data behind a pointer?",
				Options: []string{"&", "#", "*", "->"}, Answer: "C",
			}},
			{Level: "3", Amount: "$30", Category: "Interfaces", Quiz: Question{
				Text:    "How does a user-defined concrete Struct type implement an interface in the Go ecosystem?",
				Options: []string{"By explicitly adding an 'implements' keyword statement", "Implicitly, by writing methods matching the exact interface signature blueprint", "By registering it inside the main init routine", "Through explicit package class inheritance"}, Answer: "B",
			}},
			{Level: "4", Amount: "$40", Category: "Empty Interface", Quiz: Question{
				Text:    "What specialized utility signature is represented by an empty interface type (interface{})?",
				Options: []string{"It can hold absolutely any data value type", "It acts as a null data breaker", "It flags empty memory states", "It forces errors"}, Answer: "A",
			}},
			{Level: "5", Amount: "$50", Category: "Type Assertions", Quiz: Question{
				Text:    "What native syntax evaluates if an interface variable holds a specific target concrete structural type layout?",
				Options: []string{"typeof(variable)", "variable.(TargetType)", "variable == TargetType", "variable.is(TargetType)"}, Answer: "B",
			}},
			{Level: "6", Amount: "$60", Category: "Receiver Types", Quiz: Question{
				Text:    "What happens when you use a value receiver method instead of a pointer receiver method?",
				Options: []string{"Changes modify the original caller data instance state", "The method operates on a copy, leaving the original data unchanged", "The application runs faster", "Compiler errors break execution builds"}, Answer: "B",
			}},
			{Level: "7", Amount: "$70", Category: "Composition", Quiz: Question{
				Text:    "How does Go achieve reuse patterns without classic subclassing architectures?",
				Options: []string{"Using virtual methods", "Using struct embedding fields without names", "Through global namespaces", "Using interface templates"}, Answer: "B",
			}},
			{Level: "8", Amount: "$80", Category: "Constructor Idioms", Quiz: Question{
				Text:    "What is the standard idiomatic naming layout for functions acting as structural allocation constructors?",
				Options: []string{"createStruct()", "initStruct()", "NewStruct()", "makeStruct()"}, Answer: "C",
			}},
			{Level: "9", Amount: "$90", Category: "Type Switches", Quiz: Question{
				Text:    "What switch syntax construct allows conditional logic processing based on the dynamic type of an interface variable?",
				Options: []string{"switch v.Type()", "switch v.(type)", "switch type(v)", "switch typeof(v)"}, Answer: "B",
			}},
			{Level: "10", Amount: "$100", Category: "Nil Pointers", Quiz: Question{
				Text: "What occurs if your application tries to access or alter struct properties through a nil pointer reference directly?",

				Options: []string{"It returns a zero value", "It safely returns nil", "It triggers a runtime panic crash", "It fails at compilation phase"}, Answer: "C",
			}},
		},

		Expert: []Game{
			{Level: "1", Amount: "$10", Category: "Goroutines", Quiz: Question{
				Text:    "What native prefix expression schedules an asynchronous concurrent task thread lifecycle inside Go?",
				Options: []string{"thread", "go", "async", "spawn"}, Answer: "B",
			}},
			{Level: "2", Amount: "$20", Category: "Channels", Quiz: Question{
				Text:    "Which syntax notation sends data values into an active Go pipeline communication channel?",
				Options: []string{"channel <- data", "data -> channel", "channel.Send(data)", "send(channel, data)"}, Answer: "A",
			}},
			{Level: "3", Amount: "$30", Category: "Select Blocks", Quiz: Question{
				Text:    "What structural statement construct manages orchestration paths waiting across multiple target channel operations?",
				Options: []string{"switch", "select", "choose", "wait"}, Answer: "B",
			}},
			{Level: "4", Amount: "$40", Category: "Buffered Channels", Quiz: Question{
				Text:    "What is the primary difference when using a buffered channel compared to an unbuffered channel?",
				Options: []string{"It only accepts string data types", "Sends block immediately until a receiver reads the message", "Sends do not block until the buffer capacity limit is reached", "It runs automatically on independent CPU hardware pools"}, Answer: "C",
			}},
			{Level: "5", Amount: "$50", Category: "Channel Closing", Quiz: Question{
				Text:    "What execution event happens if your program mistakenly sends a new message into a channel that has already been closed?",
				Options: []string{"The data message drops silently", "The channel blocks execution infinitely", "It instantly triggers a runtime panic exception", "The system returns a default zero value"}, Answer: "C",
			}},
			{Level: "6", Amount: "$60", Category: "WaitGroups", Quiz: Question{
				Text:    "Which synchronization method from 'sync.WaitGroup' decrements the internal counter by 1?",
				Options: []string{"Add(-1)", "Done()", "Wait()", "Both A and B are correct"}, Answer: "D",
			}},
			{Level: "7", Amount: "$70", Category: "Mutexes", Quiz: Question{
				Text:    "Which package structure protects shared memory maps against fatal concurrent data race crashes?",
				Options: []string{"sync.Mutex", "sync.Pool", "runtime.Lock", "os.Signal"}, Answer: "A",
			}},
			{Level: "8", Amount: "$80", Category: "Panics", Quiz: Question{
				Text:    "What built-in system capability must be executed within a deferred call pattern block to catch a runtime panic state?",
				Options: []string{"catch()", "recover()", "handle()", "stop()"}, Answer: "B",
			}},
			{Level: "9", Amount: "$90", Category: "Unsafe Package", Quiz: Question{
				Text:    "Which specialized native library module lets developers circumvent Go type safety boundary layer rules?",
				Options: []string{"reflect", "runtime", "unsafe", "syscall"}, Answer: "C",
			}},
			{Level: "10", Amount: "$100", Category: "Garbage Collection", Quiz: Question{
				Text:    "What architectural methodology describes Go's native production garbage collection runtime process engine?",
				Options: []string{"Reference Counting", "Concurrent Tri-color Mark-and-Sweep", "Generational Stop-the-world", "Manual allocation clearing"}, Answer: "B",
			}},
		},
	}
	fmt.Println(database)
}
