# volk-go
Go (Golang) bindings for https://github.com/zeux/volk

How to use it:
```Go
	// Initialize Volk
	err := volk.Initialize()
	if err != nil {
		log.Fatalf("Volk init failed: %v", err)
	}

	// Passing Volk function to Vulkan loader
	vulkan.SetGetInstanceProcAddr(volk.GetProcAddr())

	// Initialize Vulkan wrapper
	vulkan.Init()
	...
```