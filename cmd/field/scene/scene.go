package scene

import rl "github.com/gen2brain/raylib-go/raylib"

type Properties struct {
}

type Object struct {
	Ref *any
	Properties
}

func NewObject(ref any) {

}

type OType interface {
	rl.Rectangle | rl.Camera2D
}

func (o *Object) Get[OType any](OType) OType {
	return (*o.Ref).(OType)
}

func (o *Object) GetRectangle() *rl.Rectangle {
	return (*o.Ref).(*rl.Rectangle)
}

func (o *Object) GetCamera2D() *rl.Camera2D {
	return (*o.Ref).(*rl.Camera2D)
}

type Scene struct {
	GlobalObjectTable map[string]*Object
}

func NewScene() *Scene {
	s := new(Scene)
	return s
}

func (s *Scene) AddObject(nName string, object any) {

}

func (s *Scene) DelObject(nName string) {

}
