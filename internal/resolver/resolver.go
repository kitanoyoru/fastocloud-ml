package resolver

import (
	"gonum.org/v1/gonum/mat"

	"gitlab.com/fastogt/machine-learning/face_detection/golang/pkg/utils"
)

type IResolver interface {
	Resolve(data []float32) (*string, float32)
}

type FaceResolver struct {
	embedings []*mat.Dense
	names     []string
}

func NewFaceResolver(emb [][]float32, names []string) *FaceResolver {
	embeddings := make([]*mat.Dense, 0, len(emb))

	for _, v := range emb {
		m := mat.NewDense(1, len(v), utils.ConvertToFloat64Slice(v))
		embeddings = append(embeddings, m)
	}

	return &FaceResolver{
		embedings: embeddings,
		names:     names,
	}
}

func (fr *FaceResolver) Resolve(data []float32) (*string, float32) {
	matdata := mat.NewDense(1, len(data), utils.ConvertToFloat64Slice(data))
	denom := mat.Norm(matdata, 2)

	for i := 0; i < len(data); i++ {
		v := matdata.At(0, i) / denom
		matdata.Set(0, i, v)
	}

	dist := []float64{}

	for _, emb := range fr.embedings {
		temp := mat.NewDense(1, len(data), nil)
		temp.Sub(matdata, emb)
		distItem := mat.Norm(temp, 2)
		dist = append(dist, distItem)
	}

	mind, mindi := utils.GetMinNumInSlice(dist)

	if mind < 1.1 {
		return &fr.names[mindi], float32(mind)
	} else {
		return nil, float32(mind)
	}
}
