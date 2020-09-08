<!-- e.Static("/image", "profileimage")
	e.GET("/signin", controller.SignIn)
	e.GET("/showUserInfo", controller.ShowUserInfo)
	e.POST("/signup", controller.SignUp)
	e.POST("/updateProfileImage", controller.UpdateProfileImage, middleware.JWT([]byte("secret")), jwt.VerifyAccessToken) -->

# 아만다 (아무도 만나지 못한다)

## 주요 기능 
1. 회원가입 후 자신의 프로필 사진을 등록 
2. 사진과 내용을 배경으로 사용자들이 검색 또는 메인페이지에서 프로필을 보고 평점을 부여 
3. 별점 순위를 바탕으로 자신의 프로필 사진의 위치가 메인페이지에서 달라짐.
4. 마음에 들거나 신기한 프로필을 가진 사용자의 프로필에 댓글을 달 수 있음.

## 기대 효과 
1. 답답한 기숙사 속 적적한 학우들의 마음에 조금이나마 버팀목이 될만한 소울메이트를 찾음으로써 학교생활을 더욱 즐겁고 행복하게 만들어 줌.
2. 자신의 spec과 visual을 평가받음으로써 자존감을 높이거나, 현실을 직시하여 더욱 발전하게 되는 계기가 될 수 있음.
3. 몰랐던 학우들과 친하게 지내게 되는 계기가 될 수 있음.

## Stack
|                      | WEB 차승호, 정성훈     | Server 제정민        |             
|:--------------------:|:---------------:|:------------------:|
| Development Language | React, TypeScript | GO , postgresql & GORM     |               
| Development Tool     | Visual Studio Code  | Visual Studio Code |

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