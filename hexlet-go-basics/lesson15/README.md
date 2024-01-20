Реализуйте функцию func Map(strs []string, mapFunc func(s string) string) []string, которая преобразует каждый элемент слайса strs с помощью функции mapFunc и возвращает новый слайс. 
Учтите, что исходный слайс, который передается как strs, не должен измениться в процессе выполнения.

This example work on golang 1.18
go install golang.org/dl/go1.18.0@latest
/home/maksim/go/bin/go1.18 download

for neovim before start
export GOROOT=/home/maksim/sdk/go1.18/
export PATH=$GOROOT/bin:$PATH

