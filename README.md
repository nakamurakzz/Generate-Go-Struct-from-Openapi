# Files
```bash
.
├── pkl
│   ├── Group.pkl
│   └── Person.pkl
├── presentation
│   └── response
│       ├── Group
│       │   ├── Group.pkl.go
│       │   └── init.pkl.go
│       └── person
│           ├── Person.pkl.go
│           └── init.pkl.go
├── main.go : entry point
```

- pkl: directory for pkl definition files
- presentation: directory for generated files

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

