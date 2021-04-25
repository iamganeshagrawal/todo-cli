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


### Work Progrss
- Setup Initial blank commands and sub commands
  - [X] version
  - [X] version --update
  - [X] add
  - [X] list
  - [X] list pending
  - [X] list completed
  - [X] delete
  - [X] mark

- Packages
  - [X] update : complated and linked github release tags for version check
  - [ ] database : initial file created
  - [X] version : completed



### SQL Snippits


```sql
# create table

CREATE TABLE IF NOT EXISTS "todos" (
	"id"	INTEGER,
	"task"	TEXT NOT NULL,
	"completed"	INTEGER NOT NULL DEFAULT 0 CHECK(completed IN (0,1)),
	"created_at"	TEXT DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY("id" AUTOINCREMENT)
);

```

```sql
# create index on completed field

CREATE INDEX IF NOT EXISTS "idx_todos_completed" ON "todos" (
	"completed"
);

```

```sql
# query all

SELECT id, task, completed FROM "todos";

# query completed

SELECT id, task, completed FROM todos WHERE completed=1;

# query pending

SELECT id, task, completed FROM todos WHERE completed=0;

```

```sql
# update completed status

UPDATE todos SET completed = 1 WHERE id = 2;

# completed = 1, pending = 0
```

```sql
# delete 

DELETE FROM todos WHERE id = 2;

```

```sql
# Insert

INSERT INTO todos (task) VALUES ("sdasdas dsa d as dsad sad");
select last_insert_rowid() as lastId;

```

```sql
# Count Total, Pending, Completed

SELECT count(id) as `total`, SUM(completed == 0) as `pending`, SUM(completed == 1) as `completed` from todos;

```