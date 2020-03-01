package environment


var Aliases = map[string]string {
	"c++": "c++_gcc",
	"c++14": "c++14_gcc",
	"c": "c_gcc",
	"python": "python3",
}

// NOTE: https://language-test-201603.contest.atcoder.jp/
var Environments = map[string]*Environment{
	"c++14_gcc": {
		Key:          "c++14_gcc",
		Language:     "C++14 (GCC 5.4.1)",
		SrcName:      "main.cpp",
		Template:     "internal/c++/main.cpp",
		LanguageCode: "3003",

		BuildCmd: "g++ -std=gnu++1y -O2 -o a.out main.cpp",
		Cmd:      "./a.out",
		CleanCmd: "rm ./a.out",

		DockerImage:      "docker.io/sachaos/atcoder-gcc:latest",
		BuildCmdOnDocker: "g++ -std=gnu++1y -I /opt/boost/boost_1_60_0 -L /opt/boost/boost_1_60_0 -O2 -o a.out main.cpp",
		CmdOnDocker:      "./a.out",
	},
	"c++_gcc": {
		Key:          "c++_gcc",
		Language:     "C++ (GCC 5.4.1)",
		SrcName:      "main.cpp",
		Template:     "internal/c++/main.cpp",
		LanguageCode: "3029",

		BuildCmd: "g++ -std=gnu++03 -O2 -o a.out main.cpp",
		Cmd:      "./a.out",
		CleanCmd: "rm ./a.out",

		DockerImage:      "docker.io/sachaos/atcoder-gcc:latest",
		BuildCmdOnDocker: "g++ -std=gnu++03 -I /opt/boost/boost_1_60_0 -L /opt/boost/boost_1_60_0 -O2 -o a.out main.cpp",
		CmdOnDocker:      "./a.out",
	},
	"go": {
		Key:          "go",
		Language:     "Go (1.6)",
		SrcName:      "main.go",
		Template:     "internal/go/main.go",
		LanguageCode: "3013",

		BuildCmd: "go build -o ./a.out main.go",
		Cmd:      "./a.out",
		CleanCmd: "rm ./a.out",

		DockerImage:      "docker.io/library/golang:1.6",
		BuildCmdOnDocker: "go build -o a.out main.go",
		CmdOnDocker:      "./a.out",
	},
	"python3": {
		Key:          "python3",
		Language:     "Python3 (3.4.3)",
		SrcName:      "main.py",
		Template:     "internal/python3/main.py",
		LanguageCode: "3023",

		Cmd: "python3 -B main.py",

		CmdOnDocker: "python -B main.py",
		DockerImage: "docker.io/sachaos/atcoder-python3:latest",
	},
	"python2": {
		Key:          "python2",
		Language:     "Python2 (2.7.6)",
		SrcName:      "main.py",
		Template:     "internal/python2/main.py",
		LanguageCode: "3022",

		Cmd: "python2 -B main.py",

		DockerImage: "docker.io/sachaos/atcoder-python2:latest",
		CmdOnDocker: "python -B main.py",
	},
	"ruby": {
		Key:          "ruby",
		Language:     "Ruby (2.3.3)",
		SrcName:      "main.rb",
		Template:     "internal/ruby/main.rb",
		LanguageCode: "3024",

		BuildCmd: "ruby --disable-gems -w -c main.rb",
		Cmd:      "ruby --disable-gems main.rb",

		DockerImage:      "docker.io/library/ruby:2.3.3-alpine",
		BuildCmdOnDocker: "ruby --disable-gems -w -c main.rb",
		CmdOnDocker:      "ruby --disable-gems main.rb",
	},
	"java7": {
		Key:          "java7",
		Language:     "Java7 (OpenJDK 1.7.0)",
		SrcName:      "Main.java",
		Template:     "internal/java/Main.java",
		LanguageCode: "3015",

		BuildCmd: "javac Main.java",
		Cmd:      "java -Xss256M Main",

		DockerImage:      "docker.io/library/openjdk:7",
		BuildCmdOnDocker: "javac Main.java",
		CmdOnDocker:      "java -Xss256M Main",
	},
	"java8": {
		Key:          "java8",
		Language:     "Java8 (OpenJDK 1.8.0)",
		SrcName:      "Main.java",
		Template:     "internal/java/Main.java",
		LanguageCode: "3016",

		BuildCmd: "javac Main.java",
		Cmd:      "java -Xss256M Main",

		DockerImage:      "docker.io/library/openjdk:8",
		BuildCmdOnDocker: "javac Main.java",
		CmdOnDocker:      "java -Xss256M Main",
	},
	"c_gcc": {
		Key:          "c_gcc",
		Language:     "C (GCC 5.4.1)",
		SrcName:      "main.c",
		Template:     "internal/c/main.c",
		LanguageCode: "3004",

		BuildCmd: "gcc -std=gnu11 -O2 -o a.out main.c -lm",
		Cmd:      "./a.out",

		DockerImage:      "docker.io/sachaos/atcoder-gcc:latest",
		BuildCmdOnDocker: "gcc -std=gnu11 -O2 -o a.out main.c -lm",
		CmdOnDocker:      "./a.out",
	},
}

