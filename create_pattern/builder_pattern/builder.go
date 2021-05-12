/**
 *@Description
 *@ClassName builder
 *@Date 2021/5/12 下午9:58
 *@Author ckhero
 */

package builder_pattern

import "fmt"

// 屏幕
type Screen interface {
	Desc() string
}

type Screen15 struct {
}

func (Screen15) Desc() string {
	return "15寸"
}

type Screen13 struct {
}

func (Screen13) Desc() string {
	return "13寸"
}


type Color interface {
	Desc() string
}

type ColorRed struct {
}

func (ColorRed) Desc() string {
	return "红色"
}

type ColorWhite struct {
}

func (ColorWhite) Desc() string {
	return "白色"
}

type Computer struct {
	screen Screen
	color Color
}
func(c *Computer) GetDetail() string {
	return fmt.Sprintf("%s - %s", c.screen.Desc() , c.color.Desc())
}
// 建造则

type IBuilder interface {
	SetScreen(screen Screen) IBuilder
	SetColor(color Color) IBuilder
	Build() *Computer
}

type Builder struct {
	computer *Computer
}

func NewBuilder() IBuilder {
	return &Builder{computer:&Computer{}}
}

func(b *Builder) SetScreen(screen Screen) IBuilder {
	b.computer.screen = screen
	return b
}

func(b *Builder) SetColor(color Color) IBuilder {
	b.computer.color = color
	return b
}

func(b *Builder) Build() *Computer {
	return b.computer
}

// 调用者

type Director struct {
	build IBuilder
}

func(d *Director) Create01() *Computer {
	return d.build.SetColor(&ColorRed{}).SetScreen(&Screen13{}).Build()
}

func(d *Director) Create02() *Computer {
	return d.build.SetColor(&ColorWhite{}).SetScreen(&Screen15{}).Build()
}