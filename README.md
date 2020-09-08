<!-- e.Static("/image", "profileimage")
	e.GET("/signin", controller.SignIn)
	e.GET("/showUserInfo", controller.ShowUserInfo)
	e.POST("/signup", controller.SignUp)
	e.POST("/updateProfileImage", controller.UpdateProfileImage, middleware.JWT([]byte("secret")), jwt.VerifyAccessToken) -->

# 아만다 (아무도 만나지 못한다)

# API - json, query, form 형식 모두 지원

## @GET /showMyInfo
    - Request : JWT(Require)
    - Response : 내 프로필
        
## @GET /showUserInfo
    - Request : id (Optional), name (Optional)
    - Response : 검색된 유저들의 프로필

## @GET /showAllInfo
    - Request : 없음 
    - Response : 모든 유저들의 프로필

## @GET /showUserStar
    - Request : idx (검색할 유저의 인덱스 값, Require)
    - Response : 검색된 유저의 평점과 참여자의 수

## @GET /showComment
    - Request : id (검색할 유저의 아이디, Require)
    - Response : 검색된 유저의 프로필에 있는 모든 댓글

## @GET /showReplyComment
    - Request : idx (검색할 댓글의 인덱스, Require)
    - Response : 해당 댓글의 모든 답글

## @POST /signup
    - Request : id (Require), pw (Require), name (Require), ismanager (Optional), description (Optional)
    - Response : 회원가입 성공 여부

## @POST /signin
    - Request : id (아이디, Require), pw (패스워드, Require)
    - Response : 로그인 성공 여부

## @POST /updateProfileImage
    - Request : image (등록할 사진, Require), JWT
    - Response : 프로필 이미지 등록 성공 여부
## @POST /updateStar
    - Request : idx (평가할 유저의 인덱스, Require), star (평점, Require)
    - Response : 평점 등록 성공 여부
## @POST /updateMyInfo
    - Request : name (닉네임, Optional), description (소개, Optional)
    - Response : 프로필 변경 성공 여부
## @POST /createComment
    - Request : id (작성할 대상의 아이디, Require), comment (작성할 댓글, Require), JWT
    - Response : 댓글 작성 성공 여부
## @POST /updateComment
    - Request : idx (수정할 댓글의 인덱스, Require), comment (수정할 댓글, Require), JWT
    - Response : 댓글 수정 성공 여부
## @POST /createReplyComment
    - Request : idx (답글을 달 댓글의 인덱스, Require), comment(작성할 답글, Require), JWT
    - Response : 답글 작성 성공 여부
## @POST /updateReplyComment
    - Request : idx (수정을 할 답글의 인덱스, Require), comment(작성할 답글, Require), JWT
    - Response : 답글 수정 성공 여부