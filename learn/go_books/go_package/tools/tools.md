0. 像牛人一样改进你的Go代码: https://colobu.com/2017/06/27/Lint-your-golang-code-like-a-mad-man/#gofmt
1. gofmt
    * 格式化代码
    * find . -name "\*.go" -not -path "./vendor/\*" -not -path ".git/\*" | xargs gofmt -s -d
    * vscode保存时自动会执行
2. gocyclo
    * 检查函数的复杂度
3. interfacer
    * 这个工具提供接口类型的建议，换句话说，它会对可以本没有必要定义成具体的类型的代码提出警告
4. deadcode
    * deadcode会告诉你哪些代码片段根本没用
