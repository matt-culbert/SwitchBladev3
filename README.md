# Cloak & Dagger

This is an evolution of the original Switchblade C2. Cloak refers to the C2 backend, hiding behind an mTLS reverse proxy, and Dagger is the implant which utilizes syscalls for command execution.

Run redis in a Docker container with ```docker run --name redis -p 6379:6379 -d redis```

When you run a command, you need to specify the UUID of the implant every time. To get a list of UUIDs in the redis db, enter '''5'''

![example](/img/guide/example.png)

Current commands:
- ```pwd``` gets the current working directory
- ```gcu``` gets the current user
- ```rc``` runs a command through the terminal, this can be anything (Still working on making commands work that are more than one word. So '''whoami''' works fine but '''cat /etc/passwd''' has issues
- ```rd``` reads the supplied directory. Use it with '''rd <directory path>'''

To generate keys, run the ```crypto.py``` app and copy the public key PEM contents into the implant.go file.

### Known issues:
On Kali, change the redis host in the controller and listener to 127.0.0.1 from localhost.

### Todo: 

Core items:
- [ ] Generate shellcode from the builder by adding ```go build -buildmode=pie -o shellcode.bin .\beacon.go```
- [x] Generate UUID and public/private key pair and add these to the generated beacon
- [x] When the builder is finished, update the controller to lookup keys by UUID instead of "test"
- [ ] Change it so the listeners no longer need a check in procedure, have the builder maybe write to the redis DB with the UUID? Would solve some other issues too with checking in after the DB is wiped
- [ ] Write a generator for the listeners, could be as simple as changing their listening addresses to sys.argv
- [ ] Add a .NET appdomain function for running tools like SeatBelt
