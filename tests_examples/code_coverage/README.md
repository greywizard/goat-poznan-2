# Code coverage example

Code example is taken from [https://smartystreets.com/blog/2015/02/go-testing-part-1-vanillla](https://smartystreets.com/blog/2015/02/go-testing-part-1-vanillla)

To generate coverage run:

```Bash
mkdir coverage
go test -covermode=atomic -coverprofile=coverage/test.cover
go tool cover -html=coverage/test.cover -o coverage/profile.html
open coverage/profile.html
```