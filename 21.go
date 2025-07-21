package main

import "fmt"

// Паттерн Адаптер применяется в тех случаях, когда какой-то уже существующий класс описывает нужный
// функционал, и нам нужно применить его методы, но с другими сигнатурами методов.
// Реальный пример использования - совместимость со структурами различных библиотек с самописными
// решениями.

type NewController struct{}

func (c NewController) On() {
	fmt.Println("New controller set on")
}

func (c NewController) Off() {
	fmt.Println("New controller set off")
}

type OldController struct{}

func (c OldController) Activate() {
	fmt.Println("Old controller set on")
}

func (c OldController) Deactivate() {
	fmt.Println("Old controller set off")
}

type ControllerAdapter interface {
	On()
	Off()
}

type OldAdapter struct {
	*OldController
}

func (a OldAdapter) On() {
	a.Activate()
}

func (a OldAdapter) Off() {
	a.Deactivate()
}

func NewOldAdapter(controller *OldController) *OldAdapter {
	return &OldAdapter{OldController: controller}
}

func main() {
	newController := NewController{}
	oldController := OldController{}

	controllers := [2]ControllerAdapter{newController, NewOldAdapter(&oldController)}
	for _, controller := range controllers {
		controller.On()
		controller.Off()
	}
}
