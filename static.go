type Task struct {
    id   int
    name string
}

type Node struct {
    id       int
    capacity int
    load     int
}

func createTasks(n int) []Task {
    tasks := make([]Task, n)
    for i := 0; i < n; i++ {
        tasks[i] = Task{id: i, name: fmt.Sprintf("Task-%d", i)}
    }
    return tasks
}

func createNodes(n int, capacity int) []Node {
    nodes := make([]Node, n)
    for i := 0; i < n; i++ {
        nodes[i] = Node{id: i, capacity: capacity, load: 0}
    }
    return nodes
}

func assignTasksStatic(tasks []Task, nodes []Node) {
    for i, task := range tasks {
        nodeIndex := i % len(nodes)
        nodes[nodeIndex].load++
        fmt.Printf("Assigned %s to Node-%d\n", task.name, nodeIndex)
    }
    for _, node := range nodes {
        fmt.Printf("Node-%d load: %d capacity: %d\n", node.id, node.load, node.capacity)
    }
}

func simulateExecution(nodes []Node) {
    for _, node := range nodes {
        if node.load > node.capacity {
            fmt.Printf("Node-%d overloaded with load %d\n", node.id, node.load)
        } else if node.load == 0 {
            fmt.Printf("Node-%d idle\n", node.id)
        } else {
            fmt.Printf("Node-%d running tasks\n", node.id)
        }
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())
    tasks := createTasks(30)
    nodes := createNodes(4, 5)
    assignTasksStatic(tasks, nodes)
    simulateExecution(nodes)
}
