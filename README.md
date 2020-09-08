<!-- e.Static("/image", "profileimage")
	e.GET("/signin", controller.SignIn)
	e.GET("/showUserInfo", controller.ShowUserInfo)
	e.POST("/signup", controller.SignUp)
	e.POST("/updateProfileImage", controller.UpdateProfileImage, middleware.JWT([]byte("secret")), jwt.VerifyAccessToken) -->

# 아만다 (아무도 만나지 못한다)

## API - json, query, form 형식 모두 지원
아이디 중복검사
별점 등록
댓글 작성
답글 작성

- @GET 
    - /signin
        - ID (Require), PW (Require)
    - /showUserInfo
        - ID (Optional), Name (Optional)

- @POST
    - /signup
        - ID (Require) , PW (Require), Name (Require), IsManager (Optional), Description (Require)

    - /updateProfileImage
        - image (File, Require)