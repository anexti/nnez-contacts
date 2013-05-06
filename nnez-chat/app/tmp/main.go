package main

import (
	"flag"
	"reflect"
	"github.com/robfig/revel"
	websocket "code.google.com/p/go.net/websocket"
	controllers0 "github.com/robfig/revel/modules/static/app/controllers"
	_ "github.com/robfig/revel/samples/chat/app/chatroom"
	controllers "github.com/robfig/revel/samples/chat/app/controllers"
)

var (
	runMode    *string = flag.String("runMode", "", "Run mode.")
	port       *int    = flag.Int("port", 0, "By default, read from app.conf")
	importPath *string = flag.String("importPath", "", "Go Import Path for the app.")
	srcPath    *string = flag.String("srcPath", "", "Path to the source root.")

	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

func main() {
	flag.Parse()
	revel.Init(*runMode, *importPath, *srcPath)
	revel.INFO.Println("Running revel server")
	
	revel.RegisterController((*controllers.WebSocket)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Room",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					14: []string{ 
						"user",
					},
				},
			},
			&revel.MethodType{
				Name: "RoomSocket",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "ws", Type: reflect.TypeOf((**websocket.Conn)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Refresh)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Room",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					26: []string{ 
						"user",
						"events",
					},
				},
			},
			&revel.MethodType{
				Name: "Say",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "message", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Leave",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Application)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					12: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "EnterDemo",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "demo", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.LongPolling)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Room",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					14: []string{ 
						"user",
					},
				},
			},
			&revel.MethodType{
				Name: "Say",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "message", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "WaitMessages",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "lastReceived", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Leave",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers0.Static)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Serve",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeModule",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.DefaultValidationKeys = map[string]map[int]string{ 
		"github.com/robfig/revel/samples/chat/app/controllers.Application.EnterDemo": { 
			16: "user",
			17: "demo",
		},
	}
	revel.TestSuites = []interface{}{ 
	}

	revel.Run(*port)
}
