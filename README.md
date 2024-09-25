## Linked List Implementation in Go

This repository contains a simple implementation of a singly linked list in Golang.

### Features

-   **Insertion:** Add new nodes at the beginning, end, or at a specific position.
-   **Deletion:** Remove nodes from the beginning, end, or by value.
-   **Search:** Find a node containing a specific value.
-   **Traversal:** Iterate over the linked list and print its contents.

### How to Use

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/your-username/linkedlist-go.git
    ```

2.  **Navigate to the project directory:**

    ```bash
    cd linkedlist-go
    ```

3.  **Import the linked list package:**

    ```go
    import "your-module-path/linkedlist" 
    ```

    (replace `your-module-path` with the actual path to your module)
4.  **Create a new linked list and perform operations:**

    ```go
    package main

    import (
    	"fmt"
        "your-module-path/linkedlist"
    )

    func main() {
    	list := linkedlist.NewLinkedList()
    	list.InsertAtBeginning(10)
    	list.InsertAtEnd(20)
    	// ... other operations
    	list.PrintList() 
    }
    ```

### Examples

Refer to the `main.go` file for example usage of the linked list implementation.

### Contributing

Contributions are welcome! Feel free to open issues or pull requests.

### License

This project is licensed under the MIT License. 

