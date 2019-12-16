package day8

import "fmt"

// Img TODO godoc
type Img struct {
	Width  int
	Height int
	Raw    string
	Layers [][]byte // each layer is Width x Height ints; there are N layers
}

// New TODO godoc
func New(img string, w int, h int) Img {

	i := Img{
		Width:  w,
		Height: h,
		Raw:    img,
	}

	l := make([][]byte, 0)

	b := toBytes(img)

	layerLength := w * h

	for start, i := 0, 0; start < len(b); start += layerLength {
		layer := make([]byte, layerLength)
		copy(layer, b[start:(start+layerLength)])
		l = append(l, layer)
		i++
	}

	i.Layers = l
	return i
}

func toBytes(src string) []byte {

	srcBytes := []byte(src)
	dst := make([]byte, len(srcBytes))
	copy(dst, srcBytes)

	for i := 0; i < len(dst); i++ {
		dst[i] -= 48
	}

	return dst
}

// ResolveLayers TODO godoc
func (img Img) ResolveLayers() []byte {
	final := make([]byte, img.Width*img.Height)
	for i := range final {
		final[i] = 99
	}

	for _, layer := range img.Layers {
		for i := 0; i < len(layer); i++ {
			// if the final bit is NOT set yet
			if final[i] == 99 {
				// If the current layer has a non-transparent pixel
				if layer[i] != 2 {
					final[i] = layer[i]
				}
			}
		}
	}

	return final
}

// FindLayerWithFewest0 TODO godoc
func (img Img) FindLayerWithFewest0() []byte {

	var layer []byte
	var minLayer []byte
	min := 99

	for i := 0; i < len(img.Layers); i++ {
		layer = img.Layers[i]
		n := count(layer, 0)
		fmt.Printf("Layer %d has %d zeros\n", i, n)
		if n < min {
			min = n
			minLayer = layer
		}
	}

	return minLayer
}

func count(layer []byte, val byte) int {

	n := 0
	for _, b := range layer {
		if b == val {
			n++
		}
	}

	return n
}
