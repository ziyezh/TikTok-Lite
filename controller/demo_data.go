package controller

var DemoVideos = []Video{
	{
		Id:            1,
		Author:        DemoUser,
		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         "Test Video",
	},
}

var DemoComments = []Comment{
	{
		Id:         1,
		User:       DemoUser,
		Content:    "Test Comment",
		CreateDate: "05-01",
	},
}

var DemoUser = User{
	Id:              1,
	Name:            "TestUser",
	FollowCount:     0,
	FollowerCount:   0,
	IsFollow:        false,
	Avatar:          "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
	BackgroundImage: "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
	Signature:       "Test Signature",
	TotalFavorited:  "0",
	WorkCount:       0,
	FavoriteCount:   0,
}
