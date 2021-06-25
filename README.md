# Proteus 

[![Stable Version](https://img.shields.io/github/v/tag/RobertBerglund/proteus?label=version)](https://img.shields.io/github/v/tag/RobertBerglund/proteus)
[![forthebadge](https://forthebadge.com/images/badges/works-on-my-machine.svg)](https://forthebadge.com)

Proteus is a tiny package for mapping values between structs using tags

## Install

```bash
go get github.com/RobertBerglund/proteus
```

## Examples

### Simple mapping 
```go
type Src struct {
	ValA string `dst:"Val1"`
	ValB string `dst:"Val2"`
}

type Dst struct {
	Val1 string
	Val2 string
}

src := Src{
		ValA: "value from A",
		ValB: "value from B",
	}
dst := Dst{}

proteus.New("dst").Map(src, &dst)

//{Val1:value from A Val2:value from B}
fmt.Printf("%+v", dst) 
```

### Simple mapping between structs with the same field names
```go
type SrcSimple struct {
	A string
	B string
}

type DstSimple struct {
	A string
	C string
}

src := SrcSimple{
	A: "a value",
	B: "b value",
}
dst := DstSimple{}
Map(src, &dst)
//dst.A == src.A
```

### Embedded structs

An empty tag value will "flatten" the struct

```go
type SrcEmbedded struct {
	Src `dst:""`
}

src := SrcEmbedded{}
src.ValA = "value from A"
src.ValB = "value from B"
dst := Dst{}

proteus.New("dst").Map(src, &dst)

//{Val1:value from A Val2:value from B}
fmt.Printf("%+v", dst) 
```

### Nested structs

```go
type SrcNested struct {
	SrcInner Src `dst:"DstInner"`
}

type DstNested struct {
	DstInner Dst
}

src := SrcNested{}
src.SrcInner.ValA = "value from A"
src.SrcInner.ValB = "value from B"
dst := DstNested{}

New("dst").Map(src, &dst)

//{DstInner:{Val1:value from A Val2:value from B}}
fmt.Printf("%+v", dst) 
```