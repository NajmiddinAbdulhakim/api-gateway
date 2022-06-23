package model


type User struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	FirstName            string     `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name"`
	LastName             string     `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name"`
	UserName             string     `protobuf:"bytes,4,opt,name=user_name,json=userName,proto3" json:"user_name"`
	Email                string     `protobuf:"bytes,5,opt,name=email,proto3" json:"email"`
	PhoneNumber          []string   `protobuf:"bytes,6,rep,name=phone_number,json=phoneNumber,proto3" json:"phone_number"`
	Addresses            []*Address `protobuf:"bytes,7,rep,name=addresses,proto3" json:"addresses"`
	Posts                []*Post    `protobuf:"bytes,8,rep,name=posts,proto3" json:"posts"`
	Bio                  string     `protobuf:"bytes,9,opt,name=bio,proto3" json:"bio"`
	Status               string     `protobuf:"bytes,10,opt,name=status,proto3" json:"status"`
	CreatedAt            string     `protobuf:"bytes,11,opt,name=createdAt,proto3" json:"createdAt"`
	UpdatedAt            string     `protobuf:"bytes,12,opt,name=updatedAt,proto3" json:"updatedAt"`
	DeletedAt            string     `protobuf:"bytes,13,opt,name=deletedAt,proto3" json:"deletedAt"`
}

type GetUser struct {
	userID string `json:"user_id`
}

type Address struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Country              string   `protobuf:"bytes,2,opt,name=country,proto3" json:"country"`
	City                 string   `protobuf:"bytes,3,opt,name=city,proto3" json:"city"`
	District             string   `protobuf:"bytes,4,opt,name=district,proto3" json:"district"`
	PostalCode           string   `protobuf:"bytes,5,opt,name=postal_code,json=postalCode,proto3" json:"postal_code"`
}
type Post struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description"`
	Medias               []*Media `protobuf:"bytes,5,rep,name=medias,proto3" json:"medias"`
	CreatedAt            string   `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt"`
	UpdatedAt            string   `protobuf:"bytes,7,opt,name=updatedAt,proto3" json:"updatedAt"`
	DeletedAt            string   `protobuf:"bytes,8,opt,name=deletedAt,proto3" json:"deletedAt"`
}
type Media struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	PostId               string   `protobuf:"bytes,2,opt,name=post_id,json=postId,proto3" json:"post_id"`
	Link                 string   `protobuf:"bytes,3,opt,name=link,proto3" json:"link"`
	Type                 string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type"`
}

type UpdateUserReq struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	FirstName            string     `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name"`
	LastName             string     `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name"`
	UserName             string     `protobuf:"bytes,4,opt,name=user_name,json=userName,proto3" json:"user_name"`
	Email                string     `protobuf:"bytes,5,opt,name=email,proto3" json:"email"`
	PhoneNumber          []string   `protobuf:"bytes,6,rep,name=phone_number,json=phoneNumber,proto3" json:"phone_number"`
	Addresses            []*Address `protobuf:"bytes,7,rep,name=addresses,proto3" json:"addresses"`
	Bio                  string     `protobuf:"bytes,8,opt,name=bio,proto3" json:"bio"`
	Status               string     `protobuf:"bytes,9,opt,name=status,proto3" json:"status"`
}

type RegisterUser struct {
	Id                 	 string		`json:"id"`
	FirstName            string     `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name"`
	LastName             string     `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name"`
	UserName             string     `protobuf:"bytes,4,opt,name=user_name,json=userName,proto3" json:"user_name"`
	Email                string     `protobuf:"bytes,5,opt,name=email,proto3" json:"email"`
	Password 			 string 	`json:"password"`
	PhoneNumber          []string   `protobuf:"bytes,6,rep,name=phone_number,json=phoneNumber,proto3" json:"phone_number"`
	Bio                  string     `protobuf:"bytes,9,opt,name=bio,proto3" json:"bio"`
	Status               string     `protobuf:"bytes,10,opt,name=status,proto3" json:"status"`      
	Code 				 int64      `json:"code"`
	RefreshToken 		 string 	`json:"refresh_token"`
}

type RegisterUserRes struct {
	Id string `json:"id"`
	RefreshToken string `json:"refresh_token"`
	AccessToken string `json:"access_token"`
}




// type User struct {
// 	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
// 	FirstName            string     `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name"`
// 	LastName             string     `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name"`
// 	UserName             string     `protobuf:"bytes,4,opt,name=user_name,json=userName,proto3" json:"user_name"`
// 	Email                string     `protobuf:"bytes,5,opt,name=email,proto3" json:"email"`
// 	PhoneNumber          []string   `protobuf:"bytes,6,rep,name=phone_number,json=phoneNumber,proto3" json:"phone_number"`
// 	Addresses            []*Address `protobuf:"bytes,7,rep,name=addresses,proto3" json:"addresses"`
// 	Posts                []*Post    `protobuf:"bytes,8,rep,name=posts,proto3" json:"posts"`
// 	Bio                  string     `protobuf:"bytes,9,opt,name=bio,proto3" json:"bio"`
// 	Status               string     `protobuf:"bytes,10,opt,name=status,proto3" json:"status"`
// 	CreatedAt            string     `protobuf:"bytes,11,opt,name=createdAt,proto3" json:"createdAt"`
// 	UpdatedAt            string     `protobuf:"bytes,12,opt,name=updatedAt,proto3" json:"updatedAt"`
// 	DeletedAt            string     `protobuf:"bytes,13,opt,name=deletedAt,proto3" json:"deletedAt"`
// }
// type Address struct {
// 	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
// 	Country              string   `protobuf:"bytes,2,opt,name=country,proto3" json:"country"`
// 	City                 string   `protobuf:"bytes,3,opt,name=city,proto3" json:"city"`
// 	District             string   `protobuf:"bytes,4,opt,name=district,proto3" json:"district"`
// 	PostalCode           string   `protobuf:"bytes,5,opt,name=postal_code,json=postalCode,proto3" json:"postal_code"`
// }
// type Post struct {
// 	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
// 	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id"`
// 	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title"`
// 	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description"`
// 	Medias               []*Media `protobuf:"bytes,5,rep,name=medias,proto3" json:"medias"`
// 	CreatedAt            string   `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt"`
// 	UpdatedAt            string   `protobuf:"bytes,7,opt,name=updatedAt,proto3" json:"updatedAt"`
// 	DeletedAt            string   `protobuf:"bytes,8,opt,name=deletedAt,proto3" json:"deletedAt"`
// }
// type Media struct {
// 	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
// 	PostId               string   `protobuf:"bytes,2,opt,name=post_id,json=postId,proto3" json:"post_id"`
// 	Link                 string   `protobuf:"bytes,3,opt,name=link,proto3" json:"link"`
// 	Type                 string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type"`
// }


