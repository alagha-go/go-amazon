package variables


var (
	// action variables
	CreateDepartment = "create department"
	GetDepartments = "get departments"
	GetCategories = "get categories"
	GetSubCategories = "get subcategories"
	GetTypes = "get types"
	GetSubTypes = "get subtypes"
	
	// error variables
	InternalServerError = "internal server error"
	TypeNotFound = "type with the id %s does not exist"
	DepartMentExists = "department with the title %s exists"
	DepartmentNotFound = "department with the id %s does not exist"
	CategoryExists = "category with the title %s exists"
	CategoryNotFound = "category with the id %s does not exist"
	SubCategoryExists = "subcategory with the title %s exists"
	SubCategoryNotFound = "subcategory with the id %s does not exist"
	TypeExists = "type with the title %s exists"
	InvalidID = "invalid id"
)