const input = document.querySelector("input")
const button = document.querySelector("button")
const ul = document.querySelector('ul')
const form = document.querySelector('form')

const todos = []

const BASE_URL = "http://localhost:3000"
const printList = () => {
    ul.innerHTML = todos.map(todo => `<li>${todo.name}</li>`).join('')
}

const getTodos = async () => {
    const res = await fetch(`${BASE_URL}/todos`)
    if (res.status === 200) {
        const data = await res.json()
        todos.push(...data)
        printList()
    }
}

(async () => {
    await getTodos()
})();

form.addEventListener("submit", async (e) => {
    e.preventDefault()
    const todo = {name: input.value}
    const res = await fetch(`${BASE_URL}/todos`, {
        method: "POST",
        headers: {
            'Content-Type': 'application/json',
            // 'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: JSON.stringify(todo)
    })
    if (res.status === 201) {
        todos.push(todo)
        printList()
    }
})