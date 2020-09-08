<!-- e.Static("/image", "profileimage")
	e.GET("/signin", controller.SignIn)
	e.GET("/showUserInfo", controller.ShowUserInfo)
	e.POST("/signup", controller.SignUp)
	e.POST("/updateProfileImage", controller.UpdateProfileImage, middleware.JWT([]byte("secret")), jwt.VerifyAccessToken) -->

# 아만다 (아무도 만나지 못한다)

## API - json, query, form 형식 모두 지원
유저의 별점, 참여자 수 보기
댓글 작성
답글 작성

- @GET 
    - /showMyInfo
        - JWT
        
    - /showUserInfo
        - id (Optional), name (Optional)
    
    - /showUserStar
        - idx (Require)

- @POST
    - /signup
        - id (Require) , pw (Require), name (Require), ismanager (Optional), description (Require)

    - /updateProfileImage
        - image (File, Require)
        - JWT

    - /signin
        - id (Require), pw (Require)

    - /updateStar
        - idx (Require), star (Require)
        - JWT
    
    - /updateMyInfo
        - name (Require), description (Require)
        - JWT 