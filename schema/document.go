package schema

const (
	docMetaDataKeySubIndexes   = "_sub_indexes"
	docMetaDataKeyScore        = "_score"
	docMetaDataKeyExtraInfo    = "_extra_info"
	docMetaDataKeyDSL          = "_dsl"
	docMetaDataKeyDenseVector  = "_dense_vector"
	docMetaDataKeySparseVector = "_sparse_vector"
)

type Document struct {
	ID string `json:"id"`

	Content string `json:"content"`

	MetaData map[string]any `json:"meta_data"`
}

func (d *Document) String() string { _ = "STUB: not implemented"; return "" }

func (d *Document) WithSubIndexes(indexes []string) *Document {
	_ = "STUB: not implemented"
	return nil
}

func (d *Document) SubIndexes() []string { _ = "STUB: not implemented"; return nil }

func (d *Document) WithScore(score float64) *Document { _ = "STUB: not implemented"; return nil }

func (d *Document) Score() float64 { _ = "STUB: not implemented"; return 0 }

func (d *Document) WithExtraInfo(extraInfo string) *Document { _ = "STUB: not implemented"; return nil }

func (d *Document) ExtraInfo() string { _ = "STUB: not implemented"; return "" }

func (d *Document) WithDSLInfo(dslInfo map[string]any) *Document {
	_ = "STUB: not implemented"
	return nil
}

func (d *Document) DSLInfo() map[string]any { _ = "STUB: not implemented"; return nil }

func (d *Document) WithDenseVector(vector []float64) *Document {
	_ = "STUB: not implemented"
	return nil
}

func (d *Document) DenseVector() []float64 { _ = "STUB: not implemented"; return nil }

func (d *Document) WithSparseVector(sparse map[int]float64) *Document {
	_ = "STUB: not implemented"
	return nil
}

func (d *Document) SparseVector() map[int]float64 { _ = "STUB: not implemented"; return nil }
