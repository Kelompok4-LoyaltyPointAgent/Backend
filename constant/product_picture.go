package constant

type ProductPictureTypeEnum string

const (
	ProductPictureTypeIcon  ProductPictureTypeEnum = "Icon"
	ProductPictureTypePhoto ProductPictureTypeEnum = "Photo"
)

func (x ProductPictureTypeEnum) String() string {
	return string(x)
}
