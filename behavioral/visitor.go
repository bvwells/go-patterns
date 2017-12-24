package behavioral

import "fmt"

// Element defines an interface for accepting visitors.
type Element interface {
	Accept(v Visitor)
}

// This defines a struct which is an element.
type This struct {
}

// This returns 'This' as a string.
func (t *This) This() string {
	return "This"
}

// Accept accepts a visitor.
func (t *This) Accept(v Visitor) {
	v.VisitThis(t)
}

// That defines a struct which is an element.
type That struct {
}

// That returns 'That' as a string.
func (t *That) That() string {
	return "That"
}

// Accept accepts a visitor.
func (t *That) Accept(v Visitor) {
	v.VisitThat(t)
}

// TheOther defines a struct which is an element.
type TheOther struct {
}

// TheOther returns 'TheOther' as a string./
func (t *TheOther) TheOther() string {
	return "TheOther"
}

// Accept accepts a visitor.
func (t *TheOther) Accept(v Visitor) {
	v.VisitTheOther(t)
}

// Visitor defines an interface for visiting this, that and the other.
type Visitor interface {
	VisitThis(e *This)
	VisitThat(e *That)
	VisitTheOther(e *TheOther)
}

// UpVisitor defines an up visitor.
type UpVisitor struct {
}

// VisitThis visits this.
func (v *UpVisitor) VisitThis(e *This) {
	fmt.Printf("do Up on %v\n", e.This())
}

// VisitThat visits that.
func (v *UpVisitor) VisitThat(e *That) {
	fmt.Printf("do Up on %v\n", e.That())
}

// VisitTheOther visits the other.
func (v *UpVisitor) VisitTheOther(e *TheOther) {
	fmt.Printf("do Up on %v\n", e.TheOther())
}

// DownVisitor defines a down visitor.
type DownVisitor struct {
}

// VisitThis visits this.
func (v *DownVisitor) VisitThis(e *This) {
	fmt.Printf("do Down on %v\n", e.This())
}

// VisitThat visits that.
func (v *DownVisitor) VisitThat(e *That) {
	fmt.Printf("do Down on %v\n", e.That())
}

// VisitTheOther visits the other.
func (v *DownVisitor) VisitTheOther(e *TheOther) {
	fmt.Printf("do Down on %v\n", e.TheOther())
}
