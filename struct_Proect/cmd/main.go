package main

import (
	"fmt"
	"main/struct_Proect/pkg/post"
	"main/struct_Proect/pkg/repo/post_repo"
	"main/struct_Proect/pkg/repo/user_repo"
	"main/struct_Proect/pkg/user"
)

func main() {
	us := user_repo.NewUserStorage()
	ps := post_repo.NewPostStorage()

	tony := user.NewUser("Anthony")
	jhonny := user.NewUser("Jhonny Sack")

	tony.MakeFriend(jhonny)
	p := post.NewPost(tony.Name, "some post content")

	ps.Put(p)
	us.Put(tony)
	us.Put(jhonny)

	for _, users := range us {
		fmt.Println(users)
	}

	for _, posts := range ps {
		fmt.Println(posts)
	}

}
