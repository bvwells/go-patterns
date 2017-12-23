package behavioral

import "fmt"

type Element interface {
	Accept(v Visitor)
}

type This struct {
}

func (t *This) This() string {
	return "This"
}

func (t *This) Accept(v Visitor) {
	v.VisitThis(t)
}

type That struct {
}

func (t *That) That() string {
	return "That"
}

func (t *That) Accept(v Visitor) {
	v.VisitThat(t)
}

type TheOther struct {
}

func (t *TheOther) TheOther() string {
	return "TheOther"
}

func (t *TheOther) Accept(v Visitor) {
	v.VisitTheOther(t)
}

type Visitor interface {
	VisitThis(e *This)
	VisitThat(e *That)
	VisitTheOther(e *TheOther)
}

type UpVisitor struct {
}

func (v *UpVisitor) VisitThis(e *This) {
	fmt.Printf("do Up on %v\n", e.This())
}

func (v *UpVisitor) VisitThat(e *That) {
	fmt.Printf("do Up on %v\n", e.That())
}

func (v *UpVisitor) VisitTheOther(e *TheOther) {
	fmt.Printf("do Up on %v\n", e.TheOther())
}

type DownVisitor struct {
}

func (v *DownVisitor) VisitThis(e *This) {
	fmt.Printf("do Down on %v\n", e.This())
}

func (v *DownVisitor) VisitThat(e *That) {
	fmt.Printf("do Down on %v\n", e.That())
}

func (v *DownVisitor) VisitTheOther(e *TheOther) {
	fmt.Printf("do Down on %v\n", e.TheOther())
}
