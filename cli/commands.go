package cli



// Store list of commands
var commands []*Command

type Command struct {
	Name string
	Cmd string
	Callback func(Context)
	Description string
}


// Define callback context
type Context struct {
	Name string
	Args Arguments
	Commands []*Command
}


type Arguments map[string]interface{}