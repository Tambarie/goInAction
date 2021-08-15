//package methods
//
////1. User defined types
//
//type User struct {
//	name string
//	email string
//	ext        int
//	privileged bool
//}
//
//// Admin Declaring fields based on other struct types
//// Admin represents an admin user with privileges
//type Admin struct {
//	User
//	level string
//}
//
//
//func main()  {
//// DECLARATION OF A VARIABLE OF THE STRUCT TYPE USING A STRUCT LITERAL
//	lisa := User{"Lisa",
//				"lisa@email.com",
//				123,
//				true}
//
//	//FOR ADMIN
//	fred := Admin{User{"Fred",
//		"fred@gmail.com",
//		123,
//		true},
//		"super"}
//}