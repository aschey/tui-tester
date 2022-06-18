module github.com/aschey/tui-tester

go 1.18

require (
	github.com/ActiveState/termtest v0.7.1
	github.com/ActiveState/termtest/expect v0.7.0
	github.com/ActiveState/vt10x v1.3.1
)

require (
	github.com/ActiveState/termtest/conpty v0.5.0 // indirect
	github.com/ActiveState/termtest/xpty v0.6.0 // indirect
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
	github.com/Netflix/go-expect v0.0.0-20201125194554-85d881c3777e // indirect
	github.com/creack/pty v1.1.11 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/kr/pty v1.1.8 // indirect
	golang.org/x/sys v0.0.0-20211019181941-9d821ace8654 // indirect
	golang.org/x/tools v0.1.11
)

replace github.com/ActiveState/termtest => github.com/aschey/termtest v0.7.2-0.20220618034454-44551b62ed90
