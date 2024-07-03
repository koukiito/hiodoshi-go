package services

import (
	"math/rand"
)

//Random [HI] [O] [DO] [SHI] [GO] e.g. HI O DO SHI GO
var charset[5] string = [5]string{"HI", "O", "DO", "SHI","GO"}

type HiodoshiService struct{}

func NewHiodoshiService() *HiodoshiService {
	return &HiodoshiService{}
}

func (s *HiodoshiService) GenerateRandomHiodoshi() string {
	var hiodoshi string
	for i := 0; i < 5; i++ {
		var index int = rand.Intn(len(charset))
		hiodoshi += charset[index]
	}
	return hiodoshi
}