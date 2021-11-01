# Simple make file and directories ptouch

tiny script made in go to make easy mkdir and touch

## Install
```

```

## Usage
```bash
ptouch a/b/c/c.txt

.
└── a
    └── b
        └── c
            └── c.txt
```

```bash
ptouch a/b/b.txt

.
└── a
    └── b
        ├── b.txt
        └── c
            └── c.txt
```
```bash
ptouch hello.txt

.
├── a
│   └── b
│       ├── b.txt
│       └── c
│           └── c.txt
└── hello.txt
```
