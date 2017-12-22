package behavioral

import "fmt"

// Command is a command which when executed on a person will call a given method.
type Command struct {
	person *Person
	method func()
}

// NewCommand creates a new Command.
func NewCommand(person *Person, method func()) Command {
	return Command{person, method}
}

// Execute executes the command.
func (c *Command) Execute() {
	c.method()
}

// Person with a given name and command to execute.
type Person struct {
	name string
	cmd  Command
}

// NewPerson creates a new Person.
func NewPerson(name string, cmd Command) Person {
	return Person{name, cmd}
}

// Talk talks and the executes the follow on command.
func (p *Person) Talk() {
	fmt.Fprintf(outputWriter, "%s is talking.\n", p.name)
	p.cmd.Execute()
}

// PassOn passes on by executing the follow on command.
func (p *Person) PassOn() {
	fmt.Fprintf(outputWriter, "%s is passing on.\n", p.name)
	p.cmd.Execute()
}

// Gossip gossips and then executes follow on command.
func (p *Person) Gossip() {
	fmt.Fprintf(outputWriter, "%s is gossiping.\n", p.name)
	p.cmd.Execute()
}

// Listen listens without executing follow on command.
func (p *Person) Listen() {
	fmt.Fprintf(outputWriter, "%s is listening.\n", p.name)
}
