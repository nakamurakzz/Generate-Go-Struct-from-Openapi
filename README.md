# Files
```bash
.
├── config
│   └── server
│       ├── Server.pkl.go
│       └── init.pkl.go
├── pkl
│   ├── AppConfig.pkl
│   ├── Group.pkl
│   └── Person.pkl
└── presentation
    └── response
        ├── Group
        │   ├── Group.pkl.go
        │   └── init.pkl.go
        ├── person
        │   ├── Person.pkl.go
        │   └── init.pkl.go
        └── response.go
```

- pkl directory: for pkl definition files
- config/server directory: for generated go files
- presentation/response directory: for generated go files
  - response.go: manually written file to crete a struct for response

# Usage
``` bash
pkl-gen-go pkl/Person.pkl --base-path github.com/nakamurakzz/
pkl-gen-go pkl/Group.pkl --base-path github.com/nakamurakzz/
```

# Generated files
- presentation/response/person/Person.pkl.go
- presentation/response/person/init.pkl.go
- presentation/response/group/Group.pkl.go
- presentation/response/group/init.pkl.go

