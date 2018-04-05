# mypass -- a personal password manager

mypass is a simple command-line password manager that stores your passwords
in a secure way.

mypass is a work in progress.

# Getting started

1. Clone the source code.

2. You'll likely need to `go get` a few external packages before it will
   compile. Follow compile errors to get external packages.

3. Build in Go.

Right now you must manually create the local data file before running mypass:

```
mkdir ~/.mypass
touch ~/.mypass/data.json
```

You should be able to run mypass.

# Adding passwords

`mypass add`

You will be prompted for your master password, site name (anything to help
distinguish the password), and username. This will add the password information
to your `data.json` file created above.

# List passwords

`mypass list [site]`

You can list the passwords stored in your `data.json` with `mypass list`.
Optionally you can add the site name to see additional information.

# TODO

mypass is far from complete, future features include:

- Securely copying a site password to your clipboard
- Changing passwords
- Changing master password
- Specifying custom password requirements per site
- And more!
