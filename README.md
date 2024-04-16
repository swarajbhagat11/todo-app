# Todo API

- As discussed with Gurj Bath, I've developed this application based on my interpretation. I've retrieved the first 20 `Todo` items with even indices from https://jsonplaceholder.typicode.com/todos/{index}, displaying their status and title on the command line. Additionally, I've exposed an HTTP endpoint at http://localhost:8080/todos/{index} to access `Todo` items using their index. Furthermore, I've ensured comprehensive test coverage for all unit test cases and employed Docker to facilitate the execution of the application.

- I've utilized a variable to store the todo list, as there was no specification to employ a database, and it adequately met the requirements.

- I've adjusted the indexing to begin from 1 to match the user's perspective. Please note that when using the endpoint http://localhost:8080/todos/{index}, the index parameter starts from 1, which corresponds to the element at index 0 in the array.

- The logic is to retrieve even todos, thus I'm fetching todos from https://jsonplaceholder.typicode.com/todos/{index} where the index starts from 2, 4, 6, 8, 10,... up to 20 elements.

- After retrieving the todo's, I'm displaying the `Todo` title and status on the command line.
![Todo's title and status](https://github.com/swarajbhagat11/todo-app/blob/main/images/print_todo_status.png)

- You can retrieve an element using the following endpoint: http://localhost:8080/todos/{index}.

## Endpoints

#### API
- **/** `=>`  &ensp; **GET** &ensp;  `=>` &ensp; Get status of application
- **/todos/{index} &ensp;** `=>`  &ensp; **GET** &ensp;  `=>` &ensp; Fetch the todo item located at a specific index.

## Tools used

- `App configuration` &nbsp; **=>**  &nbsp; [Viper](https://github.com/spf13/viper)
- `Logging` &nbsp; **=>**  &nbsp; [Logrus](https://github.com/sirupsen/logrus)
- `Routing` &nbsp; **=>**  &nbsp;  [Chi](https://github.com/go-chi/chi)
- `Containerization` &nbsp; **=>**  &nbsp; [Docker](http://docker.com/) + Docker Compose
- `Testing` &nbsp; **=>**  &nbsp; [Mock](https://github.com/golang/mock) + [Testify](https://github.com/stretchr/testify)


## Running the app

```bash
# start the app
make run

# run unit tests
make test
```
