package constant

type ProductTypeEnum string

const (
	ProductTypeCredit  ProductTypeEnum = "Credit"
	ProductTypePackage ProductTypeEnum = "Package"
)

func (x ProductTypeEnum) String() string {
	return string(x)
}
