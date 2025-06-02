# Go Zero - Beginner's Guide

## What is go-zero?

`go-zero` is a powerful Go framework for building web and RPC services. It's designed to be:

- **Easy to use**: Simple API definitions and code generation
- **Reliable**: Built-in features for handling high traffic
- **Fast**: Optimized for performance

## Key Features for Beginners

- **Simple API Definition**: Write your API in a simple language (.api files)
- **Automatic Code Generation**: Let the computer write the boring code for you
- **Built-in Validation**: Automatically check if user inputs are correct
- **Frontend Integration**: Generate TypeScript code for your web frontend
- **Ready for Production**: Includes features like timeouts and rate limiting

## Prerequisites

Before you start, make sure you have:

1. **Go installed**: Version 1.16 or newer
   ```bash
   # Check your Go version
   go version
   ```

2. **goctl tool**: This is go-zero's magic wand for generating code
   ```bash
   # Install goctl
   go install github.com/zeromicro/go-zero/tools/goctl@latest
   
   # Add it to your PATH
   export PATH=$PATH:$HOME/go/bin
   ```

3. **esbuild**: For compiling TypeScript (for the frontend)
   ```bash
   # Install esbuild (requires Node.js)
   npm install -g esbuild
   ```

4. **Basic Go knowledge**: Understanding of structs, HTTP, and JSON will help

## Project Overview

This project is a "Galactic Explorer" application that helps you find habitable exoplanets. It consists of:

1. **Backend API**: A Go-Zero REST API that provides exoplanet data
2. **Frontend**: A simple web interface to search for exoplanets

### What You'll Build

- A REST API endpoint to query exoplanets by distance and habitability
- A web interface to interact with the API
- Automatically generated TypeScript code for frontend-backend communication

### Project Structure

```bash
22_10_go-zero/
├── etc/
│   └── explorer-api.yaml       # Server configuration
├── internal/                   # Backend code (auto-generated)
│   ├── config/                 # Configuration handling
│   ├── handler/                # HTTP request handlers
│   ├── logic/                  # Business logic
│   ├── svc/                    # Service context
│   └── types/                  # Data types
├── frontend/                   # Web interface
│   ├── public/                 # Static HTML files
│   ├── src/                    # TypeScript code
│   └── server.go               # Frontend server
├── explorer.api                # API definition file
├── explorer.go                 # Main backend entry point
├── go.mod                      # Go dependencies
└── Makefile                    # Build automation
```

## Step-by-Step Guide

### 1. Define Your API

The first step is to define your API in the `explorer.api` file. This is a simple language that tells go-zero what your API should look like.

```
syntax = "api"

info (
	title:   "Galactic Explorer API"
	desc:    "API to query habitable exoplanets"
	author:  "Jay"
	version: "0.1"
)

type ExoplanetQueryRequest {
	MaxDistanceLy   int64   `json:"max_distance_ly,range=[1:100000]"`
	MinHabitability float64 `json:"min_habitability,range=[0:1]"`
}

type Exoplanet {
	Name         string  `json:"name"`
	DistanceLy   int64   `json:"distance_ly"`
	Habitability float64 `json:"habitability"`
}

type ExoplanetQueryResponse {
	Exoplanets []Exoplanet `json:"exoplanets"`
}

service explorer-api {
	@handler queryExoplanets
	post /exoplanets/query (ExoplanetQueryRequest) returns (ExoplanetQueryResponse)
}
```

**What's happening here?**
- We define the data structures for our request and response
- We set up validation (e.g., distance must be between 1 and 100,000 light-years)
- We create a POST endpoint at `/exoplanets/query`

### 2. Generate Code with goctl

Instead of writing all the code by hand, we use `goctl` to generate it for us:

```bash
# Generate backend Go code
goctl api go -api explorer.api -dir .

# Generate frontend TypeScript code
goctl api ts -api explorer.api -dir frontend/src
```

This creates:
- Backend code in the `internal/` directory
- TypeScript interfaces in `frontend/src/`

### 3. Implement Business Logic

The only part you need to write yourself is the business logic in `internal/logic/queryexoplanetslogic.go`:

```go
func (l *QueryExoplanetsLogic) QueryExoplanets(req *types.ExoplanetQueryRequest) (resp *types.ExoplanetQueryResponse, err error) {
	// Sample exoplanet data
	exoplanets := []ExoplanetData{
		{Name: "Kepler-442b", DistanceLy: 1200, Habitability: 0.84},
		{Name: "Proxima Centauri b", DistanceLy: 4, Habitability: 0.65},
		{Name: "TRAPPIST-1e", DistanceLy: 40, Habitability: 0.77},
	}

	// Filter exoplanets based on request criteria
	var filtered []types.Exoplanet
	for _, ep := range exoplanets {
		if ep.DistanceLy <= req.MaxDistanceLy && ep.Habitability >= req.MinHabitability {
			filtered = append(filtered, types.Exoplanet{
				Name:         ep.Name,
				DistanceLy:   ep.DistanceLy,
				Habitability: ep.Habitability,
			})
		}
	}

	return &types.ExoplanetQueryResponse{
		Exoplanets: filtered,
	}, nil
}
```

### 4. Frontend Implementation

The project includes a simple web interface:

- `frontend/public/index.html`: The HTML page with a form to search for exoplanets
- `frontend/src/*.ts`: TypeScript code for API communication
- `frontend/server.go`: A simple Go server to serve the frontend

### 5. Run the Application

The Makefile makes it easy to run everything:

```bash
# Start both backend and frontend servers
make run
```

This will:
1. Generate all necessary code
2. Compile TypeScript to JavaScript
3. Start the backend server on port 8800
4. Start the frontend server on port 3000

Then open http://localhost:3000 in your browser to use the application!

## What You've Learned

- How to define APIs using go-zero's API language
- How to generate code automatically with goctl
- How to implement business logic in go-zero
- How to connect a frontend to your go-zero API
- How to run a complete web application

## Benefits of go-zero for Beginners

- **Less Boilerplate**: Generate code instead of writing it
- **Built-in Validation**: No need to write validation code
- **Full-Stack Support**: Generate both backend and frontend code
- **Production-Ready**: Includes features needed for real-world applications

## Next Steps

To learn more about go-zero:

1. Check out the [official go-zero documentation](https://go-zero.dev/)
2. Explore more examples in the [go-zero GitHub repository](https://github.com/zeromicro/go-zero)
3. Try adding more endpoints to this example project
4. Experiment with connecting to a database instead of using static data

## References

- [Go-zero Docs](https://go-zero.dev/docs/introduction)
