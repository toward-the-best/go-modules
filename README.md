# go-modules
go module을 다루는 방법을 소개 합니다

# go mod 캐시 삭제
사내 네트워크에서
git 버전 변경이나 저장소 변경 등으로 사용중이 go.mod 의 변화가 있을때

# 빌드시 다음과 같은 패키지를 가져오지 못하는 403 에러가 발생한다.
go: github.aaa.com/ysoftman/sample@v0.1: unrecognized import path "github.aaa.com/ysoftman/sample" (https fetch: Get https://github.ysoftman.com/ysoftman/sample?go-get=1: Forbidden)

## proxy 환경 변수 중 no_proxy 에 추가한다.
export no_proxy=github.aaa.com,github.bbb.com

## 참고로 패키지 경로를 찾지 못하는 경우
https://github.com/golang/go/issues/27238#issuecomment-432793244
## mod 로 다운로드 받은 모든 패키지 삭제
go clean -modcache

## 테스트 결과물 캐시 삭제시
go clean -testcache

## 빌드 결과물 캐시 삭제시
go clean -cache


# json 관련 자료 
https://www.joinc.co.kr/w/man/12/golang/json
