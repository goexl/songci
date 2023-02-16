package songci

// Songci 对外主接口
type Songci struct {
	params *params
}

func newSongci(params *params) *Songci {
	return &Songci{
		params: params,
	}
}

func (s *Songci) Zinan(secret string) *zinanBuilder {

}
