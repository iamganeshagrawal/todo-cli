# todo-cli
CLI Based ToDo App (GoLang)

---


### commands

```zsh
# Add a todo
todo add something amazing  # task -> "something amazing"
todo add "Do testing on module client, server and db layer" # task -> "Do testing on module client, server and db layer"

# List all todo(s)
todo list

# List all pending todo(s)
todo list pending

# List all completed todo(s)
todo list completed

# Delete a todo by ID
todo delete 10      # ID -> 10

# Mark a todo completed by ID
todo mark 11        # ID -> 11

# Mark a todo pending/uncompleted by ID
todo mark 11 -u     # ID -> 11
```