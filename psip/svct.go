package psip

type svct struct {
	table
}

func newSVCT(payload []byte) VCT {
	return &svct{table(payload)}
}

func (vct svct) NumChannelsInSection() int {
	return 0
}

func (vct svct) Channels() []Channel {
	return nil
}

func (vct svct) AdditionalDescriptorsLength() uint16 {
	return 0
}
