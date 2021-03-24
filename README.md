# Proteus
Proteus is a tiny package for mapping values between structs using tags

## Examples

### Simple mapping 
```
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

### Embedded structs

An empty tag value will "flatten" the struct

```
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

```
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