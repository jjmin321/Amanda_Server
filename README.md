<!-- e.Static("/image", "profileimage")
	e.GET("/signin", controller.SignIn)
	e.GET("/showUserInfo", controller.ShowUserInfo)
	e.POST("/signup", controller.SignUp)
	e.POST("/updateProfileImage", controller.UpdateProfileImage, middleware.JWT([]byte("secret")), jwt.VerifyAccessToken) -->

# 아만다 (아무도 만나지 못한다)

## API - json, query, form 형식 모두 지원
답글 작성
답글 수정
댓글, 답글 보여주기

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
    
    - /createComment
        - id (작성할 대상의 아이디, Require), comment (작성할 댓글, Require)
        - JWT
        
    - /updateComment
        - idx (수정할 댓글의 인덱스, Require), comment (수정할 댓글, Require)
        - JWT

    - /createReplyComment
        - idx (답글을 달 댓글의 인덱스, Require), comment(작성할 답글, Require)
        - JWT
    
    - /updateReplyComment
        - idx (수정을 할 답글의 인덱스, Require), comment(작성할 답글, Require)
        - JWT